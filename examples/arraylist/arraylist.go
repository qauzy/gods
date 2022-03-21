// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/qauzy/gods/lists/arraylist"
	"github.com/qauzy/gods/utils"
)

var srcCode = `
package main

import (
	"fmt"
)
type List[T comparable] struct {
	elements []T
	size     int
}

func main() {
	var m *hashmap.Map[int,string]
    m = hashmap.New[int, string]()
	//list = arraylist.New[string]()
}

`

// ArrayListExample to demonstrate basic usage of ArrayList
func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", srcCode, 0)
	if err != nil {
		fmt.Printf("err = %s", err)
	}
	ast.Print(fset, f)

	var list *arraylist.List[string]
	list = arraylist.New[string]()
	list.Add("a")                     // ["a"]
	list.Add("c", "b")                // ["a","c","b"]
	list.Sort(utils.StringComparator) // ["a","b","c"]
	_, _ = list.Get(0)                // "a",true
	it := list.Iterator()
	for it.Next() {
		fmt.Println("Val:", it.Value())
	}
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Swap(0, 1)                       // ["b","a",c"]
	list.Remove(2)                        // ["b","a"]
	list.Remove(1)                        // ["b"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
}
