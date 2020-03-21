// Copyright (c) 2020 Ramon Quitales
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package memdb

// MemDB provides methods to to interact with the underlying
// data structure used to implement the in-memory database.
type MemDB struct {
	Data        map[string]string
	transaction [][]interface{}
}

// NewDB instantiates a MemDB object.
func NewDB() *MemDB {
	return &MemDB{
		Data: make(map[string]string),
	}
}

// Set will set the key in the database to the given value
func (d *MemDB) Set(key, value string) {
	d.Data[key] = value
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
	delete(d.Data, key)
}

// Count determines the number of keys that have the given
// values assigned to them. Returns 0, if none are found.
func (d *MemDB) Count(value string) (count int) {
	for _, v := range d.Data {
		switch v {
		case value:
			count++
		}
	}
	return
}
