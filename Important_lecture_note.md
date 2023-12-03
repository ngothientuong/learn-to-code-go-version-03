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

- When reading package'S index page
  - Under list of type, such as `type Decoder`: https://pkg.go.dev/encoding/json#pkg-index 
  - The first function usually guide you how you can get a function of such type, e.g `*Decoder` 
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

  """