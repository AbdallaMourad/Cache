package LFU_test

import (
	"fmt"
	"std/github.com/AbdallaMourad/Cache/src/LFU"
)

type Command struct {
	method          string
	data            []int
	expected_result interface{}
}

func Example() {
	lru_cache := LFU.Constructor(3)
	test_cases := []Command{
		{"put", []int{2, 2}, nil},
		{"put", []int{1, 1}, nil},
		{"get", []int{2}, 2},
		{"get", []int{1}, 1},
		{"get", []int{2}, 2},
		{"put", []int{3, 3}, nil},
		{"put", []int{4, 4}, nil},
		{"get", []int{3}, -1},
		{"get", []int{2}, 2},
		{"get", []int{1}, 1},
		{"get", []int{4}, 4},
	}

	results := make([]interface{}, len(test_cases))

	for i, test_case := range test_cases {
		if test_case.method == "put" {
			lru_cache.Put(test_case.data[0], test_case.data[1])
		} else {
			results[i] = lru_cache.Get(test_case.data[0])
		}
	}

	fmt.Println(results)

	// Output:
	// [<nil> <nil> 2 1 2 <nil> <nil> -1 2 1 4]
}
