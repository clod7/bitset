package bitset

import (
	"log"
	"testing"
)

func TestMapSet(t *testing.T) {
	s1 := NewMapSet(1, 2, 3, 7, 9)
	log.Println(s1)

	s1.Add(4)
	log.Println(s1)

	s1.Remove(2)
	log.Println(s1)

	s1 = NewMapSet(1, 2, 3, 7, 9)
	s2 := NewMapSet(2, 3, 4, 5)
	log.Println(s1.Intersection(s2))

	s1 = NewMapSet(1, 2, 3, 7, 9)
	s2 = NewMapSet(2, 3, 4, 5)
	log.Println(s1.Union(s2))

	s1 = NewMapSet(1, 2, 3, 7, 9)
	s2 = NewMapSet(2, 3, 4, 5)
	log.Println(s1.Diff(s2))
}
