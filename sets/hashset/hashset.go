// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package hashset

import (
	"fmt"
	"github.com/qauzy/gods/sets"
	"strings"
)

func assertSetImplementation() {
	var _ sets.Set[int] = (*Set[int])(nil)
}

// Set holds elements in go's native map
type Set[T comparable] struct {
	items map[T]struct{}
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[TP comparable](values ...TP) *Set[TP] {
	set := &Set[TP]{items: make(map[TP]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[T]) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[T]) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	set.items = make(map[T]struct{})
}

// Values returns all items in the set.
func (set *Set[T]) Values() []T {
	values := make([]T, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Set[T]) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}
