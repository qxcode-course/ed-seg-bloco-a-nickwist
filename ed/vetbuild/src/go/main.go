package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
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

func (v *Vector) String() string {
	return "[" + Join(v.data[:v.Size()], ", ") + "]"
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.Size(), v.Capacity())
}

func (v *Vector) PushBack(valor int) {
	if v.Size() == v.Capacity() {

		capacityNova := v.Capacity() * 2

		novoVetor := make([]int, capacityNova)

		for i := 0; i < v.Size(); i++ {
			novoVetor[i] = v.data[i]
		}

		v.data = novoVetor
		v.capacity = capacityNova

	}

	v.data[v.Size()] = valor
	v.size++
}

func (v *Vector) Get(p int) int {
	return v.data[p]
}

func (v *Vector) At(p int) (int, error) {
	if p < 0 || p >= v.Size() {
		return 0, errors.New("index out of range")
	}

	return v.Get(p), nil
}

func (v *Vector) Set(p, valor int) error {
	_, err := v.At(p)

	if err != nil {
		return err
	}

	v.data[p] = valor

	return nil
}

func (v *Vector) Clear() {
	v.data = nil
	v.size = 0
}

func (v *Vector) Reserve(c int) {
	v.capacity = c
}

func (v *Vector) PopBack() error {

	if v.Size() == 0 {
		return errors.New("vector is empty")
	}

	v.size--
	return nil
}

func (v *Vector) Insert(p, valor int) error {
	if p < 0 || p > v.size {
		return errors.New("index out of range")
	}

	if v.Size() == v.Capacity() {
		newCap := v.capacity * 2
		novo := make([]int, newCap)

		for i := 0; i < v.size; i++ {
			novo[i] = v.data[i]
		}

		v.data = novo
		v.capacity = newCap
	}
	for i := v.size; i > p; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[p] = valor
	v.size++

	return nil
}

func (v *Vector) Erase(p int) error {
	if p < 0 || p > v.size {
		return errors.New("index out of range")
	}

	v.data = append(v.data[:p], v.data[p+1:]...)
	v.size--

	return nil
}

func (v *Vector) IndexOf(valor int) int {
	for i := 0; i < v.Size(); i++ {
		if v.data[i] == valor {
			return i
		}
	}

	return -1
}

func (v *Vector) Contains(valor int) bool {
	for i := 0; i < v.Size(); i++ {
		if v.data[i] == valor {
			return true
		}
	}

	return false
}

func (v *Vector) Slice(inicio, fim int) *Vector {
	result := NewVector(v.Capacity())

	n := v.size

	for inicio < 0 {
		inicio += n
	}
	for fim < 0 {
		fim += n
	}

	inicio = inicio % n
	fim = fim % n

	// caso normal
	for i := inicio; i != fim; i = (i + 1) % n {
		result.PushBack(v.data[i])
	}

	return result
}

func (v *Vector) Size() int {
	return v.size
}

func (v *Vector) Capacity() int {
	return v.capacity
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
