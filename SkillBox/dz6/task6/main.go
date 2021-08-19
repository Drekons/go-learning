package main

import "fmt"

func main() {
	fmt.Println("*** Движение лифта ***")

	floors := 24
	firstPassengerFloor := 4
	secondPassengerFloor := 7
	thirdPassengerFloor := 10
	inElevator := 0
	maxPassengers := 2
	elevatorFlor := 1

	for firstPassengerFloor != 0 || secondPassengerFloor != 0 || thirdPassengerFloor != 0 || elevatorFlor != 24 {
		if elevatorFlor == 1 {
			fmt.Println("Пассажиров вышло из лифта:", inElevator)
			inElevator = 0
			elevatorFlor = floors
			fmt.Println("Лифт отправился на", elevatorFlor, "этаж")
			continue
		}

		elevatorFlor--

		if inElevator == maxPassengers {
			continue
		}

		if elevatorFlor == firstPassengerFloor {
			firstPassengerFloor = 0
			inElevator++
			fmt.Println("Лифт отправился на", elevatorFlor, "этаж")
			fmt.Println("В лифт зашёл пассажир, в лифте", inElevator, "пассажиров")
			continue
		}

		if elevatorFlor == secondPassengerFloor {
			secondPassengerFloor = 0
			inElevator++
			fmt.Println("Лифт отправился на", elevatorFlor, "этаж")
			fmt.Println("В лифт зашёл пассажир, в лифте", inElevator, "пассажиров")
			continue
		}

		if elevatorFlor == thirdPassengerFloor {
			thirdPassengerFloor = 0
			inElevator++
			fmt.Println("Лифт отправился на", elevatorFlor, "этаж")
			fmt.Println("В лифт зашёл пассажир, в лифте", inElevator, "пассажиров")
			continue
		}
	}
}
