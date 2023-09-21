package insertion

import "msis/pkg/sort"

type client struct {
}

func New() sort.Service {
	return &client{}
}

func (c *client) Sort(n int, arr []*int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && *arr[j] > *key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}
