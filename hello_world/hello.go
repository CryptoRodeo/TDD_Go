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

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const frenchHelloPrefix = "Bonjour, "
const french = "French"

func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }

  return greetingPrefix(language) + name
}

/**
The function signature has a named return value (prefix string).
-> This will create a variable called prefix in your function.
-> You can return whatever it's set to by just calling return rather than return prefix.

It will be assigned the "zero" value. This depends on the type, for example ints are 0 and for strings it is "".
This will display in the Go Doc for your function so it can make the intent of your code clearer.
default in the switch case will be branched to if none of the other case statements match.

The function name starts with a lowercase letter. 
-> In Go public functions start with a capital letter and private ones start with a lowercase. 
We don't want the internals of our algorithm to be exposed to the world, so we made this function private.
**/
func greetingPrefix(language string) (prefix string) {
  switch language {
    case french: 
      prefix = frenchHelloPrefix
    case spanish:
      prefix = spanishHelloPrefix
    default:
      prefix = englishHelloPrefix
  }
  return
}

func main() {
  fmt.Println(Hello("World", "Spanish"))
}
