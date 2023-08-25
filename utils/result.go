package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

// R 返回
func R(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

// RespList 返回list类型
func RList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

// RFail 返回失败
func RFail(w http.ResponseWriter, msg string) {
	R(w, -1, nil, msg)
}

// Rok 返回成功
func Rok(w http.ResponseWriter, data interface{}, msg string) {
	R(w, 0, data, msg)
}

// RespOKList 返回list Ok
func RespOKList(w http.ResponseWriter, data interface{}, total interface{}) {
	RList(w, 0, data, total)
}
