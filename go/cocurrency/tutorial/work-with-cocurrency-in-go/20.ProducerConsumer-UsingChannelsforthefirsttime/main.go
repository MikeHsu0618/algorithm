package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	fmt.Println("The Pizzeria is open for business")
	fmt.Println("=================================")
	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)
	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				fmt.Println(i.message)
				fmt.Printf("Order #%d is out for delivery", i.pizzaNumber)
				fmt.Println("")
			} else {
				fmt.Println(i.message)
				fmt.Printf("Order #%d is really mad", i.pizzaNumber)
				fmt.Println("")
			}
		} else {
			fmt.Println("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				fmt.Println("*** Error closing channel ***", err)
			}
		}
	}
	// print out the ending message
	fmt.Println("----------------------------")
	fmt.Println("Done for the day")
	fmt.Printf("We make %d pizzas, but failed to make %d, with %d attempts in total.", pizzaMade, pizzaFailed, total)

	switch {
	case pizzaFailed > 9:
		fmt.Println("Is was an awful day... ")
	case pizzaFailed >= 6:
		fmt.Println("It was not a very good day")
	case pizzaFailed >= 4:
		fmt.Println("It was an okay day")
	}
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Receivec order #%d", pizzaNumber)
		fmt.Println()
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}
		total++

		fmt.Printf("Making Pizza #%d. it will take %d seconds... \n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd < 3 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d ***", pizzaNumber)
		} else if rnd < 5 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d ***", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza Order is ready #%d", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0
	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// try to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
		// decision

	}
}
