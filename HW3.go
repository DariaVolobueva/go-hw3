package main

import (
	"fmt"
)

const (
	LocationCave  = "печера"
	LocationForest = "ліс"
	LocationCamp  = "табір"
)

type GameState struct {
	Location  string
	Inventory []string
}

func showOptions(state GameState) {
	switch state.Location {
	case LocationCave:
		fmt.Println("Ви можете:")
		fmt.Println("1. Ввійти до печери")
		fmt.Println("2. Іти до лісу")
	case LocationForest:
		fmt.Println("Ви можете:")
		fmt.Println("1. Оглянути мертву тварину")
		fmt.Println("2. Іти далі до табору")
		fmt.Println("3. Повернутися до печери")
	case LocationCamp:
		fmt.Println("Ви можете:")
		fmt.Println("1. Відпочити в наметі")
		fmt.Println("2. Іти далі")
		fmt.Println("3. Повернутися до лісу")
	default:
		fmt.Println("Невідома локація")
	}
}

func (state *GameState) handleAction(action int) {
	switch state.Location {
	case LocationCave:
		if action == 1 {
			fmt.Println("\nУ печері темно, і ви нічого не бачите.")
		} else if action == 2 {
			state.Location = LocationForest
			fmt.Println("\nВи йдете до лісу.")
		}
	case LocationForest:
		if action == 1 {
			fmt.Println("\nВи бачите мертву тварину, але вирішуєте нічого не робити.")
		} else if action == 2 {
			state.Location = LocationCamp
			fmt.Println("\nВи йдете далі і приходите до табору.")
		} else if action == 3 {
			state.Location = LocationCave
			fmt.Println("\nВи повертаєтесь до печери.")
		}
	case LocationCamp:
		if action == 1 {
			if !contains(state.Inventory, "сейф") {
				fmt.Println("\nВи відпочиваєте в наметі і знаходите сейф.")
				state.Inventory = append(state.Inventory, "сейф")
			} else {
				fmt.Println("\nВи відпочиваєте в наметі, але більше нічого не знаходите.")
			}
		} else if action == 2 {
			fmt.Println("\nВи йдете далі, але відчуваєте втому.")
			state.Location = LocationForest
		} else if action == 3 {
			state.Location = LocationForest
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
		Location:  LocationCave,
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

		state.handleAction(action)

		if state.Location == LocationCave  && action == 1 {
			fmt.Println("\nА все могло бути зовсім інакше.")
			break
		}
	}
}
