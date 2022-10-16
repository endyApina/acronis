package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	cpus := runtime.NumCPU()
	fmt.Println("starting program")

	runtime.GOMAXPROCS(cpus)
	var startMemory runtime.MemStats
	runtime.ReadMemStats(&startMemory)

	start := time.Now()

	filePathInArg := os.Args[1]
	if filePathInArg == "" {
		panic(errors.New("empty input arg. please specify input file path"))
	}
	file, err := os.Open(filePathInArg)
	if err != nil {
		fmt.Println("error opening file")
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []int

	for fileScanner.Scan() {
		lineString := fileScanner.Text()
		lineInt, _ := strconv.Atoi(lineString)
		fileLines = append(fileLines, lineInt)
	}

	sortedArray := sortArray(fileLines)
	fmt.Println("//////////SORTED AND WRITING TO OUTPUT/////////")
	writeLine(sortedArray)

	elapsed := time.Since(start)
	var endMemory runtime.MemStats
	runtime.ReadMemStats(&endMemory)

	fmt.Printf("Time elapsed %f \n", elapsed.Seconds())

	fmt.Printf("Memory before %d, memory after %d \n", startMemory.Alloc, endMemory.Alloc)
}

func writeLine(line []int) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range line {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

//sort array using the fast algorithm. Quick Array
func sortArray(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for index, _ := range arr {
		if arr[index] < arr[right] {
			arr[left], arr[index] = arr[index], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	sortArray(arr[:left])
	sortArray(arr[left+1:])

	return arr
}
