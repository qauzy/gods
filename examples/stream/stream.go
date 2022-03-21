// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/qauzy/gods/stream"
)

type ABC struct {
	Name string
	Age  int
}

type BCD struct {
	Name string
	Age  int
}

func (this *ABC) GetAge() int {
	return this.Age
}

// HashMapExample to demonstrate basic usage of HashMap
func main() {
	var ss *stream.Stream[*ABC]
	ss = stream.Of(&ABC{Name: "aa", Age: 11}, &ABC{Name: "bb", Age: 22}, &ABC{Name: "cc", Age: 33}) // empty
	_ = ss

	ss.ForEach(PrintABC)

	ss.ForEach(func(each *ABC) {
		fmt.Printf("内置 Name=%v,Age=%v\n", each.Name, each.Age)
	})

	ss.Map(ABC.GetAge)

}

func PrintABC(abc *ABC) {
	fmt.Printf("Name=%v,Age=%v\n", abc.Name, abc.Age)
}
