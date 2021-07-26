// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	var sites = []string{
// 		"https://www.google.com/",
// 		"https://www.facebook.com/",
// 		"https://golang.org/",
// 		"https://vnexpress.net/",
// 	}

// 	c := make(chan string)

// 	for _, site := range sites {
// 		go CheckSite(c, site)
// 	}

// 	for {
// 		l := <-c
// 		go func() {
// 			time.Sleep(5 * time.Second)
// 			CheckSite(c, l)
// 		}()
// 	}
// }

// func CheckSite(c chan string, url string) {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		c <- "error" + url + " " + err.Error()
// 		return
// 	}

// 	c <- url
// 	fmt.Println(url + " is " + res.Status)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(done chan bool) {
// 	fmt.Println("working...")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("done")

// 	done <- true
// }

// func main() {

// 	done := make(chan bool, 1)
// 	go worker(done)

// 	<-done

// 	fmt.Println("end")
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func ping(pings chan<- string, msg string) {
// 	time.Sleep(5 * time.Second)
// 	pings <- msg
// }

// func pong(pings <-chan string, pongs chan<- string) {
// 	msg := <-pings
// 	pongs <- msg
// }

// func main() {
// 	pings := make(chan string, 1)
// 	pongs := make(chan string, 1)
// 	go ping(pings, "passed message")
// 	go pong(pings, pongs)
// 	fmt.Println(<-pongs)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		c1 <- "one"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		c2 <- "two"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("received", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("received", msg2)
// 		default:
// 			fmt.Println("no message received")
// 		}
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 10; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
