package response

type ResponseType string

const (
	TypeSuccess ResponseType = "success"
	TypeFail    ResponseType = "fail"
	TypeError   ResponseType = "error"
)

func (t ResponseType) String() string {
	return string(t)
}
