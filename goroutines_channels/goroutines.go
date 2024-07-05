package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4"}
var wg = sync.WaitGroup{}

// To Read And Write parallely we use mutex
var lock = sync.RWMutex{}

var results = []string{}

func main() {

	fmt.Println("Go Routines are used to run async operations it will use threads and run the operations in background")
	to := time.Now()
	//Normal Call
	tSync := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("Time Taken Total For Normal :: %v \n", time.Since(tSync))

	//If you dont use wait then it will just close the program before the functions are finished
	tAsyncWG := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCallWG(i)
	}
	wg.Wait()
	fmt.Println(results)
	fmt.Printf("Time Taken Total For Async With WG :: %v \n", time.Since(tAsyncWG))

	//Async Call
	tAsync := time.Now()
	for i := 0; i < len(dbData); i++ {
		go dbCall(i)

	}
	fmt.Printf("Time Taken Total For Async Without WG :: %v \n", time.Since(tAsync))
	fmt.Printf("Time Taken Total :: %v \n", time.Since(to))
	channels()
}

func dbCall(number int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The Result is :", dbData[number])
}
func dbCallWG(number int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The Result WG is :", dbData[number])
	save(dbData[number])
	log()
	wg.Done()
}

func save(result string) {
	lock.Lock()
	results = append(results, result)
	lock.Unlock()
}

func log() {
	lock.RLock()
	fmt.Printf("Log :: Results Till Now :: %v \n ", results)
	lock.RUnlock()
}

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 5

func channels() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var webstores = []string{"walmart", "costco", "flipkart"}
	for i := range webstores {
		go checkChickenPrice(webstores[i], chickenChannel)
		go checkTofuPrice(webstores[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrice(store string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			c <- store
			break
		}
	}
}

func checkTofuPrice(store string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			c <- store
			break
		}
	}
}

func sendMessage(chick, tofu chan string) {
	fmt.Printf("\nFound a deal on chicken -> %v \n", <-chick)
	fmt.Printf("\nFound a deal on tofu -> %v \n", <-tofu)
}
