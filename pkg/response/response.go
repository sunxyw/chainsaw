// Package response 响应处理工具
package response

import (
	"gohub/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Custom 响应自定义响应
func Custom(c *gin.Context, response *Builder) {
	c.JSON(response.status, response)
}

// CustomAbort 响应自定义响应，但会中断后续表演
func CustomAbort(c *gin.Context, response *Builder) {
	c.Abort()
	Custom(c, response)
}

// Success 响应 200 和预设『操作成功！』的 JSON 数据
// 执行某个『没有具体返回数据』的『变更』操作成功后调用，例如删除、修改密码、修改手机号
func Success(c *gin.Context) {
	response := NewResponse().SetMessage("操作成功！")
	Custom(c, response)
}

// SuccessWithData 响应 200 和带 data 键的 JSON 数据
func SuccessWithData(c *gin.Context, data gin.H, msg ...string) {
	response := NewResponse().SetData(data)
	if len(msg) > 0 {
		response.SetMessage(msg[0])
	}
	Custom(c, response)
}

// Created 响应 201 和带 data 键的 JSON 数据
// 执行『更新操作』成功后调用，例如更新话题，成功后返回已更新的话题
func Created(c *gin.Context, data gin.H) {
	response := NewResponse().SetData(data).SetStatus(http.StatusCreated)
	Custom(c, response)
}

// NotFound 响应 404，未传参 msg 时使用默认消息
func NotFound(c *gin.Context, msg ...string) {
	response := NewResponse().
		SetType(TypeError).
		SetStatus(http.StatusNotFound).
		SetMessage(defaultMessage("请求资源不存在，请确认请求正确", msg...))
	CustomAbort(c, response)
}

// Forbidden 响应 403，未传参 msg 时使用默认消息
func Forbidden(c *gin.Context, msg ...string) {
	response := NewResponse().
		SetType(TypeError).
		SetStatus(http.StatusForbidden).
		SetMessage(defaultMessage("您没有权限进行此操作", msg...))
	CustomAbort(c, response)
}

// ServerError 响应 500，未传参 msg 时使用默认消息
func ServerError(c *gin.Context, msg ...string) {
	response := NewResponse().
		SetType(TypeError).
		SetStatus(http.StatusInternalServerError).
		SetMessage(defaultMessage("服务器内部错误，请稍后再试", msg...))
	CustomAbort(c, response)
}

// BadRequest 响应 400，传参 err 对象，未传参 msg 时使用默认消息
// 在解析用户请求，请求的格式或者方法不符合预期时调用
func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	response := NewResponse().
		SetType(TypeError).
		SetStatus(http.StatusBadRequest).
		SetMessage(defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...))
	CustomAbort(c, response)
}

// ValidationError 处理表单验证不通过的错误
func ValidationError(c *gin.Context, errors map[string][]string) {
	response := NewResponse().
		SetType(TypeError).
		SetStatus(http.StatusUnprocessableEntity).
		SetMessage("请求验证不通过，具体请查看 data").
		SetData(gin.H{
			"errors": errors,
		})
	CustomAbort(c, response)
}

// Unauthorized 响应 401，未传参 msg 时使用默认消息
// 登录失败、jwt 解析失败时调用
func Unauthorized(c *gin.Context, msg ...string) {
	response := NewResponse().
		SetType(TypeError).
		SetStatus(http.StatusUnauthorized).
		SetMessage(defaultMessage("认证不通过，请确定已登录并携带令牌", msg...))
	CustomAbort(c, response)
}

// Unavailable 响应 503，未传参 msg 时使用默认消息
func Unavailable(c *gin.Context, msg ...string) {
	response := NewResponse().
		SetType(TypeError).
		SetStatus(http.StatusServiceUnavailable).
		SetMessage(defaultMessage("服务器繁忙，请稍后再试", msg...))
	CustomAbort(c, response)
}

func Fail(c *gin.Context, msg ...string) {
	response := NewResponse().
		SetType(TypeFail).
		SetStatus(http.StatusUnprocessableEntity).
		SetMessage(defaultMessage("操作失败！", msg...))
	Custom(c, response)
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
