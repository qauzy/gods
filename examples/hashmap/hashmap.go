// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/qauzy/gods/maps/hashmap"
)

// HashMapExample to demonstrate basic usage of HashMap
func main() {
	var m *hashmap.Map[int, string]
	m = hashmap.New[int, string]() // empty
	m.Put(1, "x")                  // 1->x
	m.Put(2, "b")                  // 2->b, 1->x  (random order)
	m.Put(1, "a")                  // 2->b, 1->a (random order)
	_, _ = m.Get(2)                // b, true
	_, _ = m.Get(3)                // nil, false
	_ = m.Values()                 // []interface {}{"b", "a"} (random order)
	_ = m.Keys()                   // []interface {}{1, 2} (random order)

	for _, k := range m.Keys() {
		v, _ := m.Get(k)
		fmt.Println("=====", k, v)
	}
	m.Remove(1) // 2->b
	m.Clear()   // empty
	m.Empty()   // true
	m.Size()    // 0
}
