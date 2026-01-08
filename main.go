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
	numbers := []int{}

	/*operation := []string{}
	operation = append(operation, "AVG")
	operation = append(operation, "SUM")
	operation = append(operation, "MED")
	*/

	operation := map[string]func([]int) float64{
		"AVG": avg,
		"SUM": sum,
		"MED": med,
	}

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
		//if elementInArray(sourceOperation, operation) {
		//	isOperation = true
		//}

		handler, ok := operation[sourceOperation] // Проверяем, существует ли ключ
		if ok {
			handler(numbers) // Вызываем найденную функцию
			fmt.Println(sourceOperation, ":", handler(numbers))
		} else {
			fmt.Printf("Команда '%s' не найдена\n", sourceOperation)
		}
	}
}

/*func elementInArray(element string, elements []string) bool {
	//fmt.Println("elementinarray")
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return true
		}
	}
	return false
}
*/

func sum(numbers []int) float64 {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return float64(sum)
}

func avg(numbers []int) float64 {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
}

func med(numbers []int) float64 {
	sort.Ints(numbers)
	mNumber := len(numbers) / 2
	result := 0.0
	if len(numbers)%2 == 0 {
		result = (float64(numbers[mNumber-1]) + float64(numbers[mNumber])) / 2

	} else {

		result = float64(numbers[mNumber])
	}

	return result
}

/*func calculate(sourceOperation string, numbers []int) float64 {
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
			result = (float64(numbers[mNumber-1]) + float64(numbers[mNumber])) / 2

		} else {

			result = float64(numbers[mNumber])
		}
	}
	return result
}
*/
