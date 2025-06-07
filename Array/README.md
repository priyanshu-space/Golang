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

**Slice**

* ```
    arr := [...]int{10,20,30,40,50,60}
	
    slice_1 := arr[1:3] // index 1 is included but 3 is excluded

    Output: [20 30]

* [start_index:end_index]
* [:] : start from index 0 and goes till the last index.
* [:3] :  start from index 0 and goes till index 2.
* Sub slice can also created from an underlying slice.
* We can also declare a slice using **make** function.

    * syntax: `slice := make([]\<data_type>, length, capacity(optional))`
    * example: `slice := make([]int, 5, 10)`

* **cap()**: To calculate the capacity of an Slice or Array.
    
    * The capacity of the array = length of the array.
    * The capacity of the slice = length of the underlying array - starting slice starting index(array index).
    * example  : 
    ```
    arr := [4]int{10,20,30,40}
    slice := arr[1:3]
   
    len of slice = 2 (number of elements in slice)

    capcity of slice = 3 (iterate from 20 till end of the array)   

**Appending to a Slice**

* Function: `func append(s []T, \<values to append of same data type>) []T`

    * s []T: slice of an array.

    * Rest are the values to append.
    * []T : New Slice containing all the values (original slice + appended values).
    * **Note: If the initial capacity could not contain the appended values, then the newer slice generated will be the capacity of 2X the original slice. 

* `slice = append(slice, element-1, element-2)`

**Appending a slice to a new slice**

* `slice = append(slice, anotherSlice...)`

    * Note: These **...** are used for variadic function that can take arbitary no. of arguments as well.
    * anotherSlice would be appended to the first slice.

**Copying from a slice**

* `num := copy(dest_slice, src_slice)`


