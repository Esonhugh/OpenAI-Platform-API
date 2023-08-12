package main

import (
	platform "github.com/esonhugh/openai-platform-api"
	"log"
)

func main() {
	c := platform.NewUserPlatformClient("")
	// if you want accessToken based without login please use this with
	c.CheckStatus()

	// instead of
	c.LoginAuth("username", "password")

	_, err := c.GetSecretKeys()
	if err != nil {
		log.Println(err)
	}
	resp, err := c.CreateSecretKey("CreateKeyName_114514")
	if err != nil {
		log.Println(err)
	}

	log.Println(resp.Key.Object, resp.Key.SensitiveID)

	resp2, err := c.GetSecretKeys()
	if err != nil {
		log.Println(err)
	}
	var key_delete platform.Key
	for _, key := range resp2.Data {
		if key.Name == "CreateKeyName_114514" {
			key_delete = key
		}
	}
	c.DeleteSecretKey(key_delete)

}
