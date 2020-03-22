// Copyright (c) 2020 Ramon Quitales
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package memdb

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	db := NewDB()

	if got := db.Data; got == nil {
		t.Errorf("NewDB().Data = %+v; want map[string]string", got)
	}

	if got := db.index; got == nil {
		t.Errorf("NewDB().index = %+v; want map[string]set.Set", got)
	}

	if got := db.transaction; len(got) != 0 {
		t.Errorf("len(NewDB().transaction) = %d; want 0", len(got))
	}
}

func TestSet(t *testing.T) {
	db := NewDB()

	db.Set("foo", "bar")

	if value := db.Data["foo"]; value != "bar" {
		t.Errorf(`db.Set("foo", "bar") yielded foo=%v, want bar`, value)
	}
}

func TestGet(t *testing.T) {
	db := NewDB()

	db.Data["foo"] = "bar"

	if value, ok := db.Get("foo"); value != "bar" || !ok {
		t.Errorf(`db.Get("foo") yielded foo=%v, want foo=bar`, value)
	}
}

func TestDelete(t *testing.T) {
	db := NewDB()

	db.Data["foo"] = "bar"

	db.Delete("foo")

	if value, ok := db.Data["foo"]; value != "" && !ok {
		t.Errorf(`db.Delete("foo") failed, value yielded foo=%v, want foo=<nil>`, value)
	}
}

func TestCount(t *testing.T) {
	db := NewDB()

	db.Set("foo", "bar")
	db.Set("bar", "bar")

	value := db.Count("bar")

	if value != 2 {
		t.Errorf(`db.Count("bar") = %d, want 2`, value)
	}

	db.Delete("foo")
	value = db.Count("bar")

	if value != 1 {
		t.Errorf(`db.Count("bar") = %d, want 1`, value)
	}

	db.Set("bar", "bar")
	value = db.Count("bar")

	if value != 1 {
		t.Errorf(`db.Count("bar") = %d, want 1`, value)
	}

	db.Delete("bar")
	value = db.Count("bar")

	if value != 0 {
		t.Errorf(`db.Count("bar") = %d, want 0`, value)
	}
}
