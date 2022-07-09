package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `json:"status"`
	Data    gin.H  `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type Builder struct {
	btype   ResponseType
	status  int
	data    gin.H
	message string
}

func NewResponse() *Builder {
	return &Builder{
		btype:   TypeSuccess,
		status:  http.StatusOK,
		data:    gin.H{},
		message: "",
	}
}

func (b *Builder) SetType(t ResponseType) *Builder {
	b.btype = t
	return b
}

func (b *Builder) SetStatus(s int) *Builder {
	b.status = s
	return b
}

func (b *Builder) SetData(d gin.H) *Builder {
	b.data = d
	return b
}

func (b *Builder) SetMessage(m string) *Builder {
	b.message = m
	return b
}

func (b *Builder) Build() *Response {
	return &Response{
		Status:  b.btype.String(),
		Data:    b.data,
		Message: b.message,
	}
}

func (b *Builder) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
