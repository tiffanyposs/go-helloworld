// declares which package this file is apart of
// more than one file can be part of a package
// note that naming it main indicates that we are
// creating a executable file
package main

// gives our package access to some code
// written inside of another package
// "fmt" is a standard package that comes with Go
// it is short for "format", used to print out information
import "fmt"

// since "package main" creates an executable filel
// it requires a funtion named "main" to be created
// this is the function that will be automatically
// called on execution
func main() {
	fmt.Println("Hi there!")
}
