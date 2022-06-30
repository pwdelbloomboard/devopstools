# Go Os Documentation Summary

https://pkg.go.dev/os

* Unix-like design of Operating System Functionality.
* Errors may include additional information, for example, if a call that takes a filename fails, the error will include the failing file name.

### Functionality Overview

* os.Open - open files.
* os.Whatever

* Chmod
* Chtimes
* CreateTemp
* CreateTemp (Suffix)
* Expand
* ExpandEnv
* FileMode
* Getenv
* LookupEnv
* Mkdir
* MkdirAll
* MkdirTemp
* MkdirTemp (Suffix)
* OpenFile
* OpenFile (Append)
* ReadDir
* ReadFile
* Unsetenv
* WriteFile

...and more

### Individual Functions

#### Pipe

os.Pipe()

* Pipe returns a connected pair of Files; reads from r return bytes written to w. It returns the files and an error, if any.


#### Stdout

os.Stdout()

* 

#### Close

os.Close()

* Close closes the File, rendering it unusable for I/O. On files that support SetDeadline, any pending I/O operations will be canceled and return immediately with an ErrClosed error. Close will return an error if it has already been called.

# Go ioutil Summary

https://pkg.go.dev/io/ioutil

### Functionality Overview

Package ioutil implements some I/O utility functions.

ReadAll
ReadDir
ReadFile
TempDir
TempDir (Suffix)
TempFile
TempFile (Suffix)
WriteFile


### Individual Functions

#### ioutil.ReadAll

ReadAll reads from r until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.

# Go Io Documentation Summary

https://pkg.go.dev/io

### Functionality Overview

> Package io provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package os, into shared public interfaces that abstract the functionality, plus some other related primitives.

> Because these interfaces and primitives wrap lower-level operations with various implementations, unless otherwise informed clients should not assume they are safe for parallel execution.

* Copy
* CopyBuffer
* CopyN
* LimitReader
* MultiReader
* MultiWriter
* Pipe
* ReadAll
* ReadAtLeast
* ReadFull
* SectionReader
* SectionReader.Read
* SectionReader.ReadAt
* SectionReader.Seek
* SectionReader.Size
* TeeReader
* WriteString

### Individual Functions

#### io.Copy

> func Copy(dst Writer, src Reader)

> Copy copies from src to dst until either EOF is reached on src or an error occurs. It returns the number of bytes copied and the first error encountered while copying, if any.

> A successful Copy returns err == nil, not err == EOF. Because Copy is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.

> If src implements the WriterTo interface, the copy is implemented by calling src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom interface, the copy is implemented by calling dst.ReadFrom(src).

#### io.EOF

> EOF is the error returned by Read when no more input is available. (Read must return EOF itself, not an error wrapping EOF, because callers will test for EOF using ==.) Functions should return EOF only to signal a graceful end of input. If the EOF occurs unexpectedly in a structured data stream, the appropriate error is either ErrUnexpectedEOF or some other error giving more detail.

#### io.MultiWriter

> func MultiWriter(writers ...Writer) Writer

> MultiWriter creates a writer that duplicates its writes to all the provided writers, similar to the Unix tee(1) command.

> Each write is written to each listed writer, one at a time. If a listed writer returns an error, that overall write operation stops and returns the error; it does not continue down the list.

Note: a Writer is a type.

> Writer is the interface that wraps the basic Write method.

> Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early. Write must return a non-nil error if it returns n < len(p). Write must not modify the slice data, even temporarily.

> Implementations must not retain p.

#### io.PipeWriter

> A PipeWriter is the write half of a pipe.

# Go Bytes Documentation Summary

https://pkg.go.dev/bytes

> The bytes package implements functions for the manipulation of byte slices. It is similar to the strings package.

### Individual Functions

#### bytes.Buffer

A Buffer is a variable-sized buffer of bytes with Read and Write methods. The zero value for Buffer is an empty buffer ready to use.

Recall that:

> A byte in Go is an unsigned 8-bit integer. It has type uint8. A byte has a limit of 0 – 255 in numerical range. It can represent an ASCII character.

For example:

```
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
}
```