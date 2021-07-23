package main

import (
	"amp_html_edit/lib"
)

func main() {
	// fmt.Println(lib.GetHtmlDividedByTagAmp(lib.ReadInput()))
	lib.WriteOutput(lib.UnifyHtmlTokens(lib.GetHtmlDividedByTagAmp(lib.ReadInput())))
}
