package jwt

type TokenType uint

const (
	TokenTypeUser TokenType = iota
	TokenTypeService

	TokenTypeAll = TokenTypeUser | TokenTypeService
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

func getTokenTypeList(t TokenType) []TokenType {
	var list []TokenType
	if t&TokenTypeAll == TokenTypeAll {
		list = []TokenType{TokenTypeUser, TokenTypeService}
	} else {
		list = []TokenType{t}
	}
	return list
}
