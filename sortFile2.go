package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

// readNumbers đọc các số nguyên từ file và gửi vào channel
func readNumbers(filePath string, out chan<- int, wg *sync.WaitGroup) {
	defer close(out) // Đóng channel sau khi đọc xong
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Lỗi mở file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024*1024), 10*1024*1024) // Bộ đệm 10MB

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Lỗi chuyển đổi số: %v\n", err)
			continue
		}
		out <- num
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Lỗi khi đọc file: %v\n", err)
	}
}

// worker nhận số từ channel, lưu vào slice, sắp xếp và gửi kết quả về channel
func worker(id int, in <-chan int, out chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()

	var numbers []int
	for num := range in {
		numbers = append(numbers, num)
	}

	sort.Ints(numbers) // Sắp xếp phần dữ liệu

	out <- numbers // Gửi kết quả ra channel
}

// mergeSortedLists trộn danh sách đã sắp xếp thành một danh sách duy nhất
func mergeSortedLists(chunks [][]int) []int {
	if len(chunks) == 0 {
		return nil
	}

	result := chunks[0]
	for i := 1; i < len(chunks); i++ {
		result = mergeTwoLists(result, chunks[i])
	}
	return result
}

// mergeTwoLists trộn hai danh sách đã sắp xếp
func mergeTwoLists(a, b []int) []int {
	i, j := 0, 0
	merged := make([]int, 0, len(a)+len(b))

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	merged = append(merged, a[i:]...)
	merged = append(merged, b[j:]...)

	return merged
}

// writeNumbers ghi danh sách số vào file
func writeNumbers(filePath string, numbers []int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("không thể tạo file %s: %v", filePath, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		fmt.Fprintln(writer, num)
	}

	return writer.Flush()
}

// sortFile xử lý toàn bộ quá trình đọc, sắp xếp và ghi file
func sortFile(inputFile, outputFile string) error {
	numWorkers := 4 // Số worker (tùy chỉnh theo CPU)

	// Channel để truyền dữ liệu giữa các goroutine
	numChan := make(chan int, 10000)
	sortedChan := make(chan []int, numWorkers)

	var wg sync.WaitGroup

	// Goroutine đọc file
	wg.Add(1)
	go readNumbers(inputFile, numChan, &wg)

	// Goroutine xử lý sắp xếp
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, numChan, sortedChan, &wg)
	}

	// Đợi tất cả worker hoàn thành
	go func() {
		wg.Wait()
		close(sortedChan) // Đóng channel khi tất cả worker hoàn thành
	}()

	// Thu thập và trộn dữ liệu đã sắp xếp từ các worker
	var sortedChunks [][]int
	for sortedList := range sortedChan {
		sortedChunks = append(sortedChunks, sortedList)
	}

	// Trộn tất cả dữ liệu đã sắp xếp
	finalSorted := mergeSortedLists(sortedChunks)

	// Ghi kết quả ra file
	if err := writeNumbers(outputFile, finalSorted); err != nil {
		return err
	}

	fmt.Println("Write done!")
	return nil
}

func main() {
	if err := sortFile("input.txt", "output.txt"); err != nil {
		fmt.Println("Lỗi:", err)
	}
}
