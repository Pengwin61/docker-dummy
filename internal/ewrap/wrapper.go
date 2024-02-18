package ewrap

import "fmt"

func Err(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
