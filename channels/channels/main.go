package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//A channel is created to send messages through the routines
	c := make(chan string)

	//A Go routine is created for each item in the slice.
	for _, link := range links {
		go checkLink(link, c)
	}

	//Receiving channel data from the checkLink function
	//The way to expect the exact number of messages on the channel before
	//closing the go routine.
	//This is executed when we finish an http request.

	//removing the conditions from a for loop, it iterates eternally

	// for {
	// 	//The application keeps expecting messages from the channel
	// 	//A value coming from c, is a "blocking" line of code
	// 	go checkLink(<-c, c)
	// }
	//The control variable for the loop, does not handle an iteration, thus being infinite.

	//ALTERNATIVE: if you put this condition on a loop, this will make the same stuff as line 35
	for l := range c {
		//this is the function that makes you wait a little bit after requesting a new go routine.
		//BE CAREFUL WHEN DOING THAT BECAUSE THAT MIGHT CAUSE BACKED UP MESSAGES IN A CHANNEL
		//time.Sleep(5 * time.Second)
		//go checkLink(l, c)

		//THIS IS A LITERAL FUNCTION
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link

		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
