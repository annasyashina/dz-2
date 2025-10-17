package main

import "fmt"

func main() {

	operation := []string{}
	operation = append(operation, "AVG")
	operation = append(operation, "SUM")
	operation = append(operation, "MED")

	var str string
	fmt.Println("Введите строку неограниченное число чисел через запятую")
	fmt.Scan(&str)

	isOperation := false
	var sourceOperation string

	for isOperation == false {
		fmt.Println("Введите операцию", operation, ":")
		fmt.Scan(&sourceOperation)
		if elementInArray(sourceOperation, operation) {
			isOperation = true
		}
	}

}

func elementInArray(element string, elements []string) bool {
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return true
		}
	}
	return false
}
