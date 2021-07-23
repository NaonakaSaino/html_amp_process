package lib

import (
	"fmt"
	"io/ioutil"
)

func WriteOutput(html string) {

	err := ioutil.WriteFile("output.html", []byte(html), 0664)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("書き込み完了")
}
