package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Item struct {
	Next  *Item
	Value string
}

func (i *Item) Print() {
	for i != nil {
		fmt.Println(i.Value)
		i = i.Next
	}
}

func main() {
	/*
		if len(os.Args) != 2 {
			fmt.Println("args not enough: need 2, catched:", len(os.Args))
			return
		}
		f, err := os.Open(os.Args[1])
	*/
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("file open fail:", err.Error())
		return
	}

	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println("file close fail:", err.Error())
		}
	}()

	r := bufio.NewReader(f)
	var it *Item
	for {
		val, err := r.ReadString('\n')
		if len(val) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read file fail:", err.Error())
			return
		}
		val = strings.TrimSuffix(val, "\n")

		newItem := &Item{
			Value: val,
			Next:  it,
		}
		it = newItem
	}

	fmt.Println("order after read:")
	it.Print()

	fmt.Println("===")

	var first, middle, third *Item
	first = it
	if it != nil {
		middle = it.Next
	}
	if middle != nil {
		third = it.Next.Next
	}
	first.Next = nil

	for third != nil {
		middle.Next = first
		first, middle, third = middle, third, third.Next
	}
	middle.Next = first

	fmt.Println("order after reverse:")
	middle.Print()
}
