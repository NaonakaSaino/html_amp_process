package lib

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type HtmlToken struct {
	Content          string
	TokenType        html.TokenType
	HtmlTag          string
	StyleOrScriptFlg int
}

//NGpath:PapillonPath
var replacePath = map[string]string{
	"": "",
}

func GetHtmlDividedByTagAmp(str string) []HtmlToken {

	tr := html.NewTokenizer(strings.NewReader(str))
	tt := tr.Next()
	loop := true
	HtmlTokens := make([]HtmlToken, 0)
	for loop {
		if tt == html.ErrorToken {
			return HtmlTokens
		}
		t := tr.Token()
		ht := t.Data
		var content string
		flg := 0

		switch tt {
		case html.TextToken:
			content = MakeTextTag(t.Data)
		case html.StartTagToken:
			content, flg = MakeStartorSelfClosingTag(t.Data, t.Attr, 0)
		case html.EndTagToken:
			content, flg = MakeEndTag(t.Data)
		case html.CommentToken:
			content = MakeCommentTag(t.Data)
		case html.DoctypeToken:
			content = MakeDoctype(t.Data)
		case html.SelfClosingTagToken:
			content, flg = MakeStartorSelfClosingTag(t.Data, t.Attr, 1)
		}
		HtmlTokens = append(HtmlTokens, HtmlToken{
			Content:          content,
			TokenType:        tt,
			HtmlTag:          ht,
			StyleOrScriptFlg: flg,
		})
		tt = tr.Next()
	}
	return nil
}

//Make〇〇Tag系メソッドは、"<img src="https://~">"のような文字列を返す。
func MakeStartorSelfClosingTag(content string, attrs []html.Attribute, tagType int) (string, int) {
	var tags []string
	tags = append(tags, "<", content, " ")
	if content == "style" || content == "script" {
		return "", 1
		//attrがnilでない場合の処理
		//attrをループさせて、style,class属性の値を保持する。
	} else if attrs != nil {
		//styleとclassの中身を入れるsliceを用意する。
		//それぞれ要素は１つしかない想定だが、一応スライスにする
		var styles, classes []string
		others := map[string]string{}
		for _, atr := range attrs {
			k := atr.Key
			v := atr.Val
			if k == "style" {
				styles = append(styles, v)
			} else if k == "class" {
				classes = append(classes, v)
			} else {
				others[k] = v
			}
			//とりあえずこのループで作っておく
			tags = append(tags, k, "=\"", v, "\"", " ")
		}
		//img,style, script,テキスト部以外かつインラインスタイルありのタグの処理。
		if styles != nil {
			//　style属性がある場合、一度tagsをリセットする。
			//stylesとclassesを渡してclassesを返す
			classes = ConvertStyleToClass(styles, classes)
			tags = nil
			tags = append(tags, "<", content, " class=\"", strings.Join(classes, ""), "\"", " ")
			for k, v := range others {
				tags = append(tags, k, "=\"", v, "\"", " ")
			}
		}
		if content == "img" {
			tags = nil
			tags = MakeImgTag(strings.Join(classes, ""), others)
		}
	}
	//共通処理
	//最後の要素が"" "のため
	tags = tags[0 : len(tags)-1]
	if tagType == 0 {
		tags = append(tags, ">")
	} else {
		tags = append(tags, "/>")
	}
	tagContent := strings.Join(tags, "")
	return tagContent, 0
}

func MakeImgTag(classStr string, others map[string]string) []string {
	var tags, ngList []string
	//img.lazy要素の変換
	if strings.Contains(classStr, "lazy") {
		classStr = strings.Replace(classStr, "lazy", "", 1)
		others["src"] = others["data-src"]
		delete(others, "data-src")
	}
	if !strings.Contains(others["src"], "admin.moneytimes.jp") {
		if strings.Contains(others["src"], "cdn") {
			ngList = append(ngList, others["src"])
		} else {
			//amp-imgの必須属性を記述する width="1" height="1" layout="responsive"
			if v, exists := others["width"]; !exists || v == "" {
				others["width"] = "1"
			}
			if v, exists := others["height"]; !exists || v == "" {
				others["height"] = "1"
			}
			if v, exists := others["layout"]; !exists || v == "" {
				others["layout"] = "responsive"
			}
		}
	}
	if classStr != "" {
		tags = append(tags, "<img", " class=\"", classStr, "\"", " ")
	} else {
		tags = append(tags, "<img", " ")
	}
	for k, v := range others {
		tags = append(tags, k, "=\"", v, "\"", " ")
	}
	if ngList != nil {
		fmt.Println(fmt.Sprintf("下記画像パスはcdnのパスを直接指定しているため使用できません。\n%s", strings.Join(ngList, "\n")))
	}
	return tags
}

//マークダウンの置換
func MakeTextTag(text string) string {
	regex := regexp.MustCompile(`\!\[(.*?)\]\((.*?)\)`)
	if regex.MatchString(text) {
		for regex.MatchString(text) {
			a := regex.FindStringSubmatch(text)
			text = strings.Replace(text, a[0], fmt.Sprintf("<img alt=\"%s\" src=\"%s\" height=\"1\" width=\"1\" layout=\"responsive\">", a[1], a[2]), 1)
			if strings.Contains(a[2], "cdn") {
				fmt.Println(fmt.Sprintf("下記画像パスはcdnのパスを直接指定しているため使用できません。\n%s", a[2]))
			}
		}
	}
	return text
}

func MakeEndTag(content string) (string, int) {
	var tags []string
	if content == "style" || content == "script" {
		return "", 1
	}
	tags = append(tags, "</", content, ">")
	tagContent := strings.Join(tags, "")
	return tagContent, 0
}

func MakeDoctype(content string) string {
	var tags []string
	tags = append(tags, "<!DOCTYPE ", content, ">")
	tagContent := strings.Join(tags, "")
	return tagContent
}

func MakeCommentTag(content string) string {
	var tags []string
	tags = append(tags, "<!--", content, "-->")
	tagContent := strings.Join(tags, "")
	return tagContent
}

func UnifyHtmlTokens(Tokens []HtmlToken) string {
	var contentList []string
	for i := 0; i < len(Tokens); i++ {
		if Tokens[i].StyleOrScriptFlg == 1 {
			if i+2 < len(Tokens) {
				if Tokens[i+2].StyleOrScriptFlg == 1 {
					i++
					continue
				}
			}
			continue
		}
		contentList = append(contentList, Tokens[i].Content)
	}
	return strings.Join(contentList, "")
}
