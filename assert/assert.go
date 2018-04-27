package assert

import (
	"testing"
)

// 内置一个测试对象
var t *testing.T

func BeginTest(tin *testing.T) {
	t = tin
}

func EndTest() {
	t = nil
}

func True(val bool, msg string) {
	if !val {
		t.Errorf("assert fails, expect true: %s\n", msg)
	}
}

func Equal(got, expect interface{}, msg string) {
	if got != expect {
		t.Errorf("assert fails, got(%v) but expect(%v): %s\n", got, expect, msg)
	}
}

func InSlice(val interface{}, slice []interface{}, msg string) {
	for _, v := range slice {
		if v == val {
			return
		}
	}
	t.Errorf("assert fails, element(%v) not in slice: %s\n", val, msg)
}

// 只适用于以 string 为键的 map 类型
func HasHash(val interface{}, hash map[string]interface{}, msg string) {
	for _, v := range hash {
		if v == val {
			return
		}
	}
	t.Errorf("assert fails, element(%v) not in hash: %s\n", val, msg)
}

func InHash(key string, hash map[string]interface{}, msg string) {
	if _, ok := hash[key]; ok {
		return
	}
	t.Errorf("assert fails, key(%v) not in hash: %s\n", key, msg)
}
