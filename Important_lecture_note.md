# Important Note from Tim

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
  """
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
  """
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
- A value can be sent to a BUFFERED channel by main goroutine and the main goroutine can also receive from the channel later
  
- Birectional Channel:  `c := make(chan int)`
- Directional Channel, read from left to right:
  - `cr := make(<-chan int)` // receive
  - `cs := make(chan<- int)` // send