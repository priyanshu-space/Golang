package encrypt

func Nimbus(str string) string { //to export a package identifier its name should start with a capital letter (nimbus --> won't work)

	encryptedStr := ""

	for _,c := range str {
		asciiCode := int(c)
		character := string(asciiCode+3)
		encryptedStr += character
	}
	return encryptedStr

}


/*
Note: If we want to export a package identifier its name should start with captilization(Nimbus), otherwise(nimbus) it would be scoped inside package only.

Golang uses capitalization to scope the package identifier.
*/


