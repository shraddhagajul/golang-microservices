package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T){
	els := []int{9,8,7,6,5}
	Sort(els)
	assert.EqualValues(t,5,len(els))
	assert.EqualValues(t,5,els[0])
	assert.EqualValues(t,6,els[1])
	assert.EqualValues(t,7,els[2])
	assert.EqualValues(t,8,els[3])
	assert.EqualValues(t,9,els[4])
}

func TestBubbleSortBestCase(t *testing.T){
	els := []int{5,6,7,8,9}
	Sort(els)
	assert.EqualValues(t,5,len(els))
	assert.EqualValues(t,5,els[0])
	assert.EqualValues(t,6,els[1])
	assert.EqualValues(t,7,els[2])
	assert.EqualValues(t,8,els[3])
	assert.EqualValues(t,9,els[4])
}

func TestBubbleSortNil(t *testing.T){
	Sort(nil)
}

func getElements(n int) []int{
	result := make([]int, n)
	i := 0
	for j:= n-1; j>=0; j-- {
		result[i] = j 
		i++
	}
	return result
}

func BenchmarkBubbleSort10(b *testing.B){
	els := getElements(10)
	for i:=0;i<b.N;i++{
		Sort(els)
	}

}

// func BenchmarkSort10(b *testing.B){
// 	els := getElements(10)
// 	for i:=0;i<b.N;i++{
// 		Sort(els)
// 	}

// }

func BenchmarkBubbleSort1000(b *testing.B){
	els := getElements(1000)
	for i:=0;i<b.N;i++{
		Sort(els)
	}

}

// func BenchmarkSort1000(b *testing.B){
// 	els := getElements(1000)
// 	for i:=0;i<b.N;i++{
// 		Sort(els)
// 	}

// }


func BenchmarkBubbleSort100000(b *testing.B){
	els := getElements(100000)
	for i:=0;i<b.N;i++{
		Sort(els)
	}

}

// func BenchmarkSort100000(b *testing.B){
// 	els := getElements(100000)
// 	for i:=0;i<b.N;i++{
// 		Sort(els)
// 	}

// }
