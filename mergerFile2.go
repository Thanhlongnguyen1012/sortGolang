package main

/*import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"sync"
)

type Item struct {
	value int
	index int
}
type minHeap []Item

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Goroutine đọc từng file và gửi dữ liệu qua channel
func readFile(fileName string, index int, out chan<- Item, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		out <- Item{value: num, index: index}
	}
}

func mergeFile(inputFiles []string, outputFile string) {
	// Channel nhận dữ liệu từ goroutines
	dataCh := make(chan Item, 1000)
	var wg sync.WaitGroup

	// Khởi chạy goroutines để đọc file song song
	for i, file := range inputFiles {
		wg.Add(1)
		go readFile(file, i, dataCh, &wg)
	}

	// Goroutine đóng channel sau khi đọc xong
	go func() {
		wg.Wait()
		close(dataCh)
	}()

	// Mở file output
	fileOut, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer fileOut.Close()
	writer := bufio.NewWriter(fileOut)

	// Min-heap để sắp xếp dữ liệu
	h := &minHeap{}
	heap.Init(h)

	// Thu thập dữ liệu từ channel và push vào heap
	for item := range dataCh {
		heap.Push(h, item)
	}

	// Xuất kết quả từ heap ra file
	for h.Len() > 0 {
		minItem := heap.Pop(h).(Item)
		writer.WriteString(strconv.Itoa(minItem.value) + "\n")
	}
	writer.Flush()
}
*/
