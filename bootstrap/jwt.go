package bootstrap

import (
	"errors"
	"gohub/pkg/config"
	"gohub/pkg/jwt"
)

// SetupJWT 初始化 JWT
func SetupJWT() {
	selectedProvider := config.Get[string]("jwt.token_provider")

	switch selectedProvider {
	case "local":
		jwt.InitWithProvider(jwt.NewLocalProvider())
	case "firebase":
		jwt.InitWithProvider(jwt.NewFirebaseProvider())
	default:
		panic(errors.New("unknown jwt token provider"))
	}
}
