**Why Functions**

* Reusability
* Abstraction

**Declaration**

* Syntax: `func <func_name>(<params>) <return type> { // body of the function}`

**Return**

```
func addNumbers(a int, b int) int {
    sum := a+b
    return sum
}
```

**Calling a function**

`<function_name>(<arguments(s))`

**Parameters Vs Arguments**

* Parameters: names listed in functions's defination.
* Arguments: Real values passed into the function.

**Variadic functions**

* function that accepts variable number of arguments.
* to declare a variadic function, the type of the final parameter is preceded by an ellipsis **"..."**.

**Declaration of variadic function**

* `func <func_name>(param-1 type, param-2 type, param-3 type) <return_type>`
* `func sumNumbers(numbers ...int) int`

    * all the integer passed will be stored in **slice** numbers.

**Recursion**

* Recursion is a concept where a function calling itself by direct or indirect means.
* The function keeps calling until it reaches a base case.
* Used to solve a problem where solution is dependent on the smaller instance of the same problem.

**Anonymous functions**

Anonymous functions (also called function literals) are functions `defined without a name`. 
* Can Call them immediately.
* Or assign them to a variable for later use.
* Or pass them as arguments to other functions.

**High Order Functions**

* function that receives a function as an argument or returns a function as output.
* Why

    * composition:

**Defer Statement**

* A defer statement delays the execution of a function until the surrounding function returns.