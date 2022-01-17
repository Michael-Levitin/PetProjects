package main

import (
	"fmt"
)

type consumables struct {
	water, milk, beans, cups, money, whiskey int
}

var (
	espresso = consumables{
		water: 250,
		beans: 16,
		money: 4,
	}
	latte = consumables{
		water: 350,
		milk:  75,
		beans: 20,
		money: 7,
	}
	cappuccino = consumables{
		water: 200,
		milk:  100,
		beans: 12,
		money: 6,
	}
	irish = consumables{
		water:   200,
		milk:    100,
		beans:   12,
		money:   8,
		whiskey: 40,
	}
)

func main() {
	in := consumables{
		water: 400,
		milk:  540,
		beans: 120,
		cups:  9,
		money: 550,
	}
	in.selectOption()
}

func (c *consumables) selectOption() {
	var choice string
	for choice != "exit" {
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		fmt.Scan(&choice)
		switch choice {
		case "buy":
			c.buyCoffee()
		case "fill":
			c.fillConsumables2()
		case "take":
			c.takeMoney()
		case "remaining":
			c.printConsumables()
		case "exit":
			return
		default:
			fmt.Println("I don't understand this option")
		}
	}
}

func (c *consumables) buyCoffee() {
	var choice string
	fmt.Printf("\nWhat do you want to buy? 1 - espresso = %d$, 2 - latte = %d$, "+
		"3 - cappuccino = %d$, 4 - irish coffee = %d$, back - to main menu:",
		espresso.money, latte.money, cappuccino.money, irish.money)
	fmt.Scan(&choice)
	switch choice {
	case "1":
		err := c.checkConsumables(espresso)
		if err == nil {
			c.prepareCoffee(espresso)
		} else {
			fmt.Println(err)
			return
		}
	case "2":
		err := c.checkConsumables(latte)
		if err == nil {
			c.prepareCoffee(latte)
		} else {
			fmt.Println(err)
			return
		}
	case "3":
		err := c.checkConsumables(cappuccino)
		if err == nil {
			c.prepareCoffee(cappuccino)
		} else {
			fmt.Println(err)
			return
		}
	case "4":
		err := c.checkConsumables(irish)
		if err == nil {
			c.prepareCoffee(irish)
		} else {
			fmt.Println(err)
			return
		}
	case "back":
		return
	default:
		fmt.Println("I don't understand this option, returning to main menu")
		return
	}
	fmt.Println("I have enough resources, making you a coffee!")
}

func (c *consumables) fillConsumables() {
	var temp int
	fmt.Println("\nWrite how many ml of water you want to add:")
	fmt.Scan(&temp)
	c.water += temp
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&temp)
	c.milk += temp
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&temp)
	c.beans += temp
	fmt.Println("Write how many disposable coffee cups you want to add:")
	fmt.Scan(&temp)
	c.cups += temp
}

//arabika (100%) arabika/robusta (50/50), 10
func (c *consumables) fillConsumables2() {
	var water, milk, beans, cups, whiskey int
	for {
		fmt.Println("\nWrite how much consumables you want to add: \nwater(ml), milk(ml), coffee beans(gr), cups, whiskey(ml)")
		_, err := fmt.Scan(&water, &milk, &beans, &cups, &whiskey)
		if err == nil {
			break
		} else {
			fmt.Println(err)
		}
	}

	c.water += water
	c.milk += milk
	c.beans += beans
	c.cups += cups
	c.whiskey += whiskey
}

func (c *consumables) prepareCoffee(p consumables) {
	c.water -= p.water
	c.milk -= p.milk
	c.beans -= p.beans
	c.whiskey -= p.whiskey
	c.cups--
	c.money += p.money
}

func (c *consumables) checkConsumables(t consumables) error {
	switch {
	case t.cups > c.cups:
		return fmt.Errorf("Sorry, not enough cups")
	case t.beans > c.beans:
		return fmt.Errorf("Sorry, not enough coffee beans")
	case t.milk > c.milk:
		return fmt.Errorf("Sorry, not enough milk")
	case t.water > c.water:
		return fmt.Errorf("Sorry, not enough water")
	case t.whiskey > c.whiskey:
		return fmt.Errorf("Sorry, not enough liquor")
	}
	return nil
}

func (c *consumables) printConsumables() {
	fmt.Println("\nThe coffee machine has:")
	fmt.Println(c.water, "ml of water")
	fmt.Println(c.milk, "ml of milk")
	fmt.Println(c.beans, "gr of coffee beans")
	fmt.Println(c.cups, "of disposable cups")
	fmt.Println(c.money, "of money")
	fmt.Println(c.whiskey, "ml of liquor")
}

func (c *consumables) takeMoney() {
	fmt.Println("\nI gave you", c.money)
	c.money = 0
}

func (c *consumables) oldTasks() {
	fmt.Println("Write how many cups of coffee you will need:")
	var cupsCanMake, cupsWanted, n int
	fmt.Scan(&n)
	fmt.Printf("For %d cups of coffee you will need:", n)
	fmt.Println(n*200, "ml of water")
	fmt.Println(n*50, "ml of milk")
	fmt.Println(n*15, "g of coffee beans")

	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&c.water)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&c.milk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&c.beans)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cupsWanted)

	cupsCanMake = MinIntSlice([]int{c.water / 200, c.milk / 50, c.beans / 15})
	if cupsWanted == cupsCanMake {
		fmt.Println("Yes, I can make that amount of coffee")
	} else if cupsWanted > cupsCanMake {
		fmt.Printf("No, I can make only %d cups of coffee", cupsCanMake)
	} else {
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)", cupsCanMake-cupsWanted)
	}
}

func MinIntSlice(v []int) (m int) {
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return
}