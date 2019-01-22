package Library

import (
	"github.com/labstack/echo"
	"net/http"
)

/**
* author wood
* api 返回格式封装
*/

type Response struct {
	Context   echo.Context
	parameter *RetParameter
}

type RetParameter struct {
	Code int
	Data interface{}
	Msg  string
}

const DefaultCode = 1
var HttpStatus = http.StatusOK

// 初始化Response
func NewResponse(c echo.Context) *Response {

	R := new(Response)
	R.Context = c
	R.parameter = new(RetParameter)
	R.parameter.Data = nil
	return R
}

// 设置返回的Status值默认http.StatusOK
func (this *Response)SetStatus(i int) {
	HttpStatus = i
}

func (this *Response)SetMsg(s string) {
	this.parameter.Msg = s
}

func (this *Response)SetData(d interface{}) {
	this.parameter.Data = d
}

// 返回自定自定义的消息格式
func (this *Response)RetCustomize(code int, d interface{}, msg string) error {

	this.parameter.Code = code
	this.parameter.Data = d
	this.parameter.Msg = msg

	return this.Ret(this.parameter)
}

// 返回成功的结果 默认code为1
func (this *Response)RetSuccess(d interface{}) error {

	this.parameter.Code = DefaultCode
	this.parameter.Data = d

	return this.Ret(this.parameter)
}

// 返回失败结果
func (this *Response)RetError(e error, c int) error {

	this.parameter.Code = c
	this.parameter.Msg = e.Error()

	return this.Ret(this.parameter)
}

// 返回结果 - 目前仅支持json格式
func (this *Response)Ret(par interface{}) error {
	return this.Context.JSON(HttpStatus, par)
}

// 输出返回结果
func (this *Response)Write(b []byte) {

	_, e := this.Context.Response().Write(b)
	if e != nil {
		print(e.Error())
	}
}