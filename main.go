package main

import (
	"permutations/dictionary"
	"permutations/priorityqueue"
	"fmt"
)

func distance(a string, b string) int {
	d := make([][]int, len(a) + 1)
	for i := range d {
		d[i] = make([]int, len(b) + 1)
		d[i][0] = i
	}

	for j := 0; j <= len(b); j++ {
		d[0][j] = j
	}

	min := func (a int, b int, c int) int {
		if a <= b && a <= c {
			return a
		} else if b <= a && b <= c {
			return b
		} else {
			return c
		}
	}

	for j := 1; j <= len(b); j++ {
		for i := 1; i <= len(a); i++ {
			substCost := 0
			if a[i-1] != b[j-1] {
				substCost = 1
			}

			d[i][j] = min(d[i-1][j] + 1, d[i][j-1] + 1, d[i-1][j-1] + substCost)
		}
	}

	return d[len(a)][len(b)]
}

func permutationsCount(in string, out string) int {
	type item struct {
		value string
		cost int
	}

	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	unqitems := make(map[string]bool)

	pq := priorityqueue.NewQueue()

	addItem := func(v item) {
		if !dictionary.DoesWordExist(v.value) {
			return
		}

		_, exist := unqitems[v.value]
		if !exist {
			pq.Push(v, v.cost + distance(v.value, out))
			unqitems[v.value] = true
		}
	}

	addItem(item{in, 0})

	for pq.Len() != 0 {
		val, _ := pq.Pop()
		it := val.(item)

		fmt.Println(it.value)
		if it.value == out {
			return it.cost
		}

		for i := 0; i < len(it.value); i++ {
			for j := 0; j < len(alphabet); j++ {
				//try to put new char before current position
				temp := it.value[:i] + string(alphabet[j]) + it.value[i:]
				addItem(item{temp, it.cost + 1})

				//try to replace current char
				if alphabet[j] != it.value[i] {
					temp = it.value[:i] + string(alphabet[j]) + it.value[i+1:]
					addItem(item{temp, it.cost + 1})
				}

				if i == len(it.value) - 1 {
					//try to put new char to the end of a string
					temp = it.value + string(alphabet[j])
					addItem(item{temp, it.cost + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	a := "hub"
	b := "mug"
	fmt.Println("Finding amount of permutations between ", a, " and " , b)
	c := permutationsCount(a, b)
	if c == -1 {
		fmt.Println("Can not find a valid path between two words")
	} else {
		fmt.Println("The minimum amount of required premutations is ", c)
	}
}