package main

// global object
var (
	router  Router
	setting Setting
)

// global marco
const (
	RESULT_FAILED = 0
	RESULT_OK     = 1

	ADD_OP     = 0
	SUBTRAC_OP = 1
)

// error message
const (
	SUCCESS_MSG                   = "成功"
	ERROR_MSG_INVALID_METHOD      = "错误的请求方法"
	ERROR_MSG_SERVER_ERROR        = "服务器内部错误"
	ERROR_MSG_PARSE_PARAMS_FAILED = "解析参数失败"
	ERROR_MSG_INVALID_PARAMS      = "无效的参数"
)
