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
