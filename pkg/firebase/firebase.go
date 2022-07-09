package firebase

import (
	"context"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"path/filepath"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var once sync.Once
var app *firebase.App

func InitFirebase() {
	once.Do(func() {
		serviceAccountKeyPath := config.Get[string]("firebase.key_path")
		serviceAccountKeyFilePath, err := filepath.Abs(serviceAccountKeyPath)
		logger.LogIf(err)

		opt := option.WithCredentialsFile(serviceAccountKeyFilePath) //Firebase admin SDK initialization
		app, err = firebase.NewApp(context.Background(), nil, opt)
		logger.LogIf(err)
	})
}

func GetAuth() *auth.Client {
	auth, err := app.Auth(context.Background())
	logger.LogIf(err)

	return auth
}
