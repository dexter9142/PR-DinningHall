package main

import "time"

func Div(str string) string {
	return "<div>" + str + "</div>"
}

func genTime() int64 {
	return time.Now().UnixNano() / int64(timeVar)
}
