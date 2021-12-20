# Variable

 - The var statement declares a list of variables; as in function argument lists, the type is last.

 - A var statement can be at package or function level. We see both in this example.

 - When a variable is declared, and not explicitly initialized, it will be allocated storage with a default value, appropriate for the type, known as the type’s zero value

```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

Output :

```
0 false false false
```

url : [Variable](https://play.golang.org/p/XzV619chZOD)

## Variables with initializers

 - A var declaration can include initializers, one per variable.

 - If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

 - variable in go can only be assign with the same type

```
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

Output :

```
1 2 3 true false no!
```

url : [Variables with initializers](https://play.golang.org/p/Pvvcjdiyz9y)

## Short variable declarations

 - Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

 - Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

Output :

```
1 2 3 true false no!
```

url : [Short variable declarations](https://play.golang.org/p/Zd7IrcAyrHh)

 - The example shows variables declarations may be "factored" into blocks.

```
package main

import "fmt"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
}
```

Output :

```
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
```

url : [Short variable declarations](https://play.golang.org/p/BlYSduUuBOE)

## Variabel Underscore `_`

 - Unused variable not allowed in go

 - So to receive unused variable from return function usually using `_` variable

 - `_` is predefined variable, no need to declare it, you cannot use the variable because no data will be store there

```
package main

import "fmt"

var c, python, java bool

func main() {
	var i, _ int
	fmt.Println(i, c, python, java)
}
```

Output :

```
0 false false false
```

url : [Variable _](https://play.golang.org/p/kXQwPsPGRIB)