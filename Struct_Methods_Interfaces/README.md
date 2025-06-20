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