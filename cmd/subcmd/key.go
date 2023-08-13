package subcmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/esonhugh/openai-platform-api/cmd/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"time"
)

var KeyCmd = &cobra.Command{
	Use:   "key",
	Short: "Key sub command for key management",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var KeyDelCmd = &cobra.Command{
	Use:     "del",
	Aliases: []string{"rm"},
	Short:   "delete the key",
	Run: func(cmd *cobra.Command, args []string) {
		if client.Client == nil {
			log.Errorln("No available client")
			os.Exit(-1)
		}
		resp, err := client.Client.GetSecretKeys()
		if err != nil {
			log.Errorln("Get Secret Keys hits error", err)
			log.Debug(client.Client.LastResponse())
			os.Exit(-2)
		}

		var option []string
		for _, key := range resp.Data {
			log.Infof("Detect %v, Creatd at %v, Lasted Used at %v, name: \"%v\" value: %v",
				key.Object,
				time.Unix(int64(key.Created), 0).Format("2006-01-02"),
				time.Unix(int64(key.LastUse), 0).Format("2006-01-02"),
				key.Name,
				key.SensitiveID)
			option = append(option, fmt.Sprintf("%v - %v", key.Name, key.SensitiveID))
		}

		var selectedKey string
		err = survey.AskOne(&survey.Select{
			Message: "Choose the key you need delete",
			Options: option,
		}, &selectedKey)
		if err != nil {
			log.Errorln("Exit without asking", err)
			os.Exit(-3)
		}

		for _, key := range resp.Data {
			if selectedKey == fmt.Sprintf("%v - %v", key.Name, key.SensitiveID) {
				respDelete, err := client.Client.DeleteSecretKey(key)
				if err != nil {
					log.Errorln("Delete Key Errored", err)
					log.Debug(client.Client.LastResponse())
					os.Exit(-3)
				}
				log.Infoln("Delete Key ", respDelete.Result)
				os.Exit(0)
			}
		}
		log.Warnln("The key is not found.")
	},
}

var KeyListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "get"},
	Short:   "list the keys currently",
	Run: func(cmd *cobra.Command, args []string) {
		if client.Client == nil {
			log.Errorln("No available client")
			os.Exit(-1)
		}
		resp, err := client.Client.GetSecretKeys()
		if err != nil {
			log.Errorln("Get Secret Keys hits error", err)
			log.Debug(client.Client.LastResponse())
			os.Exit(-2)
		}
		for _, key := range resp.Data {
			log.Infof("Detect %v, Creatd at %v, Lasted Used at %v, name: \"%v\" value: %v",
				key.Object,
				time.Unix(int64(key.Created), 0).Format("2006-01-02"),
				time.Unix(int64(key.LastUse), 0).Format("2006-01-02"),
				key.Name,
				key.SensitiveID)
		}
	},
}

var KeyTempCmd = &cobra.Command{
	Use:     "temp",
	Aliases: []string{"t"},
	Short:   "create a temp key, delete with graceful shutdown",
	Run: func(cmd *cobra.Command, args []string) {
		if client.Client == nil {
			log.Errorln("No available client")
			os.Exit(-1)
		}
		name := "temp"
		if len(args) > 1 {
			name = args[0]
		}
		resp, err := client.Client.CreateSecretKey(name)
		if err != nil {
			log.Errorln("Get Secret Keys hits error", err)
			log.Debug(client.Client.LastResponse())
			os.Exit(-2)
		}
		log.Infoln("Create Key ", resp.Result)
		log.Infof("%v: %v", resp.Key.Object, resp.Key.SensitiveID)

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Kill)
		<-quit
		// quit
		respDelete, err := client.Client.DeleteSecretKey(resp.Key)
		if err != nil {
			log.Errorln("Delete Key Errored", err)
			log.Debug(client.Client.LastResponse())
			os.Exit(-3)
		}
		log.Infoln("Delete Key ", respDelete.Result)
	},
}
