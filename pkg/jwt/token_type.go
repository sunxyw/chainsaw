package jwt

type TokenType uint

const (
	TokenTypeUser TokenType = iota
	TokenTypeService
)

func (t TokenType) String() string {
	switch t {
	case TokenTypeUser:
		return "user"
	case TokenTypeService:
		return "service"
	}
	return ""
}

func (t TokenType) Prefix() string {
	switch t {
	case TokenTypeUser:
		return "usr_"
	case TokenTypeService:
		return "svc_"
	}
	return ""
}
