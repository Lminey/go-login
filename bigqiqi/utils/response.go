package utils

import (
	"reflect"
)

type DataResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (this *DataResponse) SetError(errCode int, errMsg string) {
	this.Msg = errMsg
	this.Status = errCode
}

func (this *DataResponse) SetData(data interface{}) {
	this.Data = data
}

func Response(errCode int, msg string, data interface{}) DataResponse {
	responseData := make([]map[string]interface{}, 0)
	if data == nil {
		data = responseData
	} else {
		if reflect.TypeOf(data).String() == "string" {
			if reflect.ValueOf(data).Len() == 0 {
				data = responseData
			}
		}
		if reflect.TypeOf(data).String() == "[]interface{}" {
			if reflect.ValueOf(data).Len() == 0 {
				data = responseData
			}
		}
		if reflect.TypeOf(data).String() == "[]orm.Params" {
			if reflect.ValueOf(data).Len() == 0 {
				data = responseData
			}
		}
	}

	resp := DataResponse{
		Status: errCode,
		Msg:    msg,
		Data:   data,
	}

	return resp
}

func FormatRes(rsp *DataResponse) map[string]interface{} {
	res := make(map[string]interface{})
	res["msg"] = rsp.Msg
	res["status"] = rsp.Status
	res["data"] = rsp.Data

	return res
}

func NewData() map[string]interface{} {
	return make(map[string]interface{})
}
