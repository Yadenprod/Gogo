package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TechnicalPartArab(num1 string, operator string, num2 string) int {
	var numArab1, numArab2, result int
	arabicNum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num1Int, _ := strconv.Atoi(num1)
	num2Int, _ := strconv.Atoi(num2)
	for i := range arabicNum {
		if arabicNum[i] == num1Int {
			numArab1 += num1Int
		}
	}
	for i := range arabicNum {
		if arabicNum[i] == num1Int {
			numArab2 += num2Int
		}
	}

	if numArab1 > 0 && numArab2 > 0 {
		switch operator {
		case "+":
			result = numArab1 + numArab2
		case "-":
			result = numArab1 - numArab2
		case "*":
			result = numArab1 * numArab2
		case "/":
			result = numArab1 / numArab2
		}
	}
	return result
}
func TechnicalPartRoman(num1 string, operator string, num2 string) string {
	var a, b int
	var result int
	romanConv := []string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
		"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	for i := range romanConv {
		if num1 == romanConv[i] {
			a = i
		}
		if num2 == romanConv[i] {
			b = i
		}
	}
	if a <= 0 || b <= 0 || a > 10 || b > 10 {
		err := errors.New("Косячнул ты))) Введи число от '1' до '10' включительно:)")
		fmt.Println("", err)
		return " Исправляйся)"
	}
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	if result < 1 {
		fmt.Println("Ответ не может быть меньше единицы")
		return "Error"
	}
	RoomeResult := romanConv[result]
	return RoomeResult
}
func main() {
	var primer string
	gogo := true
	primer = Scan1()
	if strings.Count(primer, " ") != 2 {
		fmt.Println("Введите числа по типу x + y, не более 1 операции с пробелами между знаками и оператором")
		os.Exit(1)
	}
	priimer := strings.Split(primer, " ")
	num1Int, _ := strconv.Atoi(priimer[0])
	num2Int, _ := strconv.Atoi(priimer[2])
	if num1Int == 0 && num2Int == 0 {
		fmt.Println(TechnicalPartRoman(priimer[0], priimer[1], priimer[2]))
		gogo = false
	}
	if num1Int == 0 && num2Int > 0 || num1Int > 0 && num2Int == 0 {
		fmt.Println("Вводи числа либо на римских, либо на арабских. с 1 до 10 включительно")
		os.Exit(2)
	}
	if gogo == true {
		if num1Int <= 0 || num1Int > 10 || num2Int <= 0 || num2Int > 10 {
			fmt.Println("Вводи числа либо на римских, либо на арабских. с 1 до 10 включительно")
			os.Exit(3)
		}
		fmt.Println(TechnicalPartArab(priimer[0], priimer[1], priimer[2]))
	}
}
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
