**Array** 
* Collection of similar data elements stored at contiguous memory location.
* fixed length
* has a poiner and has a property called length.

**Decalaration of array**
* syntax: var \<array name> [size of the array] \<data type>
* example: var grades [5] int

**Array initialization**

* `var grade [3] int = [3]int{10,20,30}`
* `grades := [3]int{10,20,30}`

    * {elements} : the no. of elements should be lesser than the size of the array.
* `grades := [...]int{10,20,30}`

    * [...] : the three  dots are called ellipses.
    * It detects the array length automatically. No need to provide the length of the array.

**len()**: returns the no. of elements stored in the array. 

**indexes in array**: `0 <= index <= length-1`

**Looping through array**

* ```grades:=[...]int{90,80,85,99,100}

	for i:=0; i<len(grades); i++ {
		fmt.Println(grades[i])
	}

* **range**: This keyword is mainly used in for loops in order to iterate over all the elements of an array, slice or map.

    * It sets the scope of iteration upto the length of an array.

    ```
    for index, element := range grades {
        fmt.Println(index, "=>", element)
    }
    ```
    * 1st element: returns index
    * 2nd element: returns element

**Multidimensional Array**

* Syntax:

    * `arr := [3][2]int{{1,2},{3,4},{5,6}}`
    * [3]: represents size of the main array
    * [2]: represents size of the sub array

