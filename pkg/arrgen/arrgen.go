package arrgen

import (
	"math/rand"
	"time"
)

var DefaultArraySizes = []int{10, 25, 50, 100, 250, 500, 1000, 2000, 3000}
var Limit = 1000

func randomInt() (num int) {
	rand.Seed(time.Now().UnixNano())
	num = rand.Intn(Limit)
	return num
}

func Generate(n int) (arr []*int, err error) {
	arr = make([]*int, n)
	for i := 0; i < n; i++ {
		num := randomInt()
		arr[i] = &num
	}
	return arr, err
}

func CopyArr(arr []*int) (newArr []*int) {
	newArr = make([]*int, len(arr))
	for i := 0; i < len(arr); i++ {
		num := *arr[i]
		newArr[i] = &num
	}
	return newArr
}
