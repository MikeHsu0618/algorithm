package main

import (
	"fmt"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		fmt.Printf("%s goes to the waiting  room to check for clients", barber)
		fmt.Println()
		for {
			// if no clients, the barber goes to sleep
			if len(shop.ClientChan) == 0 {
				fmt.Printf("There is nothing to do, so %s takes a nap", barber)
				fmt.Println()
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientChan
			if !shopOpen {
				// shop is closed, so send the barber home and close the goroutine
				shop.sendBarberHome(barber)
				return
			}
			if isSleeping {
				fmt.Printf("%s wakes %s up.", client, barber)
				fmt.Println()
				isSleeping = false
			}
			// cut hair
			shop.curHair(barber, client)
		}

	}()
}

func (shop *BarberShop) curHair(barber, client string) {
	fmt.Printf("%s is cutting %s's hair", barber, client)
	fmt.Println()
	time.Sleep(shop.HairCutDuration)
	fmt.Printf("%s is finished cutting %s's hair", barber, client)
	fmt.Println()

}

func (shop *BarberShop) sendBarberHome(barber string) {
	fmt.Printf("%s is goinh home. \n", barber)
	shop.BarberDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	fmt.Println("Closing shop for the day")
	close(shop.ClientChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarberDoneChan
	}
	close(shop.BarberDoneChan)

	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("The barbershop is now closed for the day, and every one has gone home")
}

func (shop *BarberShop) addClient(client string) {
	fmt.Printf("**** client %s arrives ***", client)
	fmt.Println()
	if !shop.Open {
		fmt.Printf("The shop is already closed, so %s leaves.", client)
		return
	}

	select {
	case shop.ClientChan <- client:
		fmt.Printf("%s takes a sear in the waiting room", client)
		fmt.Println()
	default:
		fmt.Printf("The waiting room is full, so %s leaves.", client)
		fmt.Println()
	}
}
