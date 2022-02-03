package slice

import "reflect"

// SpliceInsertValue 实现在 slice 中任意位置插入元素
// value 为要插入的元素 index 为插入索引位置
// 异常时候会panic
func SpliceInsertValue(slice interface{}, value interface{}, index int) interface{} {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("slice需要是切片类型")
	}
	if reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		panic("插入的元素需要和切片同种类型")
	}

	sliceLen := reflect.ValueOf(slice).Len()
	if index > sliceLen || index < 0 {
		panic("插入位置异常")
	}

	oldSlice := reflect.MakeSlice(reflect.TypeOf(slice), sliceLen, sliceLen)
	reflect.Copy(oldSlice, reflect.ValueOf(slice))

	pre := reflect.ValueOf(slice).Slice(0, index)

	newL := reflect.Append(pre, reflect.ValueOf(value))
	newL = reflect.AppendSlice(newL, oldSlice.Slice(index, sliceLen))
	return newL.Interface()
}
