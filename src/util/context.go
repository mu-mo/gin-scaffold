package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

// ErrorRes ErrorResponse
type ErrorRes struct {
	Status int    `json:"status"`
	ErrMsg string `json:"err_msg"`
}

// DataRes DataResponse
type DataRes struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// BindGetJSONData bind the json data of method GET
// body must be a point
func BindGetJSONData(url string, param req.Param, body interface{}) error {
	r, err := req.Get(url, param)
	if err != nil {
		return err
	}
	err = r.ToJSON(body)
	if err != nil {
		return err
	}
	return nil
}

// RetError response error, wrong response
func RetError(c *gin.Context, code, status int, errMsg string) {
	c.JSON(code, ErrorRes{
		Status: status,
		ErrMsg: errMsg,
	})
}

// RetData response data, correct response
func RetData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, DataRes{
		Status: 200,
		Data:   data,
	})
}
