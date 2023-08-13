package subcmd

import (
	"fmt"
	platform "github.com/esonhugh/openai-platform-api"
	"github.com/esonhugh/openai-platform-api/cmd/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var BillCmd = &cobra.Command{
	Use:   "usage",
	Short: "find out the usage from arg1 to arg2. if 1 arg, from arg1 to Today. if 0 arg, get last month",
	Run: func(cmd *cobra.Command, args []string) {
		if client.Client == nil {
			log.Errorln("No available client")
			os.Exit(-1)
		}
		var resp platform.UsageResponse
		var err error
		var from, to string
		if len(args) == 2 {
			from = args[0]
			to = args[1]
		} else if len(args) == 1 {
			now := time.Now()
			from = args[0]
			to = fmt.Sprintf("%v-%02d-%v", now.Year(), now.Month(), now.Day())
		} else {
			from, to = platform.GetLastMonth()
		}
		resp, err = client.Client.UsageWithSessionToken(from, to)
		if err != nil {
			log.Errorln("Get Usage hits error", err)
			log.Debug(client.Client.LastResponse())
			return
		}
		log.Infof("From %v To %v", from, to)
		log.Infof("Total count: $%v", resp.TotalUsage/100)
	},
}
