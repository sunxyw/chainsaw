package jwt

type TokenProvider interface {
	IssueToken(uid string) string
	ParseToken(token string) (uid string, err error)
}
