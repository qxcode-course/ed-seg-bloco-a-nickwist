package main

import (
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	txt := ""
		txt += "[ "

		node := l.Front()

		for i := 0; i < l.Size(); i++ {
			if node == sword {
				txt += fmt.Sprintf("%d> ", node.Value)
			} else {
				txt += fmt.Sprintf("%d ", node.Value)
			}

		node = Next(l, node)	
		}
	txt += "]"

	return txt
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it == nil || l.Size() == 0 {
		return nil
	}

	if it.next == it.root {
		return l.Front()
	}

	return it.next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
