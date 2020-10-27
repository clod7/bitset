package bitset

import (
	"math/bits"
)

const (
	shift = 6    // 2^n=64, n=6
	mask  = 0x3f // 2^n-1
)

type BitSet struct {
	data []uint64
	size int
}

func row(n int) int {
	return n >> shift
}

func column(n int) uint64 {
	return 1 << uint64(n&mask)
}

func NewBitSet(nums ...int) *BitSet {
	if len(nums) == 0 {
		return new(BitSet)
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	if max < 0 {
		return new(BitSet)
	}

	rows := row(max)
	s := &BitSet{
		data: make([]uint64, rows+1),
		size: 0,
	}

	for _, num := range nums {
		col := column(num)
		row := row(num)
		s.data[row] |= col
		s.size++
	}

	return s
}

func (s *BitSet) Add(n int) *BitSet {
	if n < 0 {
		return s
	}

	r := row(n)
	if r >= len(s.data) {
		newData := make([]uint64, r+1)
		copy(newData, s.data)
		s.data = newData
	}

	s.data[r] |= column(n)
	s.size++
	return s
}

func (s *BitSet) Contain(n int) bool {
	if n < 0 {
		return false
	}

	r := row(n)
	if r >= len(s.data) {
		return false
	}

	c := column(n)
	return s.data[r]&c != 0
}

func (s *BitSet) Remove(n int) *BitSet {
	if n < 0 {
		return s
	}

	r := row(n)
	if r >= len(s.data) {
		return s
	}

	c := column(n)
	s.data[r] &^= c
	s.size--
	return s
}

func (s *BitSet) Size() int {
	return s.size
}

func (s *BitSet) refreshSize() int {
	count := 0
	for _, data := range s.data {
		count += bits.OnesCount64(data)
	}

	return count
}

func (s *BitSet) Intersection(input *BitSet) *BitSet {
	if len(s.data) < len(input.data) {
		s, input = input, s
	}

	newSet := &BitSet{
		data: make([]uint64, len(input.data)),
	}
	for i := range input.data {
		newSet.data[i] = s.data[i] & input.data[i]
	}
	newSet.size = newSet.refreshSize()

	return newSet
}

func (s *BitSet) Union(input *BitSet) *BitSet {
	if len(s.data) < len(input.data) {
		s, input = input, s
	}

	for i := range input.data {
		s.data[i] |= input.data[i]
	}
	s.size = s.refreshSize()

	return s
}

func (s *BitSet) Diff(input *BitSet) *BitSet {
	sizeS, sizeInput := len(s.data), len(input.data)
	if sizeS < sizeInput {
		sizeS, sizeInput = sizeInput, sizeS
	}
	for i := 0; i < sizeInput; i++ {
		s.data[i] &^= input.data[i]
	}
	s.size = s.refreshSize()

	return s
}
