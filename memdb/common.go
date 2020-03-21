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
	transaction [][]interface{}
}

// NewDB instantiates a MemDB object.
func NewDB() *MemDB {
	return &MemDB{
		Data:  make(map[string]string),
		index: make(map[string]set.Set),
	}
}

// Set will set the key in the database to the given value
func (d *MemDB) Set(key, value string) {
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
}

// Get returns the value for the given name, else,
// a nil pointer is returned.
func (d *MemDB) Get(key string) *string {
	if value, ok := d.Data[key]; ok {
		return &value
	}

	return nil
}

// Delete will remove a key from the database.
func (d *MemDB) Delete(key string) {
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
