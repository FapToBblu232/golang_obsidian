package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data     []string
	len      int
	capacity int
	first    int //индекс первого
	last     int // индекс последнего элемента
}

func (deque *Deque) pushb(input string) {
	if deque.len == deque.capacity {
		fmt.Println("overflow")
		return
	}
	deque.data[deque.last] = input
	deque.last = (deque.last + 1) % deque.capacity
	deque.len++
}

func (deque *Deque) pushf(input string) {
	if deque.len == deque.capacity {
		fmt.Println("overflow")
		return
	}
	deque.first = (deque.first + deque.capacity - 1) % deque.capacity
	deque.data[deque.first] = input
	deque.len++
}

func (deque *Deque) popf() string {
	if deque.len == 0 {
		return "underflow"
	}
	answ := deque.data[deque.first]
	deque.first = (deque.first + 1) % deque.capacity
	deque.len--
	return answ
}

func (deque *Deque) popb() string {
	if deque.len == 0 {
		return "underflow"
	}
	deque.last = (deque.last + deque.capacity - 1) % deque.capacity
	answ := deque.data[deque.last]
	deque.len--
	return answ
}

func (deque *Deque) print() {
	if deque.len == 0 {
		fmt.Println("empty")
		return
	}
	for i := 0; i < deque.len; i++ {
		ind := (deque.first + i) % deque.capacity
		fmt.Printf("%v ", deque.data[ind])
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var deque Deque
	setted := false
	for scanner.Scan() {
		request := scanner.Text()
		SPACE := 0
		for i := range request {
			if string(request[i]) == " " {
				SPACE++
			}
		}

		if request == "" {
			continue
		}
		fields := strings.Fields(request)

		if !setted {
			if len(fields) != 2 || fields[0] != "set_size" {
				fmt.Println("error")
				continue
			}
			size, _ := strconv.Atoi(fields[1])
			if size < 0 {
				fmt.Println("error")
				continue
			}
			deque = Deque{
				data:     make([]string, size),
				len:      0,
				capacity: size,
				first:    0,
				last:     0,
			}
			setted = true
			continue
		}

		switch fields[0] {
		case "pushb":
			if len(fields) != 2 {
				fmt.Println("error")
				continue
			}
			input := fields[1]
			deque.pushb(input)

		case "pushf":
			if len(fields) != 2 {
				fmt.Println("error")
				continue
			}
			input := fields[1]
			deque.pushf(input)
		case "popf":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}
			if SPACE > 0 {
				fmt.Println("error")
				continue
			}
			answ := deque.popf()
			fmt.Println(answ)

		case "popb":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}
			if SPACE > 0 {
				fmt.Println("error")
				continue
			}
			answ := deque.popb()
			fmt.Println(answ)

		case "print":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}
			if SPACE > 0 {
				fmt.Println("error")
				continue
			}
			deque.print()

		default:
			fmt.Println("error")
		}
	}
}
