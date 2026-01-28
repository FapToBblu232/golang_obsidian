package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Ban(
	n, p, b, b_max, now int,
	attempts []int,
) int {
	filtered := make([]int, 0, len(attempts))
	for _, t := range attempts {
		if now-t <= 2*b_max {
			filtered = append(filtered, t)
		}
	}
	if len(filtered) < n {
		return 0
	}

	sort.Ints(filtered)

	currentBlock := b
	flag := false
	var end int

	left := 0
	for left < len(filtered)-n+1 {
		for left+n-1 < len(filtered) && filtered[left+n-1]-filtered[left] > p {
			left++
		}

		if left+n-1 >= len(filtered) {
			break
		}

		if flag {
			currentBlock *= 2
			if currentBlock > b_max {
				currentBlock = b_max
			}
		}

		start := filtered[left+n-1]
		end = start + currentBlock
		flag = true

		left += n
	}
	if flag && end > now {
		return end
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n, p, b, b_Max, now int
	var attempts []int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		n, _ = strconv.Atoi(fields[0])
		p, _ = strconv.Atoi(fields[1])
		b, _ = strconv.Atoi(fields[2])
		b_Max, _ = strconv.Atoi(fields[3])
		now, _ = strconv.Atoi(fields[4])
		break
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		temp, _ := strconv.Atoi(line)
		attempts = append(attempts, temp)
	}

	end := Ban(n, p, b, b_Max, now, attempts)

	if end == 0 {
		fmt.Println("ok")
	} else {
		fmt.Println(end)
	}
}

// Время:
// Добавление всех элементов через append - O(n)
// Сортировка целочисленных - O(n*Logn) В Go - стандартной является MergeSort с улучшениями
// Два указателя. Левай не перегонит правый, потому, там не цикл в цикле,
// А левый догоняет правый - O(n)
// Остальное проверки - O(1)
// Итого: Время O(n*logn)

// Память
// Храним все попытки - O(n)
// filtered - ещё O(n)
// Остальное - O(const)
// Итого: Если считать входные данные - O(n),
//        Иначе - можно добавить на каждой итерации проверку и вовсе избавиться от фильтрации в начале
// 		  тогда будет O(1)
