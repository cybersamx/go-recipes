package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type taskFunc func(timer time.Time)

func getPickFromInput() (int64, time.Duration) {
	var strategy int64
	var freq int64

	for strategy <= 0 || strategy > 4 {
		fmt.Println("1 - Using time.Sleep()\n2 - Using Ticker #1\n3 - Using Ticker #2\n4 - Using time.After")
		fmt.Print("Select a scheduler strategy: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal("error capturing strategy")
		}
		strategy, err = strconv.ParseInt(input, 10, 64)
		if err != nil {
			log.Fatal("error parsing strategy")
		}

		fmt.Print("Frequency (in seconds): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			log.Fatal("error capturing frequency")
		}
		freq, err = strconv.ParseInt(input, 10, 64)
		if err != nil {
			log.Fatal("error parsing frequency")
		}
	}

	return strategy - 1, time.Duration(freq) * time.Second
}

func theTask() taskFunc {
	return func(t time.Time) {
		// The time received can be used to perform a bunch a stuff including
		// setting a timestamp, compare against another time to check for
		// expiration, etc.

		// For this recipe, we'll just print the time.
		fmt.Println(t)
	}
}

func schedStrategy1(task taskFunc, dur time.Duration) {
	for {
		task(time.Now())
		time.Sleep(dur)
	}
}

func schedStrategy2(task taskFunc, ticker *time.Ticker) {
	for {
		select {
		case t := <-ticker.C:
			task(t)
		}
	}
}

func schedStrategy3(task taskFunc, ticker *time.Ticker) {
	for t := range ticker.C {
		task(t)
	}
}

func schedStrategy4(task taskFunc, dur time.Duration) {
	for {
		select {
		case t := <-time.After(dur):
			task(t)
		}
	}
}

func main() {
	// Capture system signals to exit gracefully.
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	var ticker *time.Ticker

	// Mux for calling the right scheduler.
	strategy, freq := getPickFromInput()
	switch strategy {
	case 0:
		fmt.Println("Selected time.Sleep()")
		go schedStrategy1(theTask(), freq)
	case 1:
		fmt.Println("Selected Ticker #1")
		ticker = time.NewTicker(freq)
		go schedStrategy2(theTask(), ticker)
	case 2:
		fmt.Println("Selected Ticker #2")
		ticker = time.NewTicker(freq)
		go schedStrategy3(theTask(), ticker)
	case 3:
		fmt.Println("Selected time.After")
		go schedStrategy4(theTask(), freq)
	default:
		fmt.Println("Unsupported schedule strategy selected")
	}

	select {
	case <-quitChan:
		log.Println("received signal to quit gracefully")
		if ticker != nil {
			// Stop the ticker to avoid resource leak. For a simple program, placing the
			// ticker at the end of the program makes sense. But for more complex timer
			// that may have multiple start and stop logic. This mean that you may end up
			// creating multiple Tickers. If so, you want to make sure you stop the Ticker
			// prior to creating a new Ticker.
			ticker.Stop()
		}
	}
}
