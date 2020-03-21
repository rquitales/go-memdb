// Copyright (c) 2020 Ramon Quitales
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package set

// Set implements a basic set data structure
// when we only require a single value to be hashed.
type Set map[string]struct{}

// Add will add a new value to our set
func (s *Set) Add(value string) {
	(*s)[value] = struct{}{}
}

// Count returns how many values are in our set
func (s *Set) Count() int {
	return len(*s)
}

// Delete removes a value from our set
func (s *Set) Delete(value string) {
	delete(*s, value)
}
