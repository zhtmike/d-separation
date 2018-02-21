package dsep

import (
	"reflect"
	"testing"
)

func TestFindDSeperatioWithCommonCause(t *testing.T) {
	var adjlist = make([][]int, 3)
	adjlist[0] = []int{1, 2}
	adjlist[1] = []int{}
	adjlist[2] = []int{}
	ans, err := FindDSeperation(adjlist, 1, []int{0})
	if err != nil {
		t.Error(err)
	} else {
		expect := []int{2}
		if !reflect.DeepEqual(ans, expect) {
			t.Error("Expected ", expect, " got ", ans)
		}
	}
}

func TestFindDSeperatioWithTrail(t *testing.T) {
	var adjlist = make([][]int, 3)
	adjlist[0] = []int{1}
	adjlist[1] = []int{2}
	adjlist[2] = []int{}
	ans, err := FindDSeperation(adjlist, 0, []int{1})
	if err != nil {
		t.Error(err)
	} else {
		expect := []int{2}
		if !reflect.DeepEqual(ans, expect) {
			t.Error("Expected ", expect, " got ", ans)
		}
	}
}

func TestFindDSeperatioWithCommonEffect(t *testing.T) {
	var adjlist = make([][]int, 3)
	adjlist[0] = []int{}
	adjlist[1] = []int{0}
	adjlist[2] = []int{0}
	ans, err := FindDSeperation(adjlist, 1, []int{})
	if err != nil {
		t.Error(err)
	} else {
		expect := []int{2}
		if !reflect.DeepEqual(ans, expect) {
			t.Error("Expected ", expect, " got ", ans)
		}
	}
}
