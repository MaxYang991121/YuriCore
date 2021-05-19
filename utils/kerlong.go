package utils

import "fmt"

//IntAbs 绝对值
func IntAbs(num int) int {
	ans, ok := Ternary(num > 0, num, -num).(int)
	if ok {
		return ans
	}
	return 0
}

//Ternary 三目运算符
func Ternary(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}

func IsAllNumber(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] > '9' || str[i] < '0' {
			return false
		}
	}
	return true
}

//ScanLine 得到一行
func ScanLine() (line string) {
	var buffer []rune
	for {
		var c rune
		n, err := fmt.Scanf("%c", &c)

		//FIXME: in windows,line feeds are '\r\n',but this may cause some problem in UNIX or MAXOS
		if nil != err ||
			1 != n ||
			//'\r' == c ||
			'\n' == c {
			break
		}
		buffer = append(buffer, c)
	}
	return string(buffer)
}
