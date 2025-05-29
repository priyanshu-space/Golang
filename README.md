# Golang
--------

**Basics**

Go is a compiled language.

* **Packages**

    * Every *.go* file starts with a *package* declaration.
    * *package main* is used to build the executable programs.
    * Without `package main`, Go won’t compile it as an executable.

* **import** 

    * Allow to use code from other packages.

    ```
    package main      // This file belongs to the 'main' package

    import "fmt"      // This file imports and uses the 'fmt' package

    func main() {
    fmt.Println("Hello, Go!")
    }
    ```
    
**Statically v/s Dynamically typed language**

* Statically
   * Data type: checked at compile time.
   * Variable declaration: Required or inferred, but fixed.
* Dynamically
   * Data type: checked at runtime.
   * variable declaration: Not required can change dynamically.

* Note :
   * Go is a statically typed language.
   * Inferred : The compiler guessed the type of a variable based on the value you assigned.
      * `age := 25`, it is same as `var age int = 25`. Now `age = "twenty" // throws compile-time error`

**Why Data Types**
* To get the details of memory allocation for the operation.

**Different Kinds of DataTypes** 
1. Numeric
   * Integer : `int`, `int8`, `int16`, `int32`, `int64`
   * Unsigned Integer : `uint`, `uint8`, `uint16`,`uint32`, `uint64`
   * Float : `float32`, `float64`
   * Complex : `complex64`, `complex128`
   * Alias : 
3. Boolean
4. String
5. Derived

**Declaring variables in Golang**

* `var <var-name> <data-type> = <value>`
   * `var s string = "Hello world"`
   * `var i int = 100`
   * `var b bool = true`
   * `var f float64 = 77.90`
* Shorthand way
   * var (

     s string = "foo"

     i int = 5) 

* Other shorthand way
   * `s := "Hello World"` `//strictly to use this style inside func only.` 
