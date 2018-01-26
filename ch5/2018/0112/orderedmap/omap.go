package omap

import (
	"reflect"
	"bytes"
	"fmt"
)

type OrderedMap interface {
	//获取给定键值对应的元素值,若没有对应元素值则返回nil
	Get(key interface{}) interface{}
	//添加键值对,并返回与给定简直对应的旧元素值.
	// 若没有旧元素值则返回(nil,true)
	Put(key interface{}, elem interface{}) (interface{}, bool)
	//删除与给定键值对应的键值对,并返回就得元素值.若没有旧元素值则返回nil
	Remove(key interface{}) interface{}
	//清除所有的键值对
	Clear()
	//获取键值对的数量
	Len() int
	//判断是否包含给定的键值
	Contains(key interface{}) bool
	//获取第一个键值,若无任何键值对则返回nil
	FirstKey() interface{}
	//获取最后一个键值,若无任何键值对则返回nil
	LastKey() interface{}
	//获取由小于键值toKey的键值所对应的键值对组成的OrderedMap类型值
	HeadMap(toKey interface{}) OrderedMap
	//获取由小于键值toKey且大于等于键值formKey的键值所对应的键值对
	// 组成的OrderedMap类型值
	SubMap(fromKey interface{}, toKey interface{}) OrderedMap
	//获取由大于等于键值fromKey的键值对所对应的键值对
	// 组成的OrderedMap类型值
	TailMap(fromKey interface{}) OrderedMap
	//获取已排序的键值所组成的切片值
	Keys() []interface{}
	//获取已排序的元素值所组成的切片值
	Elems() []interface{}
	//获取已排序的键值对所组成的字典值
	ToMap() map[interface{}]interface{}
	//获取键的类型
	KeyType() reflect.Type
	//获取元素的类型
	ElemType() reflect.Type
	String() string
}

type myOrderedMap struct {
	keys     Keys
	elemType reflect.Type
	m        map[interface{}]interface{}
}

func NewOrderedMap(keys Keys, elemType reflect.Type) OrderedMap {
	return &myOrderedMap{
		keys:     keys,
		m:        make(map[interface{}]interface{}),
		elemType: elemType,
	}
}

func (omap *myOrderedMap) IsAcceptableElem(e interface{}) bool {
	if e == nil {
		return false
	}
	if reflect.TypeOf(e) != omap.elemType {
		return false
	}
	return true
}

func (omap *myOrderedMap) Get(key interface{}) interface{} {
	k, ok := omap.m[key]
	if !ok {
		return nil
	}
	return k
}

//添加键值对,并返回与给定键值对应的旧元素值.
// 若没有旧元素值则返回(nil,true)
func (omap *myOrderedMap) Put(key interface{}, elem interface{}) (interface{}, bool) {
	if !omap.IsAcceptableElem(key) {
		return nil, false
	}
	oldElem, ok := omap.m[key]
	omap.m[key] = elem
	if !ok {
		omap.keys.Add(key)
	}
	return oldElem, true
}

//删除与给定键值对应的键值对,并返回旧得元素值.若没有旧元素值则返回nil
func (omap *myOrderedMap) Remove(key interface{}) interface{} {
	if !omap.IsAcceptableElem(key) {
		return nil
	}
	elem, ok := omap.m[key]
	if ok {
		delete(omap.m, key)
		omap.keys.Remove(key)
	}
	return elem
}

func (omap *myOrderedMap) Clear() {
	omap.m = make(map[interface{}]interface{}, 0)
	omap.keys.Clear()
}

func (omap *myOrderedMap) Len() int {
	return len(omap.m)
}

func (omap *myOrderedMap) Contains(key interface{}) bool {
	_, ok := omap.m[key]
	return ok
}

func (omap *myOrderedMap) FirstKey() interface{} {
	if omap.Len() == 0 {
		return nil
	}

	return omap.keys.Get(0)
}

func (omap *myOrderedMap) LastKey() interface{} {
	if omap.Len() == 0 {
		return nil
	}
	return omap.keys.Get(omap.Len() - 1)
}

func (omap *myOrderedMap) HeadMap(tokey interface{}) OrderedMap {
	return omap.SubMap(nil, tokey)
}

func (omap *myOrderedMap) TailMap(fromKey interface{}) OrderedMap {
	return omap.SubMap(fromKey, nil)
}

func (omap *myOrderedMap) SubMap(fromKey interface{}, toKey interface{}) OrderedMap {

	newOrderedMap := &myOrderedMap{
		keys:     NewKeys(omap.keys.CompareFunc(), omap.keys.ElemType()),
		elemType: omap.elemType,
		m:        make(map[interface{}]interface{}, 0),
	}
	omapLen := omap.Len()
	if omapLen == 0 {
		return newOrderedMap
	}
	beginIndex, contains := omap.keys.Search(fromKey)
	if !contains {
		beginIndex = 0
	}
	endIndex, contains := omap.keys.Search(toKey)
	if !contains {
		endIndex = omapLen
	}
	var key, elem interface{}
	for i := beginIndex; i < endIndex; i++ {
		key = omap.keys.Get(i)
		elem = omap.m[key]
		newOrderedMap.Put(key, elem)
	}
	return newOrderedMap
}

func (omap *myOrderedMap) Keys() []interface{} {
	if omap.Len() == 0 {
		return nil
	}
	initialLen := omap.keys.Len()
	keysVal := make([]interface{}, initialLen)
	actualLen := 0
	for _, k := range omap.keys.GetAll() {
		if actualLen < initialLen {
			keysVal[actualLen] = k
		} else {
			keysVal = append(keysVal, k)
		}
		actualLen++
	}
	if actualLen < initialLen {
		return keysVal[:actualLen]
	}
	return keysVal
}

func (omap *myOrderedMap) Elems() []interface{} {
	if omap.Len() == 0 {
		return nil
	}
	initialLen := omap.Len()
	elems := make([]interface{}, initialLen)
	actualLen := 0
	for _, v := range omap.keys.GetAll() {
		elem := omap.m[v]
		if actualLen < initialLen {
			elems[actualLen] = elem
		} else {
			elems = append(elems, elem)
		}
		actualLen++
	}
	if actualLen < initialLen {
		return elems[:actualLen]
	}
	return elems
}

func (omap *myOrderedMap) ToMap() map[interface{}]interface{} {
	replica := make(map[interface{}]interface{})
	for k, v := range omap.m {
		replica[k] = v
	}
	return replica
}

func (omap *myOrderedMap) KeyType() reflect.Type {
	return omap.keys.ElemType()
}

func (omap *myOrderedMap) ElemType() reflect.Type {
	return omap.elemType
}

func (omap *myOrderedMap) String() string {
	var buf bytes.Buffer
	buf.WriteString("orderedMap<")
	buf.WriteString(omap.keys.ElemType().Kind().String())
	buf.WriteString(",")
	buf.WriteString(omap.elemType.Kind().String())
	buf.WriteString(">{")
	first := true
	for _, v := range omap.Keys() {
		if !first {
			buf.WriteString(" ")
		} else {
			first = false
		}

		buf.WriteString(fmt.Sprintf("%v", v))
		buf.WriteString(":")
		buf.WriteString(fmt.Sprintf("%v\n", omap.Get(v)))
	}
	buf.WriteString("}")
	return buf.String()
}
