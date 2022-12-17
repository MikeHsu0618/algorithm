package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 睡眠的理髮師問題
//1. 問題描述
//理髮店理有一位理髮師、一把理髮椅和n把供等候理髮的顧客坐的椅子，要求：
//
//如果沒有顧客，理髮師便在理髮椅上睡覺
//一個顧客到來時，它必須叫醒理髮師
//如果理髮師正在理髮時又有顧客來到，則如果有空椅子可坐，就坐下來等待，否則就離開
//2.問題分析
//理髮師和顧客是同步關係，理髮師等待顧客來，然後為顧客服務，顧客來了之後叫醒理髮師，執行上是有先後順序的，所以應該有兩個同步信號量，且散在兩個進程裡；另一方面，顧客對椅子的操作又是互斥的，屬於競爭關係，所以需要互斥信號量來保證椅子的數量準確。

// This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem
// which illustrates the complexities that arise when there are multiple operating system processes. Here, we have
// a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is
// open, and clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the
// waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to
// sleep until a new client arrives. So the rules are as follows:
//
//		- if there are no customers, the barber falls asleep in the chair
//		- a customer must wake the barber if he is asleep
//		- if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and
//		  sits in an empty chair if it's available
//		- when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers
//		  and falls asleep if there are none
// 		- shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is
//	      empty
//		- after the shop is closed and there are no clients left in the waiting area, the barber
//		  goes home
//
// The Sleeping Barber was originally proposed in 1965 by computer science pioneer Edsger Dijkstra.
//
// The point of this problem, and its solution, was to make it clear that in a lot of cases, the use of
// semaphores (mutexes) is not needed.

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our rand number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	fmt.Println("The Sleeping Barber Problem")
	fmt.Println("---------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarberDoneChan:  doneChan,
		ClientChan:      clientChan,
		Open:            true,
	}

	fmt.Println("The shop is open for today!")

	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Frank1")
	shop.addBarber("Frank2")
	shop.addBarber("Frank3")
	shop.addBarber("Frank4")
	shop.addBarber("Frank")
	shop.addBarber("Frank1")
	shop.addBarber("Frank2")
	shop.addBarber("Frank3")
	shop.addBarber("Frank4")
	shop.addBarber("Frank")
	shop.addBarber("Frank1")
	shop.addBarber("Frank2")
	shop.addBarber("Frank3")
	shop.addBarber("Frank4")
	shop.addBarber("Frank")
	shop.addBarber("Frank1")
	shop.addBarber("Frank2")
	shop.addBarber("Frank3")
	shop.addBarber("Frank4")
	// start and close the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1
	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliSeconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliSeconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
