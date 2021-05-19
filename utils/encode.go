package utils

import (
	iconv "github.com/djimenez/iconv-go"
)

var (
	CVtolocal *iconv.Converter
	CVtoutf8  *iconv.Converter
)

func InitConverter(local string) bool {
	cv, err := iconv.NewConverter("utf-8", local)
	if err != nil {
		panic(err)
	}
	CVtolocal = cv
	cv, err = iconv.NewConverter(local, "utf-8")
	if err != nil {
		panic(err)
	}
	CVtoutf8 = cv
	return true
}

func Utf8ToLocal(str string) (b string, err error) {
	buf, err := CVtolocal.ConvertString(str)
	return string(buf), err
}

func LocalToUtf8(str string) (b string, err error) {
	buf, err := CVtoutf8.ConvertString(str)
	return string(buf), err
}
