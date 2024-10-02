package main

import (
	"fmt"
	"time"
)

func doWork(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Second)
		fmt.Print(".")
	}
	time.Sleep(time.Second)
}

func worker(done chan bool) {
	fmt.Print("working")
	doWork(3)
	fmt.Println(" done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

/*
package main

import (
	"fmt"
	"log"
	"time"
)

func countup(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println(from, ":", i)
	}
}

func main() {
	log.Println("Begin main")

	log.Println("Begin countup direct")
	countup("direct")

	log.Println("Begin countup goroutine")
	go countup("goroutine")

	log.Println("Begin anon goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	log.Println("Begin making channels")
	messages := make(chan string)
	inbox := make(chan string)

	log.Println("Begin sending to channels")
	go func() { messages <- "ping" }()
	go func() { inbox <- "pong" }()

	log.Println("Begin dumping channels to vars")
	msg := <-messages
	mail := <-inbox

	log.Println("Begin print vars")
	fmt.Println(msg)
	fmt.Println(mail)

	time.Sleep(time.Second * 3)

	log.Println("End main")
}
*/
