package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

//添加元素
func (set *HashSet) Add(e interface{}) bool {
	//判断是否已经存在set里面
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

//删除元素
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

//清空
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

//判断是否存在某个数据
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

//获取长度
func (set *HashSet) Len() int {
	return len(set.m)
}

//判断两个hashSet类型数据的值是否相等
func (set *HashSet) Same(other Set) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	//这里多加了判断是考虑到我们在迭代hashset的时候,
	//hashset里面的数据可能会发生变化
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}
