// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraystack

//func assertSerializationImplementation() {
//	var _ containers.JSONSerializer = (*Stack)(nil)
//	var _ containers.JSONDeserializer = (*Stack)(nil)
//}

// ToJSON outputs the JSON representation of the stack.
func (stack *Stack[T]) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

// FromJSON populates the stack from the input JSON representation.
func (stack *Stack[T]) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}
