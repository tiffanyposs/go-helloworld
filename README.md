# GO

TK Course link

## Notes

**How do we run the code in our project?**

```
$ go run main.go

```

* `$ go build` - compiles a bunch of go source code files
* `$ go run` - compiles and executes one or two files
* `$ go fmt` - formats all code in each file in current directory
* `$ go install` - compiles and "installs" a package
* `$ go get` - downloads raw source code of someone else's package
* `$ go test` -  runs any tests associated with the current project


```go

package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

```

**What does 'package main' mean?**

Declaring `package <name>` at the top of a file indicates which package this particular file belongs to within a project.

There are two types of packages in Go: `executable` and `reusable`

* `Executable` - generates a file that we can run.
* `Reusable` - code used as "helpers". Good place to put reusable logic.

**Note:** Naming the package "main" makes it an executable file.

**Note:** The file with the package "main" must have a function named "main"

**What does 'import "fmt"'mean?**

`import fmt` is simply importing the "format" package (included with standard Go library). using `import` is how you include packages into a file.

Other standard Go packages include:

* `fmt`
* `debug`
* `math`
* `encoding`
* `crypto`
* `io`

You can see a list and documentation for all packages here: https://golang.org/pkg/ 

**What's that 'func' thing?**

`func` is just short for `function`

**How is the main.go file organized?**

Every file in Go will have the same organization

1. Package declaration
2. Import other packages needed
3. Declare functions and tell Go to do things
