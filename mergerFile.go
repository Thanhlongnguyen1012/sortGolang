package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

type Item struct {
	value int
	index int
}
type minHeap []Item

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func mergeFile(inputFile []string, outputFile string) {
	files := make([]*os.File, len(inputFile))
	scanners := make([]*bufio.Scanner, len(inputFile))
	for i, fileName := range inputFile {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		files[i] = file
		scanners[i] = bufio.NewScanner(file)
	}
	fileOut, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer fileOut.Close()
	writer := bufio.NewWriter(fileOut)
	h := &minHeap{}
	heap.Init(h)

	for i, scanner := range scanners {
		if scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			heap.Push(h, Item{value: num, index: i})
		}
	}

	for h.Len() > 0 {
		minItem := heap.Pop(h).(Item)
		writer.WriteString(strconv.Itoa(minItem.value) + "\n")
		if scanners[minItem.index].Scan() {
			num, err := strconv.Atoi(scanners[minItem.index].Text())
			if err != nil {
				panic(err)
			}
			heap.Push(h, Item{value: num, index: minItem.index})
		}
	}
	writer.Flush()
}
