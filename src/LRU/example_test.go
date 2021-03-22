package LRU

import "fmt"

type Command struct {
	method          string
	data            []int
	expected_result interface{}
}

func Example() {
	lru_cache := Constructor(2)
	test_cases := []Command{
		{"put", []int{1, 1}, nil},
		{"put", []int{2, 2}, nil},
		{"get", []int{1}, 1},
		{"put", []int{3, 3}, nil},
		{"get", []int{2}, -1},
		{"put", []int{4, 4}, nil},
		{"get", []int{1}, -1},
		{"get", []int{3}, 3},
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
	// [<nil> <nil> 1 <nil> -1 <nil> -1 3 4]
}
