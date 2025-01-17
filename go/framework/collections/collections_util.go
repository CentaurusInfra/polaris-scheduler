package collections

import (
	"container/list"
)

var (
	_ LessFunc[string] = StringLessFn
)

// Defines a comparison function that returns true iff itemA is less than itemB, otherwise false.
type LessFunc[T any] func(itemA T, itemB T) bool

// Returns the key for the specified value (can be used for maps or similar data structures).
type KeyFunc[K ~int | ~string, V any] func(value V) (K, error)

func ConvertToLinkedList[T any](elements []T) *list.List {
	linkedList := list.New()
	for _, e := range elements {
		linkedList.PushBack(e)
	}
	return linkedList
}

func ConvertToSlice[T any](linkedList *list.List) []T {
	slice := make([]T, linkedList.Len())
	i := 0
	for e := linkedList.Front(); e != nil; e = e.Next() {
		slice[i] = e.Value.(T)
		i++
	}
	return slice
}

// Appends the elements of src to the end of dest, without making any copies.
func AppendToLinkedList(dest *list.List, src *list.List) {
	for e := src.Front(); e != nil; e = e.Next() {
		dest.PushBack(e.Value)
	}
}

// Swaps the elements in the positions i and j.
func Swap[T any](slice []T, i int, j int) {
	temp := slice[j]
	slice[j] = slice[i]
	slice[i] = temp
}

// LessFunc implementation for strings.
func StringLessFn(itemA string, itemB string) bool {
	return itemA < itemB
}
