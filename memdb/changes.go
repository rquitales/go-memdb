// Copyright (c) 2020 Ramon Quitales
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package memdb

// Changes contains a list of changes in the current transaction.
type Changes []Change

// Change describes the key/valu before and after an action.
type Change struct {
	Before interface{}
	After  interface{}
	Key    string
}

// Created determines if a new key/value entry was created in the database.
func (c *Change) Created() bool {
	return c.Before == nil && c.After != nil
}

// Updated determines if the action change the key/value data stored.
func (c *Change) Updated() bool {
	return c.Before != nil && c.After != nil
}

// Deleted determines if the action deletd a key from the database.
func (c *Change) Deleted() bool {
	return c.Before != nil && c.After == nil
}

// None determines if the action had no effect on the data ultimately.
func (c *Change) None() bool {
	return c.Before == c.After
}

// Pop gets the latests change from a slice of changes.
func (c *Changes) Pop() (val Change) {
	val, (*c) = (*c)[len(*c)-1], (*c)[:len(*c)-1]

	return
}

// Push appends a new action and its changes into a list of changes.
func (c *Changes) Push(before, after interface{}, key string) {
	*c = append(*c, Change{
		Before: before,
		After:  after,
		Key:    key,
	})
}
