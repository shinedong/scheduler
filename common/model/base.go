package model

import (
	"encoding/json"
	"time"
)

const (
	DefaultSuccessCode = "0"
	DefaultSuccessMag  = "success"
)

type BaseModel struct {
	Id           string    `json:"id"`
	Creator      string    `json:"creator"`
	Modifier     string    `json:"modifier"`
	CreateTime   time.Time `json:"create_time"`
	ModifiedTime time.Time `json:"modified_time"`
}

type BaseResult struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildResult(code string, message string, data interface{}) (resp []byte, err error) {
	var baseResult BaseResult
	baseResult.Code = code
	baseResult.Message = message
	baseResult.Data = data
	resp, err = json.Marshal(baseResult)
	return
}

func BuildSuccessResult(data interface{}) (resp []byte, err error) {
	return BuildResult(DefaultSuccessCode, DefaultSuccessMag, data)
}

func (result *BaseResult) isSuccess() (flag bool) {
	return result.Code == DefaultSuccessCode
}

func main() {
	var result *BaseResult = &BaseResult{}
	if result.isSuccess() {

	}
}
