/**
Writing tests:

Writing a test is just like writing a function, with a few rules
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
- For now, it's enough to know that your t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail.
**/

package main

import "testing"
import "fmt"

func TestHello(t *testing.T) {

  t.Run("says hello to user", func(t *testing.T) {
    user := "Bryan"
    got := Hello(user)
    want := fmt.Sprintf("Hello, %s", user)

    if got != want {
      t.Errorf("Got %q want %q", got, want)
    }
  })

  t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
    got := Hello("")
    want := "Hello, World"

    if got != want {
      t.Errorf("Got %q want %q", got, want)
    }
  })
}
