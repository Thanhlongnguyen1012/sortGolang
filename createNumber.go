package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func creaFile(s string) {
	var count int = 200_000_000
	var maxValue uint64 = 18_000_000_000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	file, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < count; i++ {
		num := r.Uint64() % (maxValue + 1)
		_, err = writer.WriteString(strconv.FormatUint(num, 10) + "\n")
	}
	if err != nil {
		panic(err)
	}
	writer.Flush()
	fmt.Println("done!")
}
