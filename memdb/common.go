// Copyright (c) 2020 Ramon Quitales
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package memdb

import (
	"github.com/rquitales/go-memdb/set"
)

// MemDB provides methods to to interact with the underlying
// data structure used to implement the in-memory database.
type MemDB struct {
	Data        map[string]string
	index       map[string]set.Set
	transaction []Changes
}

// NewDB instantiates a MemDB object.
func NewDB() *MemDB {
	return &MemDB{
		Data:        make(map[string]string),
		index:       make(map[string]set.Set),
		transaction: []Changes{},
	}
}

// Set will set the key in the database to the given value
func (d *MemDB) Set(key, value string) {
	// Check if the key existed before
	var before interface{}
	var ok bool
	if before, ok = d.Data[key]; !ok {
		before = nil
	}

	// Check if a transaction exists, if so, log
	// the action to the transaction log
	if len(d.transaction) != 0 {
		temp := d.transaction[len(d.transaction)-1]

		temp.Push(before, value, key)
		d.transaction[len(d.transaction)-1] = temp
	}

	d.Data[key] = value

	// Add value to 'index' hash map for O(1) look-up for COUNT
	{
		temp := d.index[value]

		// Initialise a new set if required
		if temp == nil {
			temp = make(set.Set)
		}

		temp.Add(key)
		d.index[value] = temp
	}

	// Remove "old" value from index
	if ok && before.(string) != value {
		temp := d.index[before.(string)]
		if temp != nil {
			temp.Delete(key)
			d.index[before.(string)] = temp
		}
	}
}

// Get returns the value for the given name, else,
// a nil pointer is returned.
func (d *MemDB) Get(key string) (value string, ok bool) {
	value, ok = d.Data[key]

	return
}

// Delete will remove a key from the database.
func (d *MemDB) Delete(key string) {
	// Check if a transaction exists, if so, log
	// the action to the transaction log
	if len(d.transaction) != 0 {
		temp := d.transaction[len(d.transaction)-1]

		var before interface{}
		var ok bool
		if before, ok = d.Data[key]; !ok {
			before = nil
		}

		temp.Push(before, nil, key)
		d.transaction[len(d.transaction)-1] = temp
	}

	value := d.Data[key]
	delete(d.Data, key)

	// Delete key from index on 'values'
	{
		temp := d.index[value]
		if temp != nil {
			temp.Delete(key)
			d.index[value] = temp
		}
	}
}

// Count determines the number of keys that have the given
// values assigned to them. Returns 0, if none are found.
func (d *MemDB) Count(value string) (count int) {
	temp := d.index[value]

	return temp.Count()
}

// Begin starts a new database transaction.
func (d *MemDB) Begin() {
	d.transaction = append(d.transaction, Changes{})
}

// Commit will commit the current transaction. As we always keep
// writing/deleting/updating to the database, we commit by deleting
// the last transaction log.
func (d *MemDB) Commit() {
	if len(d.transaction) == 0 {
		return
	}

	d.transaction = d.transaction[0 : len(d.transaction)-1]
}

// Rollback will undo any changes made in the current transaction.
func (d *MemDB) Rollback() bool {
	if len(d.transaction) == 0 {
		return false
	}

	lastTx := d.transaction[len(d.transaction)-1]

	for len(lastTx) > 0 {
		change := lastTx.Pop()

		if change.None() {
			continue
		} else if change.Created() {
			RollbackCreate(d, change)
		} else if change.Deleted() {
			RollbackDelete(d, change)
		} else if change.Updated() {
			RollbackUpdate(d, change)
		}
	}

	// All change in latest transaction now undone, so remove from log/stack
	d.transaction = d.transaction[0 : len(d.transaction)-1]

	return true
}

// RollbackUpdate will revert any write updates to the database
// and re-index values.
func RollbackUpdate(d *MemDB, change Change) {
	d.Data[change.Key] = change.Before.(string)

	// Add old value to 'index' hash map for O(1) look-up for COUNT
	{
		temp := d.index[change.Before.(string)]

		// Initialise a new set if required
		if temp == nil {
			temp = make(set.Set)
		}

		temp.Add(change.Key)
		d.index[change.Before.(string)] = temp
	}

	// Remove "new" value from index
	{
		temp := d.index[change.After.(string)]
		if temp != nil {
			temp.Delete(change.Key)
			d.index[change.After.(string)] = temp
		}
	}
}

// RollbackDelete will undo any deletions in the database.
func RollbackDelete(d *MemDB, change Change) {
	d.Data[change.Key] = change.Before.(string)

	// Add value to 'index' hash map for O(1) look-up for COUNT
	{
		temp := d.index[change.Before.(string)]

		// Initialise a new set if required
		if temp == nil {
			temp = make(set.Set)
		}

		temp.Add(change.Key)
		d.index[change.Before.(string)] = temp
	}
}

// RollbackCreate will rollback any entries that were created
// and remove any index values.
func RollbackCreate(d *MemDB, change Change) {
	value := d.Data[change.Key]
	delete(d.Data, change.Key)

	// Delete key from index on 'values'
	{
		temp := d.index[value]
		if temp != nil {
			temp.Delete(change.Key)
			d.index[value] = temp
		}
	}
}
