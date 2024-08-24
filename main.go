package main

import (
	"errors"
	"fmt"
)

type Deque struct {
	data []int
}

func NewDeque() *Deque {
	return &Deque{
		data: make([]int, 0),
	}
}

func (d *Deque) PushFront(val int) {
	d.data = append([]int{val}, d.data...)
}

func (d *Deque) PushBack(val int) {
	d.data = append(d.data, val)
}

func (d *Deque) PopFront() (int, error) {
	if len(d.data) == 0 {
		return 0, errors.New("empty deque")
	}
	val := d.data[0]
	d.data = d.data[1:]
	return val, nil
}

func (d *Deque) PopBack() (int, error) {
	if len(d.data) == 0 {
		return 0, errors.New("empty deque")
	}
	val := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return val, nil
}

func (d *Deque) Front() (int, error) {
	if len(d.data) == 0 {
		return 0, errors.New("empty deque")
	}
	return d.data[0], nil
}

func (d *Deque) Back() (int, error) {
	if len(d.data) == 0 {
		return 0, errors.New("empty deque")
	}
	return d.data[len(d.data)-1], nil
}

func (d *Deque) Len() int {
	return len(d.data)
}

func (d *Deque) IsEmpty() bool {
	return len(d.data) == 0
}

func main() {
	deque := NewDeque()

	fmt.Println(deque.IsEmpty())

	deque.PushBack(20)
	deque.PushFront(5)
	deque.PushBack(70)

	fmt.Println(deque.IsEmpty())

	fmt.Println(deque.Len())
	fmt.Println(deque.Front())
	fmt.Println(deque.Back())

	fmt.Println(deque.PopFront())
	fmt.Println(deque.PopBack())
	fmt.Println(deque.PopFront())

	fmt.Println(deque.IsEmpty())
}
