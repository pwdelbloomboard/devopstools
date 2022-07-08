# https://pkg.go.dev/os/exec#Cmd.Run

> Run starts the specified command and waits for it to complete.

> The returned error is nil if the command runs, has no problems copying stdin, stdout, and stderr, and exits with a zero exit status.

> If the command starts but does not complete successfully, the error is of type *ExitError. Other error types may be returned for other situations.

> If the calling goroutine has locked the operating system thread with runtime.LockOSThread and modified any inheritable OS-level thread state (for example, Linux or Plan 9 name spaces), the new process will inherit the caller's thread state.

# Approaches to Error Handling of os/exec#cmd.Run

## Capturing the Error

* This is essentially the, "normal" way to run an error:

```
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
```

We're essentially capturing the error output and printing out the err itself.

## Specifically Not Capturing the Error

* This is cited to be the, "wrong" way to handle the error:

```
	_ = cmd.Run()
	log.Printf("Command finished, no error given because we used _")
```

* Basically, omitting errors is generally never a good idea, because in the future, you (the person who coded said error ignore), or someone else will change the code and not update the subsequent error handling (or lack of error handling), and the code would fail without telling the user why this is not failing
* The very least, minimum effort is to at least log the errors, to create a panidc for errors that shouldn't ever happen.

* What happens if the above code is run and there is an error, e.g. if the cmd is a bad command?  Well if for example, we feed in a non-sensical command, such as, "sleepst" as shown, then we will get the same successful output as shown below in the go run terminal, even though there in fact has been an error.

```
func main() {
	cmd := exec.Command("sleepst", "1")
	log.Printf("Running command and waiting for it to finish...")
	_ = cmd.Run()
	log.Printf("Command finished, no error given because we used _")
}


2022/07/08 17:09:15 Running command and waiting for it to finish...
2022/07/08 17:09:15 Command finished, no error given because we used _

```
* So literally, the machine output is wrong.


## Creating a Panic In the Case of an Error That Should, "Never Happen"

```
	err := cmd.Run()
	if err != nil {
		panic("this is a panic that is never supposed to happen!")
	} else {
		log.Printf("Command finished with error: %v", err)
	}
```
In this instance, if we feed in a bad command, such as, "sleepst" then we at least get a proper output panic showing that there was an error:

```
2022/07/08 17:11:39 Running command and waiting for it to finish...
panic: this is a panic that is never supposed to happen!

goroutine 1 [running]:
main.main()
	/home/volumebindmount/osexec/osexecrun.go:13 +0xce
exit status 2
```

* Of course then we also get an exit status.

## Creating an Error Handling Function with errors.New

Much like we created [here](https://github.com/pwdelbloomboard/devopstools/blob/main/about-go/go-docker/volumebindmount/retryaftererror/retryaftererror.go), we can potentially create a new error which is allowed to pass through a certain number of iterations.

This was executed properly within ./osexecretry.go, however I replaced the function f with cmd.Run().

* Note that if we make err = cmd.Run() then the result of err gets returned by the return expression.
* Hwoever if we make err := cmd.Run() then the result of err becomes nil.




# Resources

* [How to ignore returned errors - with warnings not to do so](https://stackoverflow.com/questions/31027579/how-to-ignore-returned-error-in-go)
* [Is it safe to ignore errors in Golang?](https://stackoverflow.com/questions/62594017/error-handing-in-go-when-to-ignore-errors)