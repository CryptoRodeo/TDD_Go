/**
Notes: 

- When you write a program in Go you will have a main package defined with a main func inside it. 
- Packages are ways of grouping up related Go code together.
- The func keyword is how you define a function with a name and a body.
- With import "fmt" we are importing a package which contains the Println function that we use to print.

Testing notes:
- It is good to separate your "domain" code from the outside world (side-effects). 
- The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
**/

package main

import "fmt"

func Hello(name string) string {
  return fmt.Sprintf("Hello, %s", name)
}

func main() {
  fmt.Println(Hello("Worled"))
}
