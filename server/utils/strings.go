package utils

import (
	"fmt"
	"strconv"
	"strings"
)

/*
1、func Title(s string) string
字符串s每个单词首字母大写返回
2、func ToLower(s string) string
所有小写
3、func ToLowerSpecial(_case unicode.SpecialCase, s string) string
将字符串s中所有字符按_case指定的映射转换成小写返回
4、func ToTitle(s string) string
将字符串s转换成大写返回
5、func ToTitleSpecial(_case unicode.SpecialCase, s string) string
将字符串s中所有字符按_case指定的映射转换成大写返回
6、func ToUpper(s string) string
所有小写
7、func ToUpperSpecial(_case unicode.SpecialCase, s string) string
将字符串s中所有字符按_case指定的映射转换成大写返回
*/

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}


func MaheHump(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

func CheckSpace(r rune) bool {
	if r == ' ' {
		return true
	}
	return false
}
func FindDataType(s string) string {
	if len(s) == 0 {
		return ""
	}
	s = strings.Replace(s, "（", "(", -1)
	s = strings.Replace(s, "）", ")", -1)
	for i := 0; i < len(s)-1; i++ {
		if s[i:i+1] == "(" {
			return s[:i]
		}
	}
	return s
}

func FindDataLen(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("字符串长度为0")
	}
	s = strings.Replace(s, "（", "(", -1)
	s = strings.Replace(s, "）", ")", -1)
	var l string
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" {
			l = s[i+1:]
			break
		}
	}
	println("l:", l)
	if len(l) == 0 {
		return 0, fmt.Errorf("字符串无长度'('")
	}

	var d string
	for i := 0; i < len(l); i++ {
		if l[i:i+1] == ")" {
			d = l[:i]
			break
		}
	}
	println("d:", d)
	if len(d) == 0 {
		return 0, fmt.Errorf("字符串无长度')'")
	}
	i, err := strconv.Atoi(d)
	if err != nil {
		return 0, fmt.Errorf("字符串转数字失败:%s", d)
	}
	println("i", i)
	return i, nil
}


func SubBefore(s, sep, def string) string {
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return s[:i]
		}
	}
	return def
}


func SubBeforeLast(s, sep, def string) string {
	for i := len(s) - len(sep); i > -1; i-- {
		if s[i:i+len(sep)] == sep {
			return s[:i]
		}
	}
	return def
}

func SubAfter(s, sep, def string) string {
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return s[i+len(sep):]
		}
	}
	return def
}

func SubAfterLast(s, sep, def string) string {
	for i := len(s) - len(sep); i > -1; i-- {
		if s[i:i+len(sep)] == sep {
			return s[i+len(sep):]
		}
	}
	return def
}

func TrimStart(s, trim string) string {
	if strings.HasPrefix(s, trim) {
		return s[len(trim):]
	}
	return s
}

func TrimEnd(s, trim string) string {
	if strings.HasSuffix(s, trim) {
		return s[:len(s)-len(trim)]
	}
	return s
}

func TrimBoth(s, trim string) string {
	return TrimStart(TrimEnd(s, trim), trim)
}


func StringIntersect(string1, string2 []string) []string {
	m := make(map[string]int)
	r := make([]string, 0)
	for _, v1 := range string1 {
		m[v1]++
	}
	for _, v2 := range string2 {
		_, ok := m[v2]
		if ok {
			r = append(r, v2)
		}
	}
	return r
}


func StringDifference(string1, string2 []string) []string {
	m := make(map[string]int)
	r := make([]string, 0)
	inter := StringIntersect(string1, string2)
	for _, v1 := range inter {
		m[v1]++
	}
	for _, v2 := range string1 {
		_, ok := m[v2]
		if !ok {
			r = append(r, v2)
		}
	}
	return r
}


func IntDif(int1, int2 []int) []int {
	var r []int
	for _, e := range int1 {
		i := 0
		for _, n := range int2 {
			if e != n {
				i++
			}
		}
		if i == len(int2) {
			r = append(r, e)
		}
	}
	return r
}


func stringIntersect(string1, string2 []string) []string {
	m := make(map[string]int)

	r := make([]string, 0)
	for _, v1 := range string1 {
		m[v1]++

	}

	for _, v2 := range string2 {
		_, ok := m[v2]
		if ok {
			r = append(r, v2)
		}
	}
	return r
}
