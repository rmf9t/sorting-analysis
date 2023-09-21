package main

import (
	"fmt"
	"msis/pkg/arrgen"
	"msis/pkg/factory/sort"
	"msis/pkg/models"
	"msis/pkg/plot"
	"msis/pkg/utils"
	"os"
	"strconv"
	"time"
)

func main() {
	Start()
}

func Start() {
	fmt.Println("Merge sort and Insertion sort analysis")
	PrintSortOption()

	var algo string
	fmt.Scanln(&algo)

	algoType, err := strconv.Atoi(algo)
	if err != nil {
		fmt.Printf("ERROR: Unable to convert algo type for %s. Because: %s", algo, err.Error())
		os.Exit(1)
	}

	utils.ClearScreen()

	fmt.Println("Scanned algo is:", algoType)

	if algoType == 3 {
		ms, err := sort.NewSortService(sort.Merge)
		if err != nil {
			fmt.Printf("ERROR: Unable to create a new merge sort service for algo type %d. Because: %s", algoType, err.Error())
			os.Exit(1)
		}

		ins, err := sort.NewSortService(sort.Insertion)
		if err != nil {
			fmt.Printf("ERROR: Unable to create a new insertion sort service for algo type %d. Because: %s", algoType, err.Error())
			os.Exit(1)
		}

		msTimes := []models.SortTime{}
		insTimes := []models.SortTime{}

		for _, arrSize := range arrgen.DefaultArraySizes {
			arr, err := arrgen.Generate(arrSize)
			if err != nil {
				fmt.Printf("ERROR: unable to generate array for size %d. Because: %s", arrSize, err.Error())
				os.Exit(1)
			}

			msArr := arrgen.CopyArr(arr)
			startMs := time.Now()
			ms.Sort(arrSize, msArr)
			endMs := time.Now()
			fmt.Printf("Time took for merge sort the list of size %d is: %d\n", arrSize, endMs.Sub(startMs).Nanoseconds())
			msTimes = append(msTimes, models.SortTime{
				Size:      arrSize,
				TimeTaken: endMs.Sub(startMs).Nanoseconds(),
			})

			insArr := arrgen.CopyArr(arr)

			startIns := time.Now()
			ins.Sort(arrSize, insArr)
			endIns := time.Now()
			fmt.Printf("Time took for insertion sort the list of size %d is: %d\n", arrSize, endIns.Sub(startIns).Nanoseconds())
			insTimes = append(insTimes, models.SortTime{
				Size:      arrSize,
				TimeTaken: endIns.Sub(startIns).Nanoseconds(),
			})

			fmt.Println()
		}

		filename := "points.png"
		fmt.Println("INFO: Plotting graph ...")
		if err := plot.PlotGraph(insTimes, msTimes, filename); err != nil {
			fmt.Println("ERROR: Unable to plot graph. Because:", err.Error())
			os.Exit(1)
		}

		fmt.Println("INFO: Graph plotted and saved as", filename)

	} else {
		PrintNSize()

		var n string
		fmt.Scanln(&n)

		arrSize, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("ERROR: Unable to convert algo type for %s. Because: %s", n, err.Error())
			os.Exit(1)
		}

		fmt.Println("Scanned n is:", arrSize)

		ss, err := sort.NewSortService(sort.CheckAlgo(algoType))
		if err != nil {
			fmt.Printf("ERROR: Unable to create a new sort service for algo type %d. Because: %s", algoType, err.Error())
			os.Exit(1)
		}

		arr, err := arrgen.Generate(arrSize)
		if err != nil {
			fmt.Printf("ERROR: unable to generate array for size %d. Because: %s", arrSize, err.Error())
			os.Exit(1)
		}

		start := time.Now()
		ss.Sort(arrSize, arr)
		end := time.Now()

		fmt.Println("Time took to sort the list is:", end.Sub(start).Nanoseconds())
	}
}

func PrintSortOption() {
	fmt.Println("Choose Sorting Algorithm:")
	fmt.Println("1. Merge Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Println("3. Compare merge and insertion sort")
}

func PrintNSize() {
	fmt.Println("Choose the size of the array from the list [10, 25, 50 100, 250 and 500, 1000, 2000]:")
}
