package bitset

type MapSet map[interface{}]struct{}

func NewMapSet(elements ...interface{}) *MapSet {
	set := make(MapSet)
	for _, e := range elements {
		set[e] = struct{}{}
	}

	return &set
}

// 增删改查
func (set *MapSet) Add(e interface{}) {
	(*set)[e] = struct{}{}
}

func (set *MapSet) Remove(e interface{}) {
	delete(*set, e)
}

func (set *MapSet) Contain(e interface{}) bool {
	_, ok := (*set)[e]
	return ok
}

func (set *MapSet) Cap() int {
	return len(*set)
}

// 交并差
func (set *MapSet) Intersection(input *MapSet) *MapSet {
	newSet := NewMapSet()

	if set.Cap() > input.Cap() {
		for e, _ := range *input {
			if set.Contain(e) {
				newSet.Add(e)
			}
		}
	} else {
		for e, _ := range *set {
			if input.Contain(e) {
				newSet.Add(e)
			}
		}
	}

	return newSet
}

func (set *MapSet) Union(input *MapSet) *MapSet {
	if set.Cap() > input.Cap() {
		for e, _ := range *input {
			if !set.Contain(e) {
				set.Add(e)
			}
		}
		return set
	} else {
		for e, _ := range *set {
			if !input.Contain(e) {
				input.Add(e)
			}
		}
		return input
	}
}

func (set *MapSet) Diff(input *MapSet) *MapSet {
	newSet := NewMapSet()

	for e, _ := range *set {
		if !input.Contain(e) {
			newSet.Add(e)
		}
	}

	return newSet
}
