package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data       []int
	tamanho    int
	capacidade int
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func NewSet(c int) *Set {
	return &Set{
		data:       make([]int, c), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		tamanho:    0,
		capacidade: c,
	}
}

func (s *Set) String() string {
	return "[" + Join(s.data[:s.tamanho], ", ") + "]"
}

func (s *Set) Insert(v int) {
	for i := 0; i < s.tamanho; i++ {
		if s.data[i] == v {
			return
		}
	}

	if s.tamanho == s.capacidade {
		novaCapacidade := s.capacidade * 2
		if novaCapacidade == 0 {
			novaCapacidade = 1
		}

		novoVetor := make([]int, novaCapacidade)
		copy(novoVetor, s.data)

		s.data = novoVetor
		s.capacidade = novaCapacidade
	}

	pos := 0
	for pos < s.tamanho && s.data[pos] < v {
		pos++
	}

	for i := s.tamanho; i > pos; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[pos] = v
	s.tamanho++
}

func (s *Set) Contains(valor int) bool {
	for i := 0; i < s.tamanho; i++ {
		if s.data[i] == valor {
			return true
		}
	}

	return false
}

func (s *Set) Erase(v int) error {

	for i := 0; i < s.tamanho; i++ {
		if s.data[i] == v {
			s.data = append(s.data[:i], s.data[i+1:]...)
			s.tamanho--
			return nil
		}
	}

	return errors.New("value not found")
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			s = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				s.Insert(value)
			}
		case "show":
			fmt.Println(s)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := s.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(s.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
