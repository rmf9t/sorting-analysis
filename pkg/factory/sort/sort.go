package sort

import (
	"errors"
	"msis/pkg/sort"
	"msis/pkg/sort/insertion"
	"msis/pkg/sort/merge"
)

type AlgoType int

const (
	Merge     AlgoType = 0
	Insertion AlgoType = 1
)

var UndefinedAlgoType = errors.New("undefined sort algorithm type")

// NewSortService adapts Factory pattern to produce a sort service based on the input type
func NewSortService(sortType AlgoType) (sort.Service, error) {
	switch sortType {
	case Merge:
		return merge.New(), nil
	case Insertion:
		return insertion.New(), nil
	}
	return nil, UndefinedAlgoType
}

func CheckAlgo(algo int) AlgoType {
	switch algo {
	case 1:
		return Merge
	case 2:
		return Insertion
	}
	return -1
}
