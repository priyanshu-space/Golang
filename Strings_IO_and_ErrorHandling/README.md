In Go, packages come bundled together as a library called "Standard Library"

**String**

* In Go, a string is a `read-only slice of bytes` that represents sequence of characters encoded in UTF-8.
* It is immutable.
* len(string) gives number of bytes not characters.

* byte vs rune

    * Use byte when dealing with low-level binary data or single-byte ASCII characters.
    * Use rune when working with individual characters in a string, especially when handling text that includes non-ASCII or multi-byte Unicode characters.

* Some string methods:

    * Contains
    * ReplaceAll
    * Count

go doc <package_name> <function_name>
eg; `go doc fmt Println`

**File Handling**

Packages : ??

* For metadata from a specific file : `stat` method from OS package




