package main

// https://discopal.ispras.ru/img_auth.php/9/98/Ptas-knapsack.beam.pdf
// алгоритм на основе этой статьи + то, что было на лекции про восходящую динамику
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// пришлось создать, т.к. счиатть вес + цену каждый раз при получении набора - сложно
type Set struct {
	weight int
	value  int
	idexes []int
}

func accurate(values, weights []int, maxi int) Set {
	best_for_price := map[int]Set{0: {}} //для каждой цены - индексы лучшего набора
	for i := 0; i < len(values); i++ {
		var answ []Set
		for _, item := range best_for_price { // к кажому лучшему набору попробуем добавить предмет
			cur_value := values[i] + item.value
			cur_weight := weights[i] + item.weight
			if cur_weight <= maxi {
				prev_set, ok := best_for_price[cur_value]
				if !ok || cur_weight < prev_set.weight { // Если цена впервые или вес лучше, то запомним этот набор
					cur_set := Set{
						cur_weight,
						cur_value,
						append(append([]int{}, item.idexes...), i),
					}
					answ = append(answ, cur_set)
				}
			}
		}
		for _, set := range answ {
			best_for_price[set.value] = set
		}
	}
	maxVal := 0
	for _, set := range best_for_price {
		if set.value > maxVal {
			maxVal = set.value
		}
	}
	return best_for_price[maxVal]
}

func rykzak(precision float64, values, weights []int, maxi int) Set {
	maxPrice := 0 // для ошибки берём макс значение, от него уже процент
	for _, val := range values {
		if val > maxPrice {
			maxPrice = val
		}
	}
	scale := precision * float64(maxPrice) / ((1 + precision) * (float64(len(values))))
	scaledPrices := make([]int, len(values))

	for i, val := range values {
		scaledPrices[i] = int(float64(val) / scale)
	}
	answer := accurate(scaledPrices, weights, maxi)

	normalPrice := 0
	for _, ind := range answer.idexes {
		normalPrice += values[ind]
	}
	answer.value = normalPrice
	return answer
}

func main() {
	var precision float64
	var maxi int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		temp, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("error")
			continue
		}
		if temp < 0 || temp > 1 {
			fmt.Println("error")
			continue
		}
		precision = temp
		break
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		temp, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("error")
			continue
		}
		if temp < 0 {
			fmt.Println("error")
			continue
		}
		maxi = temp
		break
	}

	values := []int{}
	weights := []int{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("error")
			continue
		}
		weight, err_1 := strconv.Atoi(fields[0])
		value, err_2 := strconv.Atoi(fields[1])
		if err_1 != nil || err_2 != nil {
			fmt.Println("error")
			continue
		}
		values = append(values, value)
		weights = append(weights, weight)
	}
	answ := rykzak(precision, values, weights, maxi)

	fmt.Println(answ.weight, answ.value)
	for _, idx := range answ.idexes {
		fmt.Println(idx + 1)
	}
}
