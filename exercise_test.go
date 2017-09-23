package golang

import (
	"reflect"
	"testing"
)

func TestAppendSlice(t *testing.T) {
	data := []string{"go", "java", "node", "clojure", "python"}
	add := []string{"react", "angular", "vue"}

	merged := concat(data, add)

	expected := []string{"go", "java", "node", "clojure", "python", "react", "angular", "vue"}

	if !reflect.DeepEqual(merged, expected) {
		t.Errorf("%v is expected but got %v", expected, merged)
	}
}

func TestGetRidOfTheFirst(t *testing.T) {
	data := []string{"go", "java", "node", "clojure", "python", "react", "angular", "vue"}

	r := delFirst(data)

	expected := []string{"java", "node", "clojure", "python", "react", "angular", "vue"}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("%v is expected but got %v", expected, r)
	}
}

func TestGetRidOfTheLast(t *testing.T) {
	data := []string{"go", "java", "node", "clojure", "python", "react", "angular", "vue"}

	r := delLast(data)

	expected := []string{"go", "java", "node", "clojure", "python", "react", "angular"}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("%v is expected but got %v", expected, r)
	}
}

func TestGetRidOfTheSecond(t *testing.T) {
	data := []string{"go", "java", "node", "clojure", "python", "react", "angular", "vue"}

	r := delSecond(data)

	expected := []string{"go", "node", "clojure", "python", "react", "angular", "vue"}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("%v is expected but got %v", expected, r)
	}
}

func TestOnlyOddShouldReturn(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	answer := chooseOdd(data)
	expected := []int{1, 3, 5, 7, 9}

	if !reflect.DeepEqual(answer, expected) {
		t.Errorf("%v is expected but got %v", expected, data)
	}
}
