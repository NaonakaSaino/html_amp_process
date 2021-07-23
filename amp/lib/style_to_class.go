package lib

import (
	"regexp"
)

//以下はMoneyTimesのもの
var cvList = map[string]string{
	"#0044db": "_color-blue",
	"#0044cc": "_color-blue",
	"#FF0000": "_color_red",
	`padding *: *0px 7px 0px 7px *; *margin-bottom *: *0px *; *border *: *2px solid #d9d9d9 *; *background-color *: *#ececec *;`: "name-box-wrap",
	"center":  "_align-center",
	"justify": "_align-justify",
	`font-family *: *"ＭＳ 明朝" *, *serif *;`:                                                                                                                  "_mincho",
	`width *: *2.5em *; *padding-left *: *0.3rem *; *margin-top *: *0px *; *margin-bottom *: *0px *; *float *: *left`:                                       "_misc-01",
	"float *: *left *; *color *: *#a48b27 *; *font-size *: *1.2em *; *margin-top *: *.3em *; *margin-bottom *: *.3em *; *margin-right *: *calc(100% - 5em)": "_misc-02",
	"inline-block": "_inline-block",
}

//順番違いの場合の対応が必要になるかも。要修正

func ConvertStyleToClass(styles []string, classes []string) []string {
	//１回しか繰り返さない想定
	for _, v := range styles {
		for s, _ := range cvList {
			r := regexp.MustCompile(s)
			if r.MatchString(v) {
				if classes == nil {
					classes = append(classes, cvList[s])
				} else {
					classes = append(classes, " ", cvList[s])
				}
			}
		}
	}
	return classes
}
