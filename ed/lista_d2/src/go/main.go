package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	valor int
	prox  *Node
	ant   *Node
}

type LList struct {
	root *Node
}

func NewLList() *LList {
	root := &Node{}

	root.prox = root
	root.ant = root

	return &LList{
		root: root,
	}
}

func (l *LList) PushFront(valor []int) {
	for _, valor := range valor {

		v := &Node{
			valor: valor,
		}

		primeiro := l.root.prox

		v.prox = primeiro
		v.ant = l.root

		primeiro.ant = v
		l.root.prox = v
	}
}

func (l *LList) String() string {
	if l.root.prox == l.root {
		return "[]"
	}

	s := "["

	for node := l.root.prox; node != l.root; node = node.prox {
		s += strconv.Itoa(node.valor)

		if node.prox != l.root {
			s += ", "
		}
	}

	s += "]"
	return s
}

func (n *Node) Next() *Node {
	return n.prox
}

func (n *Node) Prev() *Node {
	return n.ant
}

func (l *LList) Front() *Node {
	return l.root.prox
}

func (l *LList) Back() *Node {
	return l.root.ant
}

func (l *LList) Size() int {
	quantidade := 0
	node := l.root.prox
	for node != l.root {
		quantidade++
		node = node.prox
	}
	return quantidade
}

func (l *LList) Clear() {
	l.root.ant = l.root
	l.root.prox = l.root
}

func (l *LList) PushBack(valor []int) {
	for i := len(valor) - 1; i >= 0; i-- {
		v := &Node{
			valor: valor[i],
		}

		primeiro := l.root.prox

		v.prox = primeiro
		v.ant = l.root

		primeiro.ant = v
		l.root.prox = v
	}
}

func (l *LList) Search(valor int) *Node {
	for node := l.root.prox; node != l.root; node = node.prox {
		if node.valor == valor {
			return node
		}
	}
	return nil

}

func (l *LList) Insert(node *Node, valor int) {
	if node == nil || node == l.root {
		return
	}

	novo := &Node{valor: valor}

	anterior := node.ant

	novo.prox = node
	novo.ant = anterior

	anterior.prox = novo
	node.ant = novo
}

func (l *LList) Remove(node *Node) {
	if node == nil || node == l.root {
		return
	}

	anterior := node.ant
	proximo := node.prox

	anterior.prox = proximo
	proximo.ant = anterior
}

func (l *LList) PopFront() {
	if l.root.ant == l.root {
		return
	}

	node := l.root.prox

	node.ant.prox = node.prox
	node.prox.ant = node.ant
}

func (l *LList) PopBack() {
	if l.root.ant == l.root {
		return
	}

	node := l.root.ant

	node.ant.prox = node.prox
	node.prox.ant = node.ant
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			var nums []int

			for _, v := range args[1:] {
				num, err := strconv.Atoi(v)
				if err != nil {
					return
				}

				nums = append(nums, num)
			}

			ll.PushBack(nums)
		case "push_front":
			var nums []int

			for _, v := range args[1:] {
				num, err := strconv.Atoi(v)
				if err != nil {
					return
				}

				nums = append(nums, num)
			}

			ll.PushFront(nums)
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != ll.root; node = node.Next() {
				fmt.Printf("%v ", node.valor)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != ll.root; node = node.Prev() {
				fmt.Printf("%v ", node.valor)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.valor = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}

	}
}
