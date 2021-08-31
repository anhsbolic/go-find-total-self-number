package main

import (
	"fmt"
	"strconv"
)

func main() {
	total := 0

	i := 1
	for i <= 5000 {
		stringValue := strconv.Itoa(i)

		if !isNumberHaveGenerator(stringValue) {
			total = total + i
			msg := fmt.Sprintf(`%s self number`, stringValue)
			fmt.Println(msg)
		}

		i++
	}

	msgTotal := fmt.Sprintf(`Total of self numbers from 1 to 5000 is %d`, total)
	fmt.Println(msgTotal)
}

func isNumberHaveGenerator(stringValue string) bool {
	value, _ := strconv.Atoi(stringValue)
	isHaveGenerator := false

	i := 1
	for i <= value {
		stringI := strconv.Itoa(i)
		if isTheGenerator(stringI, value) {
			isHaveGenerator = true
			break
		}

		i++
	}

	return isHaveGenerator
}

func isTheGenerator(value string, dn int) bool {
	var numbers []string

	i := 0
	for i < len(value) {
		numbers = append(numbers, value[i:i+1])
		i++
	}

	v, _ := strconv.Atoi(value)
	for _, number := range numbers {
		num, _ := strconv.Atoi(number)
		v = v + num
	}

	if v == dn {
		return true
	}

	return false
}
