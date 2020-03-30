package main

// Concurancy is not parallelism

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://google.com",
		"http://golang.org",
		"http://flipkart.com",
	}

	c := make(chan string)

	for _, link := range links {
		// can run only on function
		go checkLink(link, c)
	}

	/* If we give with 4 print in main function it will wait for the message on the channel
	if we give one additional, it will hang
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println(<-c)
	*/
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up")
	c <- "Yes, its up"

}
