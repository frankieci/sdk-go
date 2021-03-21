package builder

import (
	"net/http"

	"gitlab.com/sdk-go/library/validator"

	"github.com/gin-gonic/gin"
	v "github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type ResultBuilder struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
}

func BuildSuccess(c *gin.Context, httpCode int) {
	c.JSON(httpCode, ResultBuilder{
		Success: true,
	})
}

func BuildSuccessWithData(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, ResultBuilder{
		Success: true,
		Data:    data,
	})
}

func BuildError(c *gin.Context, httpCode int, errorMessage string) {
	c.JSON(httpCode, ResultBuilder{
		Success:      false,
		ErrorMessage: errorMessage,
	})
}

func BuildBindError(c *gin.Context, err error) {
	if err == nil {
		panic(errors.New("必须向BuildBindError方法传递非空的err"))
	}
	//获取validator.ValidationErrors类型的err
	errs, ok := err.(v.ValidationErrors)
	if !ok {
		//非validator.ValidationErrors错误直接返回
		c.JSON(http.StatusBadRequest, ResultBuilder{
			ErrorMessage: "数据处理异常",
			Success:      false,
		})
	} else {
		//validator.ValidationErrors类型进行翻译
		c.JSON(http.StatusBadRequest, ResultBuilder{
			ErrorMessage: validator.GetErrMsg(errs.Translate(validator.Trans)),
			Success:      false,
		})
	}
}
