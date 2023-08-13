package main

import (
	platform "github.com/esonhugh/openai-platform-api"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	c := platform.NewUserPlatformClient("")

	// if you want accessToken based without login please use this with
	// c.LoginWithAccessToken()
	// instead of
	// c.LoginWithAuth0("username", "password")

	resp, err := c.GetSecretKeys()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)

	resp1, err := c.CreateSecretKey("CreateKeyName_114514")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(resp1.Key.Object, resp1.Key.SensitiveID)

	resp2, err := c.GetSecretKeys()
	if err != nil {
		log.Println(err)
		return
	}
	var key_delete platform.Key
	for _, key := range resp2.Data {
		if key.Name == "CreateKeyName_114514" {
			key_delete = key
		}
	}
	c.DeleteSecretKey(key_delete)

}
