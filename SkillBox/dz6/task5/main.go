package main

import "fmt"

func main() {
	fmt.Println("*** Задача на постепенное наполнение корзин разной ёмкости (if и continue и break) ***")

	var basketOneVolume int
	var basketTwoVolume int
	var basketThreeVolume int
	var basketOne int
	var basketTwo int
	var basketThree int

	fmt.Print("Введите объём первой корзины: ")
	_, _ = fmt.Scan(&basketOneVolume)
	fmt.Print("Введите объём первой корзины: ")
	_, _ = fmt.Scan(&basketTwoVolume)
	fmt.Print("Введите объём первой корзины: ")
	_, _ = fmt.Scan(&basketThreeVolume)

	var currentBasket *int
	var currentBasketVolume *int
	for {
		for i := 0; i < 3; i++ {
			if i == 0 {
				currentBasket = &basketOne
				currentBasketVolume = &basketOneVolume
			}
			if i == 1 {
				currentBasket = &basketTwo
				currentBasketVolume = &basketTwoVolume
			}
			if i == 2 {
				currentBasket = &basketThree
				currentBasketVolume = &basketThreeVolume
			}

			if *currentBasket == *currentBasketVolume {
				continue
			}

			*currentBasket++
			fmt.Println("В корзине", i+1, "теперь", *currentBasket, "яблок")

			if *currentBasket == *currentBasketVolume {
				fmt.Println("--- Корзина", i+1, "заполнена ---")
			}
		}

		if basketOneVolume == basketOne && basketTwoVolume == basketTwo && basketThreeVolume == basketThree {
			break
		}
	}
}
