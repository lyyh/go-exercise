package main

import (
	"reflect"
	"testing"
)

// 字符串包含
func StringsContains(arr []string, val string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

// 包含
func Contains(arr []interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	return
}

func BenchmarkContains(b *testing.B) {
	sa := []string{"a", "b", "c"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringsContains(sa, "a")
	}
}
