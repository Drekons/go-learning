package main

import (
	"fmt"
)

func main() {
	fmt.Println("*** Игра “Угадывание числа” ***")

	var number int
	var guess int = 5
	var yes string = "да"
	var less string = "меньше"
	var more string = "больше"
	var answer string
	var shift int = guess / 2

	fmt.Print("Введите число от 1 до 10: ")
	fmt.Scan(&number)

	fmt.Print("Ваше число ", guess, "? [", yes, ", ", more, ", ", less, "]:")
	fmt.Scan(&answer)

	if answer != yes {
		if answer == more {
			guess += shift
		} else {
			guess -= shift
		}

		fmt.Print("Ваше число ", guess, "? [", yes, ", ", more, ", ", less, "]:")
		fmt.Scan(&answer)

		if answer != yes {
			if guess < 5 {
				shift = guess / 2
			} else {
				shift = (guess - 5) / 2
			}
			if answer == more {
				guess += shift
			} else {
				guess -= shift
			}

			fmt.Print("Ваше число ", guess, "? [", yes, ", ", more, ", ", less, "]:")
			fmt.Scan(&answer)
		}

		if answer != yes {
			if guess < 5 {
				shift = guess / 2
			} else {
				shift = (guess - 5) / 2
			}
			if answer == more {
				guess += shift
			} else {
				guess -= shift
			}

			fmt.Print("Ваше число ", guess, "? [", yes, ", нет]:")
			fmt.Scan(&answer)
		}
	}

	if answer == yes {
		fmt.Print("Ваше число ", guess, "! Программа победила!")
	} else {
		fmt.Print("Программа проиграла :(")
	}

}
