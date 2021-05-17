package main

import (
	"embed"

	// _ "embed"
	"html/template"
	"os"
	// "fmt"
	// lgpd "leetcode/code/longestpalindromic"
)

var (
	data = "abbcccbbbcaaccbababcbcabca"
)

//go:embed fixtures/*
var s string

//go:embed templates/*
var tpls embed.FS

func main() {
	//fmt.Println("data result of abbcccbbbcaaccbababcbcabca", lgpd.LongestPalindrome(data))

	print(s)
	name := "en.tmpl"
	if len(os.Args) > 1 {
		name = os.Args[1] + ".tmpl"
	}
	arg := "World"
	if len(os.Args) > 2 {
		arg = os.Args[2]
	}
	t, err := template.ParseFS(tpls, "templates/*")
	if err != nil {
		panic(err)
	}
	if err = t.ExecuteTemplate(os.Stdout, name, arg); err != nil {
		panic(err)
	}
}
