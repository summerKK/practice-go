package set

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}

func NewSimpleSet() Set {
	return NewHashSet()
}

// 判断集合 one 是否是集合 other 的超集
func IsSuperset(one, other Set) bool {
	if one == nil || other == nil {
		return false
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	if oneLen < otherLen {
		return false
	}
	for _, v := range other.Elements() {
		if !one.Contains(v) {
			return false
		}
	}
	return true
}

// 生成集合 one 和集合 other 的并集
func Union(one, other Set) Set {
	if one == nil || other == nil {
		return nil
	}
	unionSet := NewSimpleSet()
	for _, v := range one.Elements() {
		unionSet.Add(v)
	}
	if other.Len() == 0 {
		return unionSet
	}
	for _, v := range other.Elements() {
		unionSet.Add(v)
	}
	return unionSet
}

// 生成集合 one 和集合 other 的交集
func Intersect(one, other Set) Set {
	if one == nil || other == nil {
		return nil
	}
	intersectSet := NewSimpleSet()
	if one.Len() > other.Len() {
		for _, v := range other.Elements() {
			if one.Contains(v) {
				intersectSet.Add(v)
			}
		}
	} else {
		for _, v := range one.Elements() {
			if other.Contains(v) {
				intersectSet.Add(v)
			}
		}
	}
	return intersectSet
}

// 生成集合 one 对集合 other 的差集
func Difference(one, other Set) Set {
	if one == nil || other == nil {
		return nil
	}
	differenceSet := NewSimpleSet()
	if other.Len() == 0 {
		for _, v := range one.Elements() {
			differenceSet.Add(v)
		}
		return differenceSet
	}

	for _, v := range one.Elements() {
		if !other.Contains(v) {
			differenceSet.Add(v)
		}
	}
	return differenceSet
}
