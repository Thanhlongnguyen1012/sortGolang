package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// hàm sắp xếp từng file
func sortFile(s1, s2 string) {
	var slice []int
	fileInput, err := os.Open(s1)
	if err != nil {
		panic(err)
	}
	defer fileInput.Close()
	scanner := bufio.NewScanner(fileInput)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		slice = append(slice, num)
	}
	sort.Ints(slice)

	fileOutput, err := os.Create(s2)
	if err != nil {
		panic(err)
	}
	defer fileOutput.Close()
	writer := bufio.NewWriter(fileOutput)
	for _, line := range slice {
		_, err = writer.WriteString(strconv.Itoa(line) + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	fmt.Println("Write done !")
}
