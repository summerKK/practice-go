package omap

import (
	"sort"
	"reflect"
)

type keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	ElemType() reflect.Type
	CompareFunc() func(interface{}, interface{}) int8
}

type mykeys struct {
	container  []interface{}
	compareFun func(interface{}, interface{}) uint8
	elemType   reflect.Type
}

func (keys *mykeys) Len() int {
	return len(keys.container)
}

func (keys *mykeys) Less(i, j int) bool {
	return keys.compareFun(keys.container[i], keys.container[j]) == -1
}

func (keys *mykeys) Swap(i, j int) {
	keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}

func (keys *mykeys) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}
	if reflect.TypeOf(k) != keys.elemType {
		return false
	}
	return true
}

func (keys *mykeys) Add(k interface{}) bool {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return false
	}
	keys.container = append(keys.container, k)
	sort.Sort(keys)
	return true
}

func (keys *mykeys) Search(k interface{}) (index int, contains bool) {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return
	}
	index = sort.Search(
		keys.Len(),
		func(i int) bool {
			return keys.compareFun(keys.container[i], k) >= 0
		})

	//
	if index < keys.Len() && keys.container[index] == k {
		contains = true
	}
	return
}
