package main

import (
	"fmt"
)

type GameState struct {
	Location  string
	Inventory []string
}

func showOptions(state GameState) {
	switch state.Location {
	case "печера":
		fmt.Println("Ви можете:")
		fmt.Println("1. Ввійти до печери")
		fmt.Println("2. Іти до лісу")
	case "ліс":
		fmt.Println("Ви можете:")
		fmt.Println("1. Оглянути мертву тварину")
		fmt.Println("2. Іти далі до табору")
		fmt.Println("3. Повернутися до печери")
	case "табір":
		fmt.Println("Ви можете:")
		fmt.Println("1. Відпочити в наметі")
		fmt.Println("2. Іти далі")
		fmt.Println("3. Повернутися до лісу")
	default:
		fmt.Println("Невідома локація")
	}
}

func handleAction(state *GameState, action int) {
	switch state.Location {
	case "печера":
		if action == 1 {
			fmt.Println("\nУ печері темно, і ви нічого не бачите.")
		} else if action == 2 {
			state.Location = "ліс"
			fmt.Println("\nВи йдете до лісу.")
		}
	case "ліс":
		if action == 1 {
			fmt.Println("\nВи бачите мертву тварину, але вирішуєте нічого не робити.")
		} else if action == 2 {
			state.Location = "табір"
			fmt.Println("\nВи йдете далі і приходите до табору.")
		} else if action == 3 {
			state.Location = "печера"
			fmt.Println("\nВи повертаєтесь до печери.")
		}
	case "табір":
		if action == 1 {
			if !contains(state.Inventory, "сейф") {
				fmt.Println("\nВи відпочиваєте в наметі і знаходите сейф.")
				state.Inventory = append(state.Inventory, "сейф")
			} else {
				fmt.Println("\nВи відпочиваєте в наметі, але більше нічого не знаходите.")
			}
		} else if action == 2 {
			fmt.Println("\nВи йдете далі, але відчуваєте втому.")
			state.Location = "ліс"
		} else if action == 3 {
			state.Location = "ліс"
			fmt.Println("\nВи повертаєтесь до лісу.")
		}
	default:
		fmt.Println("Невідома дія")
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	state := GameState{
		Location:  "печера",
		Inventory: []string{"сірники", "ліхтарик", "ніж"},
	}

	fmt.Println("\nВи прокинулись біля входу в печеру. Ви пам'ятаєте лише своє ім'я - Стівен.")
	for {
		fmt.Printf("\nВи знаходитесь в: %s\n", state.Location)
		fmt.Printf("Ваші речі: %v\n", state.Inventory)
		showOptions(state)
		fmt.Print("Введіть номер вашої дії: ")

		var action int
		_, err := fmt.Scan(&action)
		if err != nil {
			fmt.Println("Будь ласка, введіть номер дії.")
			continue
		}

		handleAction(&state, action)

		if state.Location == "печера" && action == 1 {
			fmt.Println("\nА все могло бути зовсім інакше.")
			break
		}
	}
}
