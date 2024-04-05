package main

import (
	"fmt"
	"errors"
)

type PersonError struct {
	height int
	weight int
	mess string
}

func (err *PersonError) Error() string {
	return fmt.Sprintf("%d and %d with %s", err.height, err.weight, err.mess)
}

func badPerson() (int, error) {
	return 5, &PersonError{height: 6, weight: 180, mess: "Charlie Brown"}
}

var ErrorBad = fmt.Errorf("Bad things")

func bad(thing int, thang int) (float64, error) {
	if thang == 0 {
		return -1, errors.New("Division by zero is not allowed") 
	}
	if thang == 2 {
		return -1, fmt.Errorf("You just asked for: %w", ErrorBad)
	}
	return float64(thing) / float64(thang), nil
}

func main() {
	thing, err := bad(5, 0)

	if err != nil {
		if errors.Is(err, ErrorBad) {
			fmt.Println("This was a", err)
		} else {
			fmt.Println("This was not a", err)
		}
	} else {
		fmt.Println(thing)
	}

	thing, err = bad(5, 2)

	if err != nil {
		if errors.Is(err, ErrorBad) {
			fmt.Println("This was a", err)
		} else {
			fmt.Println("This was not a", err)
		}
	} else {
		fmt.Println(thing)
	}

	_, eerr := badPerson()
	var pe *PersonError
	if errors.As(eerr, &pe) {
		fmt.Println(eerr.Error())
	}
}