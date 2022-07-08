package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {

	// simply declaring a sting to demark a failure error.
	var string1 string = "Fail"

	//string is never being changed to "pass" so we should have specified number of retries.
	// call the retry function which takes in retry(attempts, sleepduration, function() errorstruct)
	// the error struct, error {} contains what is shown below, a including a new type of error with the message, "fail"
	err := retry(5, 2*time.Second, func() error {
		// string1 never equals Pass, this if statement is never satisfied
		if string1 == "Pass" {
			fmt.Println("Pass")
			return nil
		}

		// set the return string
		var returnString string = "fail - walka walka!"

		// return a New type of error with message, "fail"
		return errors.New(returnString)

	})
	// end of retry function

	// if err is not nil, or string1 is not Pass (always true)
	// happens at the end
	if err != nil || string1 != "Pass" {
		// log the err and result
		log.Println(err)
		fmt.Println("I have logged the err above")
		return
	}

}

// retry function
func retry(attempts int, sleep time.Duration, f func() error) (err error) {

	// for i from 0 to less than the number of attempts
	for i := 0; i < attempts; i++ {
		// report the attempt number
		fmt.Println("This is attempt number", i)
		// if greater than the start position
		if i > 0 {
			// logging the retry attempt
			log.Println("retrying after error:", err)
			// sleep for a specified duration times 2
			time.Sleep(sleep)
			sleep *= 2
		}
		// error is defined as the input function
		err = f()
		// if the error is nil, return nill
		if err == nil {
			// return nil
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
