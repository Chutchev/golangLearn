package main

import "fmt"

func main() {
	fmt.Println("Введите месяц")
	var season string
	_, _ = fmt.Scan(&season)

	switch season {
	case "декабрь", "январь", "февраль":
		fmt.Println("Время года - Зима.")
	case "март", "апрель", "май":
		fmt.Println("Время года - Весна.")
	case "июнь", "июль", "август":
		fmt.Println("Время года - Лето.")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Время года - Осень.")
	}
}
