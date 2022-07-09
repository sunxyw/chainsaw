package config

import "gohub/pkg/config"

func init() {
	config.Add("firebase", func() map[string]interface{} {
		return map[string]interface{}{

			"key_path": config.Env("FIREBASE_KEY_PATH", "serviceAccountKey.json"),
		}
	})
}
