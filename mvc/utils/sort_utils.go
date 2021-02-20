package utils

import "sort"

//Native Sort is slow as compared to below method for elements <1000
func BubbleSort(elements []int) {
	keepLooping := true
	for keepLooping {
		keepLooping = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepLooping = true

			}
		}
	}
}
//Native Sort
func Sort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)

}