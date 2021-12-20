# Control Flow

## If-Else Statement


### If Statement

Go's `if` statement is structured like this:

```go
if (condition) {
    // code
}

// OR

if condition {
    // code
}
```

Condition expression need not be surrounded by parentheses `( )` but the braces `{ }` are required.

```go
// File: src/if-else/01_if-statement/main.go
package main

import "fmt"

func multipleOfTen(x int) bool {
	if x%10 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("20 is multiple of 10? ", multipleOfTen(20))
	fmt.Println("12 is multiple of 10? ", multipleOfTen(12))
}
```

The output of above code:

```shell
$ go run main.go 
20 is multiple of 10?  true
12 is multiple of 10?  false
```

### If-Else Statement

An `if` statement can be combined with an `else` which is executed if the condition specified in the if statement is `false`

```go
if condition {
    // execute this block if condition is true
} else {
    // execute this block if condition is false
}
```

Example of using `if` and `else` statement

```go
// File: src/if-else/02_if-else-statement/main.go
package main

import "fmt"

func main() {
	number := 21
	if number%10 == 0 {
		fmt.Printf("%d is multiple of 10\n", number)
	} else {
		fmt.Printf("%d is not multiple of 10\n", number)
	}
}
```

Output:

```shell
$ go run main.go 
21 is not multiple of 10
```

### If-Else-If Chain

Multiple `if` statement and making a chain of `else if` like this:

```go
// File: src/if-else/03_if-else-if-statement/main.go
package main

import "fmt"

func grading(point int) string {
	if point >= 90 {
		return "Excellent"
	} else if point >= 80 && point < 90 {
		return "Good"
	} else if point >= 70 && point < 80 {
		return "Pass"
	} else if point >= 50 && point < 70 {
		return "Unsatisfactory"
	} else {
		return "Fail"
	}
}

func main() {
	fmt.Println("91: ", grading(81))
	fmt.Println("72: ", grading(72))
	fmt.Println("50: ", grading(50))
	fmt.Println("10: ", grading(10))
}
```

Output:

```shell
$ go run main.go
91:  Good
72:  Pass
50:  Unsatisfactory
10:  Fail
```

### Short If Statement

The `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the `if` (or until end of `else` if using if-else). 

```go
// File: src/if-else/04_if-short/main.go
package main

import "fmt"

func main() {
	if number := 21; number % 10 == 0 {
		fmt.Printf("%d is multiple of 10\n", number)
	} else {
		fmt.Printf("%d is not multiple of 10\n", number)
	}

    // var number is no longer available here
}
```

## Switch

A switch statement is a shorter way to write a sequence of `if - else` statements. It runs the first case whose value is equal to the condition expression.

Unlike `switch` in other language, like C or C++, Go only runs the selected case, not all the case that follow. Also, `break` statement at the end of each case in those languages is provided automatically in Go.

Here's a simple example of switch statement:

```go
// File: src/switch/01_switch/main.go
package main

import "fmt"

func main() {
	var dayOfWeek = 6
	switch dayOfWeek {
		case 1: fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednesday")
		case 4: fmt.Println("Thursday")
		case 5: fmt.Println("Friday")
		case 6: {
			fmt.Println("Saturday")
			fmt.Println("Weekend!")
		}
		case 7: {
			fmt.Println("Sunday")
			fmt.Println("Weekend!")
		}
		default: fmt.Println("Invalid day")
	}
}
```

Output:

```shell
$ go run main.go 
Saturday
Weekend!
```

### Short switch

Similiar to `if`, `switch` can also have a short statement before the conditional expression.

The previous example can be written like this

```go
// File: src/switch/01_switch-short/main.go
package main

import "fmt"

func main() {
	switch dayOfWeek := 6; dayOfWeek {
		case 1: fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednesday")
		case 4: fmt.Println("Thursday")
		case 5: fmt.Println("Friday")
		case 6: {
			fmt.Println("Saturday")
			fmt.Println("Weekend!")
		}
		case 7: {
			fmt.Println("Sunday")
			fmt.Println("Weekend!")
		}
		default: fmt.Println("Invalid day")
	}
}
```

Output:

```shell
$ go run main.go 
Saturday
Weekend!
```


### Multiple switch case

Multiple cases can be combined into one like this:

```go
// File: src/switch/01_switch-multiple-case/main.go
package main

import "fmt"

func main() {
	switch dayOfWeek := 6; dayOfWeek {
		case 1, 2, 3, 4, 5: fmt.Println("Weekday")
		case 6, 7: {
			fmt.Println("Weekend!")
		}
		default: fmt.Println("Invalid day")
	}
}
```

Output:

```shell
$ go run main.go 
Weekend!
```

### Switch with no condition

Switch without a condition is the same as `switch true`. This can be used to write long `if-else-if` chain in more clean way.

```go
// File: src/switch/01_switch-no-condition/main.go
package main

import "fmt"

func grading(point int) string {
	switch {
        case point >= 90:
            return "Excellent"
        case point >= 80 && point < 90:
            return "Good"
        case point >= 70 && point < 80:
            return "Pass"
        case point >= 50 && point < 70:
            return "Unsatisfactory"
        default:
            return "Fail"
    }
}

func main() {
	fmt.Println("91: ", grading(81))
	fmt.Println("72: ", grading(72))
	fmt.Println("50: ", grading(50))
	fmt.Println("10: ", grading(10))
}
```

Output:

```shell
$ go run main.go 
91:  Good
72:  Pass
50:  Unsatisfactory
10:  Fail
```

## For Loop

Go has only one looping statement, the `for` loop.

This is the structure of `for` loop:

- initialization is executed before the first iteration,
- condition expression is evaluated before every iteration, and
- post statement is executed at the end of every iteration

```go

for initialization; condition_expression; post_statement {
    // loop body
}

```

Here's a simple example for `for` loop

```go
// File: src/for-loop/01_for-loop/main.go
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }

    fmt.Println(sum)
}
```

It will output:

```shell
$ go run main.go
45
```

Initialization, condition, and post statement is optional. It is possible to just write only one of them or even without any of them.

### Omitting initialization

```go
// File: src/for-loop/02_for-loop-without-init/main.go
package main

import "fmt"

func main() {
	i := 0
	for ; i < 10; i += 2 {
		fmt.Printf("%d ", i)
	}
}
```

This is the output

```shell
$ go run main.go
0 2 4 6 8
```

### Omitting condition
 
```go
// File: src/for-loop/03_for-loop-without-condition
package main

import "fmt"

func main() {
	for j := 0; ; j++ {
		fmt.Printf("%d ", j)
		if j == 10 {
			break
		}
	}
}
```

This is the output

```shell
$ go run main.go
0 1 2 3 4 5 6 7 8 9 10 
```

### Ommitting post statement

```go
// File: src/for-loop/04_for-loop-without-post-statement/main.go
package main

import "fmt"

func main() {
	i := 0
	for ; i < 10; {
		fmt.Printf("%d ", i)
        i += 2
	}
}
```

This is the output

```shell
$ go run main.go
0 2 4 6 8
```

Semicolon (`;`) can be omitted from the `for` loop without initialization and post statement, like this:

```go
package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
        i += 2
	}
}
```

### Infinite loop

Omitting initialization, condition, and post statement will give us an infinite loop

```go
package main

func main() {
	// Infinite Loop
	for {

	}
}
```

### Break and continue

- `break` statement is used in `for` loop to break out of the loop.
- `continue` statement is used in `for` loop to continue to the next iteration.

Here's the example:

```go
package main

func main() {
	for n := 0; n < 100; n++ {
        if n == 10 {
            break
        }

        if n % 2 == 0 {
            continue
        }

        fmt.Printf("%d ", n)
	}
}
```

The output will be:

```shell
$ go run main.go
1 3 5 7 9 
```

### Range

The `range` form of the `for` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```go
package main

import "fmt"

func main() {
	numbers := []int{2, 5, 7, 19, 20, 30}

	for idx, val := range numbers {
		fmt.Printf("%d = %d\n", idx, val)
	}
}
```

Output:

```shell
$ go run main.go
0 = 2
1 = 5
2 = 7
3 = 19
4 = 20
5 = 3
```

`range` provides both index and value for each entry. If we don't need one of them we can ignore it with `_`.

```go
for _, val := range numbers {
	fmt.Printf("%d\n", val)
}
```

`range` can also be used to iterate over maps. Like this:

```go
package main

import "fmt"

func main() {
	// Map iteration
	keyvalue := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, val := range keyvalue {
		fmt.Printf("key %s = %d\n", key, val)
	}
}
```

Output:

```shell
$ go run main.go
key one = 1
key two = 2
key three = 3
```


## Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 

```go
// File: src/defer/01_defer/main.go
package main

import "fmt"

func main() {
	defer fmt.Println("Executed last")

	fmt.Println("First!")
}
```

Output:

```shell
$ go run main.go 
First!
Executed last
```

There are 3 simple rules for `defer`:
1. A deferred function's arguments are evaluated when the defer statement is evaluated

    Expression `i` in example below is evaluated when `Println` is deferred. This function will print `0` after the function returns.

    ```go
    func deferEvaluation() {
        i := 0
        defer fmt.Println(i)
        i += 2
    }
    ```

2. Deferred function calls are executed in Last In First Out (LIFO) order after the surrounding function returns.

    ```go
    func deferOrder() {
        defer fmt.Println("First defer")
        defer fmt.Println("Second defer")
        defer fmt.Println("Third defer")
    }
    ```

    Those function will print from Third to First

    ```
    Third defer
    Second defer
    First defer
    ```

3. Deferred functions may read and assign to the returning function's named return values.

    From this example, a deferred function increments the return value `i` after the surrounding function returns. Thus, this function returns 5:

    ```go
    func deferReturn() (i int) {
        i = 3
        defer func() { i += 2 }()
        return i
    }
    ```

