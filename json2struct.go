package json2struct

import (
	"regexp"
)

var tabline, _ = regexp.Compile(`\n\t+\n`)

func Json2Struct(name, jsonStr string, genType GeneratorType) (string, error) {
	jt, v := toValue(jsonStr)
	node := NewNode("", v, genType, jt, NewP())
	ret, err := node.JsonToStruct(jsonStr)
	//node.debug()
	if err != nil {
		return ret, err
	}
	ret = tabline.ReplaceAllString(ret, "\n")
	ret = "type " + toProperCase(name) + " " + ret
	return ret, err
}