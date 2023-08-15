## Simple OpenAI Platform API 

Support list keys add key and delete key with it

checkout test/demo.go file about it

Idea comes from https://github.com/linweiyuan/go-chatgpt-api project. Thanks and respect.

### Lib Usage

```go

package main

import (
	"log"
	"github.com/esonhugh/openai-platform-api"
)

func main() {
	c := platform.NewUserPlatformClient("accessToken there")
	_, err := c.LoginWithAccessToken()
	if err != nil {
		log.Println(err)
		return
	}
	// if you want accessToken based without login by username and password please use this with
	//  c.LoginWithAccessToken()
	// or replace it with
	//  c.LoginWithAuth0("username", "password")
	resp, err := c.GetSecretKeys()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
```

### Cli tool guides

#### Installation

```bash
go install github.com/esonhugh/openai-platform-api/cmd/openai-cli@latest
```

#### Before use

```bash
# you need create a config file in ~/.config folder
cat << EOF > ~/.config/openai-cli/config.yaml
openai:
  username: your-openai-username
  password: your-openai-password
secret_key_prefix: temp
EOF
```

#### Basic Usage

```bash
export https_proxy=http://127.0.0.1:7890 # optional
openai-cli key ls
```

