// Package testsdk is a sample package to demonstrate some capabilities of gomobile
package testsdk

import (
	"errors"
	"fmt"
)

// anything exported (capitalized), is accessible to the mobile platform (and other Go packages)
// like a constant
const Title = "GoMobile Awesomeness"

// or a struct that will be available to the mobile platform
type MathResult struct {
	Sum        int32
	Difference int32
	Product    int32
	// this value is not directly accessible by the target platform
	// just like it wouldn't be directly accessible to any other Go package
	divisor int32
}

// we can define methods for the struct that show up as methods in the generated class on the mobile platform
func (r *MathResult) Compute(one int32, two int32) string {
	r.Sum = one + two
	r.Difference = one - two
	r.Product = one * two
	r.divisor = one / two
	return fmt.Sprintf("+: %d\n-: %d\n*: %d\n/: %d", r.Sum, r.Difference, r.Product, r.divisor)
}

// unexported attributes are still accessible if we provide a getter
func (r *MathResult) Divisor() int32 {
	return r.divisor
}

// we can also export a function that returns 0 or 1 values (it can have a second return value only if the value is an error)
func SayHello(to string) (string, error) {
	if to == "" {
		return "", errors.New("I will not greet someone with no name.")
	}

	return fmt.Sprintf("Hello, %s!", to), nil
}

// if we return a struct from a function it has to be a pointer (my guess is this is due to all objects in Java being references)
func MathResultPointer() *MathResult {
	return &MathResult{Sum: 0, Difference: 0, Product: 0, divisor: 1}
}

// this won't be available to the mobile platform
func MathResultInstance() MathResult {
	return MathResult{}
}

// we can store values to be retrieved later (probably not a great idea)
var savedVal int32

func SaveVal(val int32) {
	savedVal = val
}

func RetrieveVal() int32 {
	return savedVal
}

// exported interfaces are generated on the mobile platform and can be implemented by a Go struct or the native language
// on the mobile platform (Go structs implicitly implement interfaces if all of the required methods are defined)
type Result interface {
	Compute(int32, int32) string
}

func PrintResult(res Result, one int32, two int32) string {
	return "Result: \n" + res.Compute(one, two)
}
