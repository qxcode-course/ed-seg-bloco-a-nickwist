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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
