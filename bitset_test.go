package bitset

import (
	"log"
	"testing"
)

func TestBitSet(t *testing.T) {
	s1 := NewBitSet(1, 2, 5, 8, 64, 65)
	log.Println(s1)

	s1.Add(128)
	log.Println(s1)

	s1.Remove(65)
	log.Println(s1)

	log.Println(s1.Contain(66))
	log.Println(s1.Contain(128))

	s1 = NewBitSet(1, 2, 5, 8, 64, 65)
	s2 := NewBitSet(1, 3, 4)
	log.Println(s1.Intersection(s2))

	s1 = NewBitSet(1, 2, 5, 8, 64, 65)
	s2 = NewBitSet(1, 3, 4)
	log.Println(s1.Union(s2))

	s1 = NewBitSet(1, 2, 5, 8, 64, 65)
	s2 = NewBitSet(1, 3, 4)
	log.Println(s1.Diff(s2))
}
