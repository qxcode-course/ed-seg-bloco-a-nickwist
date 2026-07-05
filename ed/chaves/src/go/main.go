package main

import (
	"fmt"
)

func main() {
	q := NewQueue[rune]()
	var mm, nn int

	for i := 'A'; i <= 'P'; i++ {
		q.Enqueue(i)
	}

	for q.items.Len() != 1 {
		primeiro := q.Dequeue()
		segundo := q.Dequeue()

		fmt.Scan(&mm, &nn)

		if mm > nn {
			q.Enqueue(primeiro)
		} else if nn > mm {
			q.Enqueue(segundo)
		}

	}

	fmt.Printf("%c\n", q.Dequeue())
}
