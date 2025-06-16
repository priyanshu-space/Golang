**Pointers**

* A pointer is a variable that holds memory address of another variable.
* They point where the memory is allocated and provide ways to find or even change the value located at the memory location.

**Operator**

* & operator - The address of a variable can be obtained by preceding the name of a variable with an ampersand sign (&), known as address-of-operator.

* \* operator - It is known as the dereference operator. When placed before an address, it returns the value at that address.

**Decalaring and Initialization**

* Declaring: `var <pointer_name> *<data_type>`

    * example: `var ptr_i *int`

* Initilization: `var <pointer_name> *<data_type> = &<variable_name>` or `var <pointer_name> = &<variable_name>` //data type will internally determined by the compiler.

    * example: `var ptr_i *int = &i` or `var ptr_i = &i`
    * shorthand: `ptr_i := &i`

**Passing by value v/s passing by reference**

* Passing by value: parameter is copied into another location of your memory. Accessing or modifying the value will not change the original value.

* Passing by reference: the address of the variable is passed into the function call as the actual parameter. Modifying the value can change the original value.

    * Slices and maps are passed by reference by default.