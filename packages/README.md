# Go Packages

Go programs are organized into packages. It is a collection of Go source file in the same directory. We can define functions, variable, constant, or types inside those package and use it in another source files.

## Main Package

The most important package is the `main` package. It is reserved package name where our program start running and meant to be executable. Inside `main` package, there will be `main()` function as the entrypoint for our program. 

```go
// File: src/example_1/main.go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")
}
```

If we run the example above like this:

```shell
$ go run main.go
```

unsuprisingly, this is what we will get:

```
Hello World
```

## Importing Package 

We can import package using `import` followed by the package name. For multiple import, we can write `import <package-name` multiple times or group it using parenthese, known as "factored" import.

```go
import "fmt"
import "math"
import "net"
```

```go
import (
    "fmt"
    "math"
    "net"
)
```

Both are valid, but it is common to use "factored" import instead of using multiple import.

## Names Visibility

When we import a package, only the exported names are visible or can be accessed. In Go, any variable, type, or function that start with capital letter is exported and can be accessed from another package. If it does not start with capital letter, it can only be accessed inside the same package.

Let's see an example for this.

```go
// File: src/example_2/main.go
import (
    "fmt"
    "math"
)

func main() {
    // This is OK
    fmt.Println("Call Exp: ", math.Exp(1.0))

    // This will give us an error
    fmt.Println("Call expmulti: ", math.expmulti(0.1, 0.1, 1))
}
```

If we run the above code, we get an error like this:

```shell
$ go run main.go
# command-line-arguments
./main.go:13:33: cannot refer to unexported name math.expmulti
```

If we go to https://golang.org/src/math/exp.go, both `Exp` and `expmulti` can be found in `math` package. But, we get an error line 13 (`math.expmulti`) that said `cannot refer to unexported name`. Even though `expmulti` is defined in `math` package, this function name doesn't start with capital letter, so it will not visible to the outside of `math` package.

Comment those line (line 13) and now we can successfully compile and run the program

```shell
$ go run main.go
Call Exp:  2.718281828459045
```

## Create Custom Package

We already used two package, `fmt` and `math`. It is standard package provided by Go standard library, which we can find out more information about it in https://golang.org/pkg/. We can also create our own package and use it in our project. Let's try to create a simple Go project and also create packages inside it.

First, create a new folder. This folder will be our root project. 

```shell
$ mkdir customod
```

Go to that folder and initialize a Go module using `go mod init` command

```shell
$ cd customod
$ go mod init local.host/example/customod
```

Initializing module will create `go.mod` file in the root project. This file will have our module name that we specified (`local.host/example/customod`) and also a list of dependency for this module. Go module will be explained more later. 

Now, create some package and source file inside this folder. Below is the structure of this custom module that we will create.

```
customod
-- math
   -- list
      -- list.go
   -- arithmetic.go
-- text
   -- reverse.go
-- main.go
```

Go to [src/example_3](src/example_3/) to see the complete source code for this example. 

If we run it, we will get this output:

```shell
src/example_3/customod$ go run main.go
Add number: 1.5 + 2.5 =  4
Multiply number: 1.5 + 2.5 =  3.75
Constant: PI:  3.141592653589793
Add all of the number in array: sum([3.0, 4.0, 1.0, 2.5]) =  10.5
Multiply all of the number in array: mul([3.0, 4.0, 1.0, 2.5]) =  30
Reverse: rev(Hello World) =  dlroW olleH
```

### Explanation

- Package path

    The import to `local.host/example/custommod` refer to this module. So to import the package inside this module we just add the name after our module's name. 

    Package can be nested too. Our `list` package is inside `math` package, so to import `list` package we just need to give the path to this package after our module's name. 

    ```go
    // File: src/example_3/customod/main.go
    import (
        ...

        "local.host/example/customod/math"
        "local.host/example/customod/math/list"
        "local.host/example/customod/text"
    )

    ...
    ```

- Alias

    To avoid conflict between different package with similiar name, we can use alias. In this example, our custom `math` package have the same name as `math` package in Go standard libary.

    ```go
    // File: src/example_3/customod/main.go
    import (
        ...
        gomath "math"

        ...
    )
    ```

## 3rd Party Library

To add 3rd party library to our project, we can use Go modules. We can import the path to the remote package in our program, then use `go get <module_path>` to download it. After that, `go.mod` will be updated to include the path to that library. 

For example, we have a program that generate a dummy/fake data. We use library `github.com/pioz/faker` in this program.

```go
// File: src/example_4/randomtext
package main

import (
	"fmt"

	"github.com/pioz/faker"
)

func main() {
	faker.SetSeed(623)

	fmt.Println(faker.Username())
	fmt.Println(faker.String())
	fmt.Println(faker.IntInRange(1, 10))
	fmt.Println(faker.CurrencyName())
}
```

Use `go get` to download `github.com/pioz/faker` from remote repository

```shell
src/example_4/randomtext$ go get github.com/pioz/faker
go get: added github.com/pioz/faker v1.7.0
```

If we look at `go.mod`, we can see that it add this new line

```
// File: src/example_4/randomtext/go.mod

require github.com/pioz/faker v1.7.0 // indirect
```

Now run our program, and we can see something like this:

```shell
/src/example_4/randomtext$ go run main.go 
spicule
g
2
Samoan Tala
```

## Vendoring

When we download a package using `go get`, the package will be stored locally in `GOPATH/pkg/mod` and then it will be pulled by the program. Since the example projects above didn't include any 3rd party dependency in the project folder, only the path in `go.mod`, this might be a problem if we want to share those project. Because the project might use a library from private repository or maybe in the future the package that we use in our project suddenly disappear or deleted from remote repository.

To solve that problem we can use Go vendor. It will create a folder named `vendor` in our project root and store all of the dependency that we use inside it.

Using the same program as previous example, run `go mod vendor` from the project root.

```shell
src/example_5/vendoring$ go mod vendor
```

We'll see a `vendor` folder created. Inside it is the package used by this project.

When the project get larger, we might use a lot of libraries. Sometimes the dependency is a mess, unused library/package is still included in `vendor` folder or `go.mod`. So, it is recommended to run `go mod tidy` to cleans up these unused dependencies.


## Further Readings

- https://www.callicoder.com/golang-packages/ (The structure of this document is based on this post)
- https://blog.golang.org/using-go-modules (Information about how to use Go modules)
- https://golang.org/ref/mod (More detailed information about Go modules)
- https://blog.golang.org/package-names (Information about how to name a package)
- https://golang.org/doc/code (How to write Go code)

