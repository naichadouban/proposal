package main

import (
	"bytes"
	"fmt"
	"strings"
)
//语义字符
const semanticAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
// 这些常量定义了应用的版本并且遵从下面语义
// versioning 2.0.0 spec (http://semver.org/).
const (
	appMajor uint = 0
	appMinor uint = 1
	appPatch uint = 0

	// appPreRelease MUST only contain characters from semanticAlphabet
	// appPreRelease 只能包含语义字符semanticAlphabet中的字母
	// per the semantic versioning spec.

	appPreRelease = ""
)
func version(){
	version := fmt.Sprintf("%d.%d.%d",appMajor,appMinor,appPatch)
}

// normalizeVerString
// 返回通过的字符，
func normalizeVerString(str string) string{
	var result bytes.Buffer
	for _,r := range str {
		if strings.ContainsRune(semanticAlphabet,r){  // 这里说明一个字符是4字节吗
			result.WriteRune(r)
		}
	}
	return result.String()
}