package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data       []int
	tamanho    int
	capacidade int
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func NewMultiSet(c int) *MultiSet {
	return &MultiSet{
		data:       make([]int, c),
		tamanho:    0,
		capacidade: c,
	}
}

func (ms *MultiSet) String() string {
	return "[" + Join(ms.data[:ms.tamanho], ", ") + "]"
}

func (ms *MultiSet) busca(v int) (bool, int) {
	low := 0
	high := ms.tamanho - 1

	for low <= high {
		mid := (low + high) / 2

		if ms.data[mid] == v {
			return true, mid
		}

		if ms.data[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false, low
}

func (ms *MultiSet) Insert(v int) {
	if ms.tamanho == ms.capacidade {
		novaCap := ms.capacidade * 2

		novoVet := make([]int, novaCap)

		for i := 0; i < ms.tamanho; i++ {
			novoVet[i] = ms.data[i]
		}

		ms.data = novoVet
		ms.capacidade = novaCap
	}

	_, p := ms.busca(v)

	if p > ms.tamanho {
		p = ms.tamanho
	}

	for i := ms.tamanho; i > p; i-- {
		ms.data[i] = ms.data[i-1]
	}

	ms.data[p] = v
	ms.tamanho++
}

func (ms *MultiSet) Contains(valor int) bool {
	tem, _ := ms.busca(valor)
	return tem
}

func (ms *MultiSet) Erase(v int) error {
	err := ms.erase(v)

	if err != nil {
		return err
	}
	return nil
}

func (ms *MultiSet) erase(v int) error {
	for i := 0; i < ms.tamanho; i++ {
		if ms.data[i] == v {
			ms.data = append(ms.data[:i], ms.data[i+1:]...)
			ms.tamanho--
			return nil
		}
	}
	err := errors.New("value not found")
	return err
}

func (ms *MultiSet) Count(v int) int {
	primeiro := ms.LowerBound(v)
	ultimo := ms.UpperBound(v)

	if primeiro == ms.tamanho || ms.data[primeiro] != v {
		return 0
	}

	return ultimo - primeiro + 1
}

func (ms *MultiSet) LowerBound(v int) int {
	low := 0
	high := ms.tamanho - 1

	for low <= high {
		mid := (low + high) / 2

		if ms.data[mid] >= v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func (ms *MultiSet) UpperBound(v int) int {
	low := 0
	high := ms.tamanho - 1

	for low <= high {
		mid := (low + high) / 2

		if ms.data[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

func (ms *MultiSet) Unique() int {
	if ms.tamanho == 0 {
		return 0
	}

	unicos := 1

	for i := 1; i < ms.tamanho; i++ {
		if ms.data[i] != ms.data[i-1] {
			unicos++
		}
	}

	return unicos
}

func (ms *MultiSet) Clear() {
	ms.data = nil
	ms.tamanho = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
