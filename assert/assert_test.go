package assert

import (
	"testing"
)

type INT int

func TestAssert(t *testing.T) {
	BeginTest(t)

	True(1+2 == 3, "1+2=3")
	True(1+2 == 4, "1+2=3")
	True("ab"+"CD" == "abcd", "ab + CD = abCD")
	True("ab"+"CD" == "abCD", "ab + CD = abCD")

	Equal(1+2, 3, "int add error")
	Equal(1+2, 4, "int add error")
	Equal("cd", "CD", "")
	Equal("CD", "CD", "")

	/*
		var slice = []INT{1, 2, 3, 4, 5}
		InSlice(3, slice, "some error")
		InSlice(-3, slice, "some error")

		var hash = map[string]INT{"a": 1, "b": 2, "c": 3}
		InHash("b", hash, "")
		InHash("B", hash, "")
		HasHash(2, hash, "")
		HasHash(-2, hash, "")
	*/

	EndTest()
}
