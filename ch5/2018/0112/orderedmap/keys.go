package omap

import (
	"sort"
	"reflect"
	"bytes"
	"fmt"
)

type Keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	ElemType() reflect.Type
	CompareFunc() compareFunc
	String() string
}

// compareFunc的结果值：
//   小于0: 第一个参数小于第二个参数
//   等于0: 第一个参数等于第二个参数
//   大于1: 第一个参数大于第二个参数
type mykeys struct {
	container  []interface{}
	compareFun compareFunc
	elemType   reflect.Type
}

type compareFunc func(interface{}, interface{}) int8

func NewKeys(cofu compareFunc, elemType reflect.Type) *mykeys {
	return &mykeys{
		container:  make([]interface{}, 0),
		compareFun: cofu,
		elemType:   elemType,
	}
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

	if index < keys.Len() && keys.container[index] == k {
		contains = true
	}
	return
}

func (keys *mykeys) Remove(k interface{}) bool {
	index, exist := keys.Search(k);
	if !exist {
		return false
	}
	//找到对应下标的数据并且删除掉
	keys.container = append(keys.container[:index], keys.container[index+1:]...)
	return true
}

func (keys *mykeys) Clear() {
	keys.container = make([]interface{}, 0)
}

func (keys *mykeys) Get(index int) interface{} {
	if index >= keys.Len() {
		return nil
	}
	return keys.container[index]
}

func (keys *mykeys) GetAll() []interface{} {
	initialLen := keys.Len()
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for _, v := range keys.container {
		if actualLen < initialLen {
			snapshot[actualLen] = v
		} else {
			snapshot = append(snapshot, v)
		}
		actualLen++
	}
	if actualLen < initialLen {
		return snapshot[:actualLen]
	}
	return snapshot

}

func (keys *mykeys) ElemType() reflect.Type {
	return keys.elemType
}

func (keys *mykeys) CompareFunc() compareFunc {
	return keys.compareFun
}

func (keys *mykeys) String() string {
	var buf bytes.Buffer
	buf.WriteString("keys{")
	first := true
	for _, v := range keys.container {
		if !first {
			buf.WriteString(" ")
		} else {
			first = true
		}
		buf.WriteString(fmt.Sprintf("%v", v))
	}
	buf.WriteString("}")
	return buf.String()
}
