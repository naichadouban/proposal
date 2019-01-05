package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	str1 := "helloworld=xuxiaofeng"
	index := strings.Index(str1, "h")
	fmt.Println(index)
	if index == 1 {
		t.Fatalf("want %v,get %v", 1, index)
	}
}
