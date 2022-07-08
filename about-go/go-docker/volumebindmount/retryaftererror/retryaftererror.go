package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	var string1 string = "Fail"

	//string is never being changed to "pass" so we should have 50 retries.
	err := retry(5, 2*time.Second, func() error {
		if string1 == "Pass" {
			fmt.Println("Pass")
			return nil
		}
		return errors.New("fail")

	})
	if err != nil || string1 != "Pass" {
		log.Println(err)
		fmt.Println("I have logged the err above")
		return
	}

}

func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; i < attempts; i++ {
		fmt.Println("This is attempt number", i)
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		err = f()
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
