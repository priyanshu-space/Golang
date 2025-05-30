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
   * `var s, t string = "Lia", "Harry"`
* Shorthand way
   * var (

     s string = "foo"

     i int = 5) 

* Other shorthand way
   * `s := "Hello World"` `//strictly to use this style inside func only.`
 
* local v/s global variable
  * local: declared inside a function or a block, can't access outside the function or block.
  * global: declared outside of a function or a block, available throughout the lifetime of a program.

* Zero values: When you don't assign any value to the variable during declaration it chooses some default values.
   * bool: false
   * int: 0
   * float64: 0.0
   * string: ""
   * pointer, func, interfaces, maps: nil

**Read User Input**
* Scanf : takes user input
   * Returns:
     * count: the number of arguments that the function writes to.
     * err: any error thrown during the execution of the function.

**variable type**
* By using format specifier: %T
* From reflect package, fun: reflect.TypeOf()

**Type casting in Golang**
* String to int: 
   * using `Atoi` inside `strconv` package.
   * **Atoi**: converts `string to integer` and return two values- the corresponding integer, error(if any).
* int to string: 
   * using `Itoa` inside `strconv` package.
   * **Itoa**: converts `integer to string` and return one value- string formed with the given integer. 
* one type to another: `var f float64 = float(i) //i is some int variable`

**Constant**
   * Syntax: `const <variable-name> <data-type> = <value>`
   * Types: `untyped` and `typed`
   * untyped: `const age = 12`
   * typed: `const name string = "Harrry"`
   * Note: 
      * constants once assigned its value can't be changed. 
      * constant's value should be assigned and declare at the same time(not later).
      * shorthand declaration for constant wouldn't work. `const name := "Harry" // throws error`

**Operators** <Notes Pending>

* **Bitwise Operators**
   * bitwise AND (&): Converts the operand values in to bits and then perform the AND operation b/w the operands. 
   * bitwise OR (|): Converts the operand values in to bits and then perform the OR operation b/w the operands.
   * bitwise XOR (^): Output of XOR will be `1 if both bits are different` and `0 if both bits are same`. 
   * left shift (<<): Shifts all bits towards left and the bit position vacated are filled with 0.
   * right shift (>>): Shifts all bits towards right and the excess bits shifted off to the right are discarded.

**if-else**

* syntax:

```
if condition_1 {
   //execute when condition 1 is true
} else if condition_2 {
   //execute when condition 1 is false but condition_2 is true
} else if condition_3 {
   //execute when condition 1 and 2 is false but condition_3 is true
} else {
   //execute when none of the condition is true.
}
```
* Note : the `else` block will start where the `if` block ends(otherwise it will throw error).

**switch-case**

* syntax:

```
switch expression {
   case value_1:
   //execute when expression equals to value_1
   case value_1:
   //execute when expression equals to value_2
   case value_1:
   //execute when expression equals to value_3
   default:
   //execute when no match found
}
```
* Note: Go uses implicit break statement for switch-case unlike other languages(like c,java).

**Loops in Go**
* syntax-1:

```
for i := 1; i <= 5; i++ {
   fmt.Println(i)
}
```
* syntax-2:

```
i := 1
for i <= 5 { //loop only with condition
   fmt.Println(i)
   i++
}
```
* syntax-3:

```
i := 1
for {
   fmt.Println(i)  //infinite loop
}

```
