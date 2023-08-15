package subcmd

import (
	"fmt"
	platform "github.com/esonhugh/openai-platform-api"
	"github.com/esonhugh/openai-platform-api/cmd/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math"
	"os"
	"time"
)

var ui bool

func init() {
	BillCmd.Flags().BoolVarP(&ui, "ui", "u", false, "output with UI with every day data")
}

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
			now := time.Now().AddDate(0, 0, 1)
			from = args[0]
			to = fmt.Sprintf("%v-%02d-%v", now.Year(), now.Month(), now.Day())
		} else {
			from, to = platform.GetCurrentMonth()
		}
		resp, err = client.Client.UsageWithSessionToken(from, to)
		if err != nil {
			log.Errorln("Get Usage hits error", err)
			log.Debug(client.Client.LastResponse())
			return
		}
		log.Infof("From %v To %v", from, to)
		log.Infof("Total count: $%v", resp.TotalUsage/100)
		if ui {
			var days []struct {
				Cost float64
				Time float64
			}
			bigestCost := 0.0
			for _, dc := range resp.DailyCosts {
				fullCost := 0.0
				for _, li := range dc.LineItems {
					fullCost += li.Cost
				}
				bigestCost = math.Max(bigestCost, fullCost)
				days = append(days, struct {
					Cost float64
					Time float64
				}{
					Cost: fullCost,
					Time: dc.Timestamp,
				})
			}
			eachP := bigestCost / 32

			fmt.Println("------------------------------------------------------")
			fmt.Println("|    DATE    | COST |           Status               |")
			fmt.Println("------------------------------------------------------")
			for _, day := range days {
				if day.Time == 0 {
					continue
				}
				star := ""
				for i := 0; i < 32; i++ {
					if i < int(day.Cost/eachP) {
						star += "*"
					} else {
						star += " "
					}
				}
				fmt.Printf("| %v | %.2f |%v|\n", // 1 + 10 + 1 + 4 + 1 + 32 + 1
					time.Unix(int64(day.Time), 0).Format("2006-01-02"),
					day.Cost/100,
					star,
				)
			}
			fmt.Println("------------------------------------------------------")
			fmt.Printf("       TOTAL: %.2f\n", resp.TotalUsage/100)
			fmt.Println("------------------------------------------------------")
		}
	},
}
