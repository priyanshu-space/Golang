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
**data Types**
 
* To get the details of memory allocation for the operation.

**Statically v/s Dynamically typed language**

* Statically
   * Data type: checked at compile time.
   * Variable declaration: Required or inferred, but fixed.
* Dynamically
   * Data type: checked at runtime.
   * variable declaration: Not required can change dynamically.

* Note
   * Go is a statically typed language.
   * Inferred : The compiler guessed the type of a variable based on the value you assigned.
      * `age := 25`, it is same as `var age int = 25`. Now `age = "twenty" // throws compile-time error` 
