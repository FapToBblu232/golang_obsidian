package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type BitSet struct {
	data []uint64
	size int
}

func NewBitSet(size int) *BitSet {
	return &BitSet{
		data: make([]uint64, (size+63)/64),
		size: size,
	}
}

func (bs *BitSet) Get(pos int) bool {
	return bs.data[pos/64]&(1<<(pos%64)) != 0
}

func (bs *BitSet) Set(pos int) {
	bs.data[pos/64] |= 1 << (pos % 64)
}

func (bs *BitSet) Print() {
	for i := 0; i < bs.size; i++ {
		if bs.Get(i) {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
	fmt.Println()
}

type bloom struct {
	bitset     *BitSet
	size       int
	hasher_num int
	primes     []int
	mersenne31 uint64
}

//вспомогательные

func (b *bloom) getPrimes(count int) {
	b.primes = []int{}
	num := 2
	for len(b.primes) < count {
		if isPrime(num) {
			b.primes = append(b.primes, num)
		}
		num++
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// func (b *bloom) getBit(pos int) bool {
// 	return b.bits[pos/64]&(1<<(pos%64)) != 0
// }

// func (b *bloom) setBit(pos int) {
// 	b.bits[pos/64] |= 1 << (pos % 64)
// }

func (b *bloom) hash(key uint64, i int) int {
	return int(((uint64(i+1)*(key%b.mersenne31) + uint64(b.primes[i])) % b.mersenne31) % uint64(b.size))
}

// функции для пользования

func (b *bloom) Set(n int64, P float64) (int, int, error) {
	// выводим размер (m) и кол-во хэшеров (k)
	if n <= 0 || P <= 0 || P >= 1 {
		return 0, 0, errors.New("error")
	}
	// формулы
	b.size = int(math.Round(-float64(n) * math.Log2(P) / math.Log(2)))
	b.bitset = NewBitSet(b.size) // для округления вверх
	b.hasher_num = int(math.Round(-math.Log2(P)))
	if b.hasher_num == 0 {
		return 0, 0, errors.New("error")
	}
	b.mersenne31 = (1 << 31) - 1
	b.getPrimes(b.hasher_num)
	return b.size, b.hasher_num, nil
}

func (b *bloom) Has(key uint64) bool {
	for i := 0; i < b.hasher_num; i++ {
		pos := b.hash(key, i)
		if !b.bitset.Get(pos) {
			return false
		}
	}
	return true
}

func (b *bloom) Add(key uint64) {
	for i := 0; i < b.hasher_num; i++ {
		pos := b.hash(key, i)
		b.bitset.Set(pos)
	}
}

func (b *bloom) Print() {
	b.bitset.Print()
}

func main() {
	var bf bloom
	scanner := bufio.NewScanner(os.Stdin)
	setted := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		firstmas := fields[0]

		if firstmas == "set" {
			if len(fields) != 3 {
				fmt.Println("error")
				continue
			}
			n, err1 := strconv.ParseInt(fields[1], 10, 64)
			p, err2 := strconv.ParseFloat(fields[2], 64)
			if err1 != nil || err2 != nil {
				fmt.Println("error")
				continue
			}
			numBits, numHashes, err := bf.Set(n, p)
			if err != nil {
				fmt.Println("error")
				continue
			}
			fmt.Printf("%d %d\n", numBits, numHashes)
			setted = true
			break
		} else {
			fmt.Println("error")
		}
	}

	if !setted {
		// error уже вывели
		return
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		command := fields[0]

		switch command {
		case "print":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}
			bf.Print()
		case "add":
			if len(fields) != 2 {
				fmt.Println("error")
				continue
			}
			key, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				fmt.Println("error")
				continue
			}
			bf.Add(key)
		case "search":
			if len(fields) != 2 {
				fmt.Println("error")
				continue
			}
			key, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				fmt.Println("error")
				continue
			}
			if bf.Has(key) {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		default:
			fmt.Println("error")
		}
	}
}
