**Struct**

* user-defined data type, groups data of different data type together.
* Declaration:
    
    * Syntax: `type <struct_name> struct {}` 
    * example: 
    ```
    type Student struct {
        name string
        rollno int
        marks []int
        grades map[string]int
    }
    ```
* Initialization:
    
    * `var <variable_name> <struct_name>`
        
        * `var s Student`
    * `<variable_name> := new(<struct_name>)`
        
        * `st := new(Student)` 
    * 
    ```
    <variable_name> := <struct_name> {
    <field_name>: <value>,
    <field_name>: <value>,
    }

    //example

    st := Student {
        name: "joe",
        rollNo: 12,
    }

    // or

    st := Student{"joe", 12}
    ```

**Method**

* A method adds an extra section immediately after the `func` keyword that accepts a single argument. This argument is called a `receiver`.
* **A method is a function that has a defined receiver**.
* Syntax: `func (<receiver>) <method_name>(<parameter>) <return_params> { }`.
    
    * example: `func (c Circle) area() float64 { }`, `func (c *Circle) area() float64 { }`
    * Note : Any instance of Circle c has `area` method available.

**Interface**


------------

**Combined Analogy**

A. Struct — Model Data

```
type User struct {
    name  string
    email string
}

```

* Problem: You can store data, but not define any behavior

B. Method — Add Behavior to Struct

```
func (u User) Greet() {
    fmt.Println("Hello,", u.name)
}
```
* Problem: Now you can give `User` behavior but still can't write flexible code that works with any `struct` type.

C. Interface — Define a Contract

```
type Greeter interface {
    Greet()
}

//Now any type that implements Greet() becomes a Greeter.

func sayHello(g Greeter) {
    g.Greet()
}
```
**Unified Example**

```
package main

import "fmt"

// Struct
type User struct {
    name  string
    email string
}

// Method
func (u User) Greet() {
    fmt.Println("Hello,", u.name)
}

// Interface
type Greeter interface {
    Greet()
}

// Function using interface
func sendGreeting(g Greeter) {
    g.Greet()
}

func main() {
    u := User{name: "Priyanshu", email: "you@example.com"}
    sendGreeting(u) // uses interface to call Greet
}

```