// Двусвязный список
package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_list(t *testing.T) {
	list := New()

	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)
	list.PushBack(4)

	test("list.Len()", list.Len(), 4, t)

	fromFirst := []int{}
	for e := list.First(); e != nil; e = e.Next() {
		fromFirst = append(fromFirst, e.Value.(int))
	}
	test("list.fromFirst", fromFirst, []int{1, 2, 3, 4}, t)

	fromLast := []int{}
	for e := list.Last(); e != nil; e = e.Prev() {
		fromLast = append(fromLast, e.Value.(int))
	}
	test("list.fromLast", fromLast, []int{4, 3, 2, 1}, t)
}

func test(name string, got, res interface{}, t *testing.T) {
	if !reflect.DeepEqual(got, res) {
		t.Errorf("%s = %v, want %v", name, got, res)
		fmt.Printf("%s: Fail\n", name)
	} else {
		fmt.Printf("%s: Ok\n", name)
	}
}
