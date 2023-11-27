# Learn To Code Google's Go (golang) Programming Language

- This code repository is used to **learn the Go programming language**. 
- This is the code for Todd McLeod's "Learn To Code Go" course. 
- There are video lectures associated with this codebase
- **You can take the course at [mcleods.com](https://www.mcleods.com)**
  
🌞🌴😃

# Important Note from Tim

## Pointer

- `*int`: type pointer to an int
-  `*n`: defererence `address n` to get a value
- `&a`: address of `value a`

## Reference type
- A reference type is pointer

- `Reference Type` includes: interface, slices, channels, functions, pointers, maps
- Remember the phase **a`I` `S`ẽ `C`ó `F`ần `P`hở `M`ập**

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