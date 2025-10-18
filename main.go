package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	operation := []string{}
	operation = append(operation, "AVG")
	operation = append(operation, "SUM")
	operation = append(operation, "MED")

	//var str string
	fmt.Println("Введите строку неограниченное число чисел через запятую")
	//fmt.Scan(&str)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	str := scanner.Text()
	fmt.Println("captured:", str)

	//fmt.Println(str)

	str = strings.ReplaceAll(str, " ", "")

	fmt.Println(str)
	parts := strings.Split(str, ",")
	fmt.Println(parts)
	numbers := []int{}
	for i := 0; i < len(parts); i++ {
		if value, err := strconv.Atoi(parts[i]); err == nil {
			//fmt.Printf("%q looks like a number.\n", parts[i])
			numbers = append(numbers, value)
		}
	}
	isOperation := false
	var sourceOperation string

	for isOperation == false {
		//fmt.Println(isOperation)
		fmt.Println("Введите операцию", operation, ":")
		fmt.Scan(&sourceOperation)
		//	fmt.Println(sourceOperation)
		if elementInArray(sourceOperation, operation) {
			isOperation = true
		}
	}

	fmt.Println(sourceOperation, ":", calculate(sourceOperation, numbers))
}

func elementInArray(element string, elements []string) bool {
	//fmt.Println("elementinarray")
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return true
		}
	}
	return false
}

func calculate(sourceOperation string, numbers []int) float64 {
	result := 0.0
	switch sourceOperation {
	case "AVG":
		var sum int
		for _, number := range numbers {
			sum += number
		}
		result = float64(sum) / float64(len(numbers))
	case "SUM":
		var sum int
		for _, number := range numbers {
			sum += number
		}
		result = float64(sum)
	case "MED":
		sort.Ints(numbers)
		mNumber := len(numbers) / 2

		if len(numbers)%2 == 0 {
			result = float64(numbers[mNumber])
		}

		result = (float64(numbers[mNumber-1]) + float64(numbers[mNumber])) / 2
	}
	return result
}
