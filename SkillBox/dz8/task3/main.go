package main

import "fmt"

func main() {
	var bills []int

	fmt.Println("Введите сколько заплатили покупатели")
	fmt.Println("Что бы закончить ввод введите 0")

	var customer int
customerInputLoop:
	for {
		customer++

		for {
			fmt.Print("Покупатель ", customer, " [5, 10, 20]: ")

			var sum int
			_, _ = fmt.Scan(&sum)

			if sum == 0 {
				break customerInputLoop
			}

			if sum != 5 && sum != 10 && sum != 20 {
				continue
			}

			bills = append(bills, sum)
			break
		}

	}

	fmt.Printf("Счета покупателей: %v\n", bills)
	fmt.Print("Можем ли мы дать сдачу каждому покупателю: ")
	if lemonadeChange(bills) {
		fmt.Print("Да")
	} else {
		fmt.Print("Нет")
	}
	fmt.Print("\n")
}

func lemonadeChange(bills []int) bool {
	var cash5 int
	var cash10 int
	var cash20 int
	for _, bill := range bills {
		switch bill {
		case 5:
			cash5++
		case 10:
			if cash5 == 0 {
				return false
			}
			cash5--
			cash10++
		case 20:
			if (cash10 == 0 && cash5 < 3) || (cash10 > 0 && cash5 == 0) {
				return false
			}

			if cash10 > 0 {
				cash10--
				cash5--
			} else {
				cash5 -= 3
			}

			cash20++
		}
	}

	return true
}
