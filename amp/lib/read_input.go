package lib

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadInput() string {
	f, err := os.Open("input.html")
	if err != nil {
		fmt.Println("input error")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	return string(b)
}
