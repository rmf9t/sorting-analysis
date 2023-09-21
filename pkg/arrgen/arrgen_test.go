package arrgen

import (
	"fmt"
	"testing"
)

func TestClient_Sort(t *testing.T) {
	for _, num := range DefaultArraySizes {
		arr, _ := Generate(num)
		for i := 0; i < num; i++ {
			fmt.Printf("%d ", *arr[i])
		}
		fmt.Println()
	}
}
