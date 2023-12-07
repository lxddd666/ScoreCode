package test

import (
	"fmt"
	"testing"
)

func TestD(t *testing.T) {
	list := make([]int, 0)
	for i := 1; i <= 900; i++ {
		list = append(list, i)
	}
	slist := splitSlice(list, 1000)

	num := 0
	for _, i := range slist {
		num += len(i)
	}
	fmt.Println(num)
}

func splitSlice(slice []int, chunkSize int) [][]int {
	var result [][]int
	length := len(slice)
	for i := 0; i < length; i += chunkSize {
		end := i + chunkSize
		if end > length {
			end = length
		}
		result = append(result, slice[i:end])
	}
	return result
}
