package jwt

import (
	"context"
	"gohub/pkg/firebase"
	"gohub/pkg/logger"
)

type FirebaseProvider struct {
}

func NewFirebaseProvider() *FirebaseProvider {
	return &FirebaseProvider{}
}

func (ft *FirebaseProvider) IssueToken(uid string) string {
	token, err := firebase.GetAuth().CustomToken(context.Background(), uid)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

func (ft *FirebaseProvider) ParseToken(tokenString string) (uid string, err error) {
	token, err := firebase.GetAuth().VerifyIDTokenAndCheckRevoked(context.Background(), tokenString)
	if err != nil {
		return "", err
	}
	return token.UID, nil
}
