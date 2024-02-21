/// Channels
/// https://www.youtube.com/watch?v=8uiZC0l4Ajw
// Way of working with Go Routines, channels listen for updates and then can trigger options off them.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel: // If Channel is chickenChannel do this
		fmt.Printf("\nText Sent: Found deal on chicken at %v", website)
	case website := <-tofuChannel: // If Channel is tofuChannel do this
		fmt.Printf("\nEmail Sent: Found deal on tofu at %v", website)
	}
}

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel) // Check Chicken prices accross all websites
		go checkTofuPrices(websites[i], tofuChannel)       // Check Tofu prices accross all websites
	}
	sendMessage(chickenChannel, tofuChannel) // When the first channel to find a result completes trigger a message to be sent depending on the channel which completed.
}
