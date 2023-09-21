package merge

import (
	"fmt"
	"testing"
)

func TestClient_Sort(t *testing.T) {
	ms := New()
	a := []int{200, 10, 100, 0, 15}
	var i int
	ptr := make([]*int, len(a))

	for i = 0; i < 5; i++ {
		ptr[i] = &a[i]
	}

	ms.Sort(5, ptr)

	for i = 0; i < 4; i++ {
		if *ptr[i] > *ptr[i+1] {
			for j := 0; j < 5; j++ {
				fmt.Printf("%d ", *ptr[j])
			}
			t.Fatal("Array not sorted.")
		}
	}
}
