# Important Note from Tim

## TODO
- Learn about http.Handler, http.ResponseWriter, *http.Request, http.HandleFunc

## Pointer

- `*int`: type pointer to an int
-  `*n`: defererence `address n` to get a value
- `&a`: address of `value a`

## Reference type
- A reference type is pointer

- `Reference Type` includes: interface, slices, channels, functions, pointers, maps
- Remember the phase **a`I` `S`ẽ `C`ó `F`ần `P`hở `M`ập**
- `Reference type` has pointer attached to it, hence the term `reference type`

## Method Sets
- The idea of the method set is integral to how interfaces are implemented and used in Go.
    - The method set of a type T consists of all methods with receiver type T.
        - These methods can be called using variables of type T.
    - The method set of a type *T consists of all methods with receiver *T or T
        - These methods can be called using variables of type *T.
        - it can call methods of the corresponding non-pointer type as w

- Great example (**be patient** when reading it as it's **important** and **rewarding**)

```
type T struct {
}
func (t T) M1() {
// ...
}
func (t *T) M2() {
// ...
}
type X interface {
M1()
M2()
}
func main() {
var n X
Todd McLeod - Learn To Code Go - Page 121
t := T{}
n = t // This will be a compile error, because T does not implement M2()
tPtr := &T{}
n = tPtr // This is fine, *T implements M1() and M2()
}
```

- In this example, `T` implements `M1()` and `*T` implements both `M1()` and `M2()`. Therefore,
`*T` satisfies interface `X`, but `T` does not. So we can assign `&T{}` to `n`, but not `T{}`.

## READING DOCUMENTATIONS
### PACKAGE PAGES
- **Important**: You can either `initialize a type` or use a `function to produce a type` as describe below
  - This initializes a new bytes.Buffer
    ```
    var c bytes.Buffer
    ```
  - This produces a new bytes.Buffer using a function: d := bytes.NewBufferString(`my string`)
    ```
    d := bytes.NewBufferString(`my string`)
    ```
  
- When reading package's index page:
  - Under list of type, such as `type Decoder`: https://pkg.go.dev/encoding/json#pkg-index 
  - The first function usually guide you how you can get a function to produce such type (or most often `pointer of its type`), e.g `*Decoder`
  - `*T` can be used as `T`
    - `func NewDecoder(r io.Reader) *Decoder`
  - The rest of the functions in the same type are methods implemented by that type
    - type `*Decoder` have medthods `Buffered()`, `Decode(v interface{})`, `More()`, `Token()`, `UseNumber()`

  

## MAP VS JSON OBJECT
### MAP
- Map is with format: `map[T1]T2`
  - Example: 
    ```
    am := map[string]int{
            "Todd":   42,
            "Henry":  16,
            "Padget": 14,
        }
    ```
  - Adding members into map
    `am["Lucas"] = 28`

### JSON OBJECT
- Data structure encoded in JSON format with byte of slice such as raw string literal, pointer.
- To encode a Go map type it must be of the form map[string]T (where T is any Go type supported by the json package).
  - Example 1: raw string is converted into byte slice
    ```
    var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
    ```
  - Example 2: a VALUE of a struct is converted into byte slice
    ```
    type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)  // b is byte slice
    os.Stdout.Write(b)
    // Output: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
    ```

## ENCODE & DECODE WITH WRITER INTERFACE
### TECHNIQUE TO CONVERT ANY GOLANG DATA TO JSON OBJECT
- Great example: https://medium.com/what-i-talk-about-when-i-talk-about-technology/go-code-snippet-json-encoder-and-json-decoder-818f81864614
- Example:
  ```
  package main

  import (
    "bytes"
    "encoding/json"
    "fmt"
    "os"
  )

  type user struct {
    First   string
    Last    string
    Age     int
    Sayings []string
  }

  func main() {
    u1 := user{
      First: "James",
      Last:  "Bond",
      Age:   32,
      Sayings: []string{
        "Shaken, not stirred",
        "Youth is no guarantee of innovation",
        "In his majesty's royal service",
      },
    }

    u2 := user{
      First: "Miss",
      Last:  "Moneypenny",
      Age:   27,
      Sayings: []string{
        "James, it is soo good to see you",
        "Would you like me to take care of that for you, James?",
        "I would really prefer to be a secret agent myself.",
      },
    }

    users := []user{u1, u2}

    fmt.Println(users)

    // your code goes here
    // User bytes.Buffer to as space to write out the json string
    b := new(bytes.Buffer)
    err := json.NewEncoder(b).Encode(&users)
    if err != nil {
      fmt.Println("We did something wrong and here is the error: ", err)
    }
    fmt.Println("\n")
    fmt.Printf("Here is byte buffer in json with type\n: %+v \t %T", b.String(), b)
  }
  ```
- Summary:
  - `json.NewEncoder(new(bytes.Buffer))`: produces `*Encoder` (pointer to an Encoder)
  -  `json.NewEncoder(new(bytes.Buffer)).Encode(users)` encode `users` data into `new(bytes.Buffer)`
- Thought process:
  - Need to **encode** the data `[]user` (or `users`, `&users`), you can use function `func (enc *Encoder) Encode(v any) error` from `json` package where `v any` can be any Golang data including `[]user` (or `users`, `&users`)
  - Need to genrate the pointer to an Encoder (`*Encoder`) via function `func NewEncoder(w io.Writer) *Encoder`
  - Need `io.Writer` interface to as place holder or as receiver in the function `func NewEncoder(w io.Writer) *Encoder`
  - Any types that has method `Write()` is a type of the interface `io.Writer`, including `*File`, `bytes.Buffer`, `os.Stdout` as `os.Stdout` is a `File.NewFile` function producing pointer to a file (`*File`)


## GOBUILD, GOINSTALL, GOOS
```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOARCH, runtime.GOOS)
}
```
- `Go build` command
  - cd <folder>
  - go build
  - You'll see a package with the same name as <folder>

- `Go install` command
  - cd <folder>
  - go install
  - the package is installed in the bin folder ~/go/bin/. Need to configure $PATH to call reference ~/go/bin

## Manage, upgrade dependencies
- Reference: https://go.dev/doc/modules/managing-dependencies

## CHANNELS
- A Goroutine can send a value to the channel and another Goroutine can pick up. This has to be done by 2 separate Goroutines.
  - **Channel Block**: We need to `block` the MAIN GoRoutine until the value is completely sent to the channel by the other Goroutine AND is successfully received by the Main Goroutine
- A value can be sent to a `BUFFERED` channel by main goroutine and the main goroutine can also receive from the channel later

### Channels - Directional Channels
- Birectional Channel:  `c := make(chan int)`
- Directional Channel, read from left to right:
  - `cr := make(<-chan int)` // receive
  - `cs := make(chan<- int)` // send  - On this channel, we're sending an int

### Channels - Range over a receive channel
- `Range` clause: We can range over the receive channel to pull values off from the channel `UNTIL THE CHANNEL IS CLOSED`
  - **Important**: That means in the `send` channel, we must have a close statement, e.g. `close(c)`. This means this also block in the MAIN Goroutine until the receive operation completes!
### Channels - Select Statement
- `Select` statement: We can use `select` statement in tandem with infinite loop to pull values from multiple channel at the same time, whichever value is ready first will be pulled. If pulled a value from a quit channel, then exit the infinite loop -> need `NOT` to close the channel

### Channels - Fanin & Fanout
- `Fanin`: is a concept of pulling values from multiple Goroutines of receive channels and sending those values into 1 single receive channel.
  - Tood Code example: https://go.dev/play/p/_CyyXQBCHe
  - Rob Pike’s code example: https://go.dev/play/p/buy30qw5MM
- `Fanout`: taking a chunk of work, send to multiple Goroutines to run in parallel, e.g. need to encode 1000 videos -> create 1000 Goroutines each encode a video.
  - Example 1: https://go.dev/play/p/iU7Oee2nm7
  - Example 2 - throttle to just 10 Goroutines: https://go.dev/play/p/RzR3Kjrx7q

### Channels - `Context` - Advance topic
- `Context`: Often you have a process that spawn other Goroutines. Sometimes you shutdown the process, you want the Goroutines spawned by the process to be canceled so no leaked resources happen. `context` helps you achieve that
  - Further read:
    - https://go.dev/blog/context
    - https://medium.com/@matryer/context-has-arrived-per-request-state-in-go-1-7-4d095be83bd8
    - https://peter.bourgon.org/blog/2016/07/11/context.html
- For example, you can spawn a Goroutine has cxt.Done() so you can cancel the Goroutine from the main thread! 
```
go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
        // Do something
        // Send output or error checking into a channel c
        // In the main thread, pull output from channel c and call cancel() in the main thread to cancel this Goroutine IF you spot error!!!!
			}
		}
	}()
```
- You can spawn MULTIPLE of the above Goroutines and send output in the the same channel. A single cancel() will cancel all of those GoRoutines!

# Exported Functions - Struct fields - Struct
- The first character is capitalized if you want to export the functions, Struct Field or Struct

# ERROR HANDLING
- Code sample: https://github.com/GoesToEleven/go-programming/tree/master/code_samples/006-error-handling
- A `error interface` has type Error(). Type `errorString` that contains method `Error()` whichs returns a `string` => Type `errorString` implements the `error interface`
- `errors.New(string)` produces a type `errorString` which implement the `error interface`
- Any type that has method `Error()` implements type `error interface`

- From package built-in/errors https://pkg.go.dev/builtin#error:
  ```
    type error interface {
      Error() string
    }
  ```
- From source code of `errors.New` function
  ```
    package errors

    // New returns an error that formats as the given text.
    // Each call to New returns a distinct error value even if the text is identical.
    func New(text string) error {
      return &errorString{text}
    }

    // errorString is a trivial implementation of error.
    type errorString struct {
      s string
    }

    func (e *errorString) Error() string {
      return e.s
    }
  ```
## Error hanlding - Printing and Logging
- Similarly to logger in python, the below are recommended to be used which include timestamp:
  - log.Println()
  - log.Fatalln()
    - os.Exit()
  - log.Panicln()
    - deferred functions run
    - can use “recover”

- Other options you can use but less recommended
  - panic()
  - fmt.Println()

# WRITING DOC
## `go doc` cmd line - prefer documentation page
  - go help doc
```
    go doc <pkg>
    go doc <sym>[.<methodOrField>]
    go doc [<pkg>.]<sym>[.<methodOrField>]
    go doc [<pkg>.][<sym>.]<methodOrField>
```

  - Example:
    ```
            go doc json.Decoder.Decode
            go doc json.decoder.decode
            go doc json.decode
            cd go/src/encoding/json; go doc decode
    ```

## `godoc` - one word
- Install:  `sudo apt install golang-golang-x-tools`
- Run local http server to access the documentation
```
godoc -http=:8080
```
- Open web browser and access `http://localhost:8080`

## `pkg.go.dev` - HOW TO SHARE YOU CODE
- TODO: Find out how to share your code to `pkg.go.dev`

## Writing Documentation - convention
- Put comment right above package name, function (or identifier)
- Always start with the name of the package, function (or identifier) in the comment
- If the documentation for packge is lengthy, create a doc.go file in the package folder. See `fmt` package, scroll down to the bottom to see the doc.go file for example

# WRITING TESTS
- Tests must: 
  - be in a file that ends with “_test.go”
  - put the file in the same package as the one being tested
  - be in a func with a signature “func TestXxx(*testing.`T`)”
- `*testing.T`: pointer of type T in package `testing`: https://pkg.go.dev/testing#T  
  - Have methods Error(), Errorf(), Parallel(), ...
- Similarly, there are `*testing.B`, `*testing.F` which has a set of their own methods!

## Writing test - Example - 256-writing-test:03-writing-test-example
- `Example` are tests: https://pkg.go.dev/testing#hdr-Examples
- Write `Example` in test file

## Writing test - Golint, go vet, gofmt
- `golint``: reports poor coding style
- `go vet`: reports suspicious constructs
- `gofmt`: formats go code

## Benchmark - 256-writing-test:04-writing-test-benchmark
- Measure how fast your code runs.
- Write benchmark in test file
- `go test -bench . ./...` or `go test -bench <functionName> ./...`
  - `./...` means from this directory and all sub-directories
- `go help testflag`

## Coverage
- `go test -cover`
- `go test -coverprofile c.out`
- show in browser: `go tool cover -html=c.out`
- `go tool cover -h`
- `go test -cover -bench . ./...` : cover and bench mark at the same time

# Misc
- String is a slice of byte

# HTTP SERVER & CLIENT - FROM EFFECTIVE_GO - Interfaces and methods section
- Any object that implements Handler can serve HTTP requests.
```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
- `ResponseWriter` is itself an interface that provides access to the methods needed to return the response to the client. Those methods include the standard Write method, so an http.ResponseWriter can be used wherever an io.Writer can be used. 
- `Request` is a struct containing a parsed representation of the request from the client.
ResponseWriter is itself an interface that provides access to the methods needed to return the response to the client. Those methods include the standard Write method, so an http.ResponseWriter can be used wherever an io.Writer can be used. Request is a struct containing a parsed representation of the request from the client.

For brevity, let's ignore POSTs and assume HTTP requests are always GETs; that simplification does not affect the way the handlers are set up. Here's a trivial implementation of a handler to count the number of times the page is visited.
```
// Simple counter server.

type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```
(Keeping with our theme, note how Fprintf can print to an http.ResponseWriter.) In a real server, access to ctr.n would need protection from concurrent access. See the sync and atomic packages for suggestions.

For reference, here's how to attach such a server to a node on the URL tree.
```
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
```
But why make Counter a struct? An integer is all that's needed. (The receiver needs to be a pointer so the increment is visible to the caller.)
```
// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}
```
What if your program has some internal state that needs to be notified that a page has been visited? Tie a channel to the web page.
```
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")
}
```
Finally, let's say we wanted to present on /args the arguments used when invoking the server binary. It's easy to write a function to print the arguments.
```
func ArgServer() {
    fmt.Println(os.Args)
}
```
How do we turn that into an HTTP server? We could make ArgServer a method of some type whose value we ignore, but there's a cleaner way. Since we can define a method for any type except pointers and interfaces, we can write a method for a function. The http package contains this code:
```
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
```
`HandlerFunc` is a type with a method, ServeHTTP, so values of that type can serve HTTP requests. Look at the implementation of the method: the receiver is a function, f, and the method calls f. That may seem odd but it's not that different from, say, the receiver being a channel and the method sending on the channel.

To make ArgServer into an HTTP server, we first modify it to have the right signature.
```
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
```
`ArgServer` now has the same signature as `HandlerFunc`, so it can be converted to that type to access its methods, just as we converted Sequence to IntSlice to access IntSlice.Sort. The code to set it up is concise:
```
http.Handle("/args", http.HandlerFunc(ArgServer))
```
When someone visits the page `/args`, the handler installed at that page has value `ArgServer` and type `HandlerFunc`. The HTTP server will invoke the method ServeHTTP of that type, with `ArgServer` as the receiver, which will in turn call `ArgServer` (via the invocation `f(w, req) inside HandlerFunc.ServeHTTP`). The arguments will then be displayed.

In this section we have made an HTTP server from a struct, an integer, a channel, and a function, all because interfaces are just sets of methods, which can be defined for (almost) any type.