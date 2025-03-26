package main

import (

	//"math/rand"
	//"os"
	//"strconv"

	"fmt"
	"sync"
	"time"
)

func main() {
	//start := time.Now()
	//create file
	/*a := []string{"number1.txt", "number2.txt", "number3.txt", "number4.txt", "number5.txt", "number6.txt", "number7.txt", "number8.txt", "number9.txt", "number10.txt"}
	var wg sync.WaitGroup
	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			creaFile(s)
		}(a[i])
	}
	wg.Wait()

	for i := 0; i < len(a); i++ {
		creaFile(a[i])
	}
	delta := time.Now().Sub(start)
	fmt.Println("time: ", delta)
	*/

	start := time.Now()
	//create file
	a := []string{"number1.txt", "number2.txt", "number3.txt", "number4.txt", "number5.txt", "number6.txt", "number7.txt", "number8.txt", "number9.txt", "number10.txt", "number11.txt", "number12.txt", "number13.txt", "number14.txt", "number15.txt", "number16.txt", "number17.txt", "number18.txt", "number19.txt", "number20.txt"}
	b := []string{"number1_Sort.txt", "number2_Sort.txt", "number3_Sort.txt", "number4_Sort.txt", "number5_Sort.txt", "number6_Sort.txt", "number7_Sort.txt", "number8_Sort.txt", "number9_Sort.txt", "number10_Sort.txt", "number11_Sort.txt", "number12_Sort.txt", "number13_Sort.txt", "number14_Sort.txt", "number15_Sort.txt", "number16_Sort.txt", "number17_Sort.txt", "number18_Sort.txt", "number19_Sort.txt", "number20_Sort.txt"}
	var wg sync.WaitGroup
	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			creaFile(s)
		}(a[i])
	}
	wg.Wait()
	for i := 0; i < len(a); i++ {
		sortFile(a[i], b[i])
	}
	mergeFile(b, "outPut.txt")
	// for i := 0; i < len(a); i++ {
	// 	creaFile(a[i])
	// }
	// for i := 0; i < len(a); i++ {
	// 	wg.Add(1)
	// 	go func(s1, s2 string) {
	// 		defer wg.Done()
	// 		sortFile(s1, s2)
	// 	}(a[i], b[i])
	// }

	// wg.Wait()
	// var wg sync.WaitGroup

	// for i := 0; i < len(b); i++ {
	// 	wg.Add(1)
	// 	go func(s string) {
	// 		defer wg.Done()
	// 		testOutput(s)
	// 	}(b[i])
	// }
	delta := time.Now().Sub(start)
	fmt.Println("time: ", delta)
}
