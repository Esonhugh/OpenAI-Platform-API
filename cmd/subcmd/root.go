package subcmd

import (
	platform "github.com/esonhugh/openai-platform-api"
	"github.com/esonhugh/openai-platform-api/cmd/client"
	"github.com/esonhugh/openai-platform-api/cmd/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(BillCmd, KeyCmd)
	KeyCmd.AddCommand(KeyListCmd, KeyTempCmd, KeyDelCmd, KeyAddCmd)
}

var RootCmd = &cobra.Command{
	Use:   "openai-cli",
	Short: "Openai platform API cli tools",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.DebugLevel)
		config.Init()
		if config.C.Config.AccessToken != "" {
			client.Client = platform.NewUserPlatformClient(config.C.Config.AccessToken)
			_, err := client.Client.DashboardOnBoarding()
			if err == nil {
				log.Infoln("Connect to platform successful")
				return
			}
			log.Warnln("AccessToken Expired or bad one")
		}
		client.Client = platform.NewUserPlatformClient("")
		log.Infoln("Try login as username and password....")
		username := config.C.Config.Openai.Username
		password := config.C.Config.Openai.Password
		if username == "" || password == "" {
			log.Errorln("Empty username or password")
			client.Client = nil
			return
		}
		err := client.Client.LoginWithAuth0(username, password)
		if err == nil {
			config.C.Set("access_token", client.Client.AccessToken())
			config.C.WriteConfig()
			return
		}
		log.Errorln("Login fail")
		client.Client = nil
		return
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
