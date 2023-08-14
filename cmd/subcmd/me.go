package subcmd

import (
	"github.com/esonhugh/openai-platform-api/cmd/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var MeCmd = &cobra.Command{
	Use:     "me",
	Aliases: []string{"user", "u"},
	Short:   "Operating personal data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var f = log.Infoln

var MeSubCmd = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"subscription", "pay", "payment"},
	Short:   "Show user subscription data",
	Run: func(cmd *cobra.Command, args []string) {
		if client.Client == nil {
			log.Errorln("No available client")
			os.Exit(-1)
		}

		f("===== User Subscription Data =====")
		resp3, err := client.Client.DashboardSubscription()
		if err != nil {
			log.Errorln("payment subscription hits error", err)
			log.Debug(client.Client.LastResponse())
			os.Exit(-3)
		}
		f("Account Name:", resp3.AccountName)
		f("Soft Limit: $", resp3.SoftLimitUsd)
		f("Hard Limit: $", resp3.HardLimitUsd)
		f("Bill Address Country:", resp3.BillingAddress.Country)
		f("Bill Address State:", resp3.BillingAddress.State)
		f("Bill Address City:", resp3.BillingAddress.City)
		f("Bill Address Ln.1:", resp3.BillingAddress.Line1)
		f("Bill Address Ln.2:", resp3.BillingAddress.Line2)
		f("Post Code:", resp3.BillingAddress.PostalCode)

		f("===== User Payment Information =====")
		if resp3.HasPaymentMethod {
			resp2, err := client.Client.PaymentMethod()
			if err != nil {
				log.Errorln("payment data hits error", err)
				log.Debug(client.Client.LastResponse())
				os.Exit(-3)
			}
			for i, data := range resp2.Data {
				f("Payment #", i)
				if data.IsDefault == true {
					f("* This is Default Payment")
				}
				f("Payment Type:", data.Type)
				f("Card Type:", data.Card.Brand)
				f("Card Number: ******", data.Card.Last4)
				f("Card Exp:", data.Card.ExpYear, "-", data.Card.ExpMonth)
				f("Card Country:", data.Card.Country)
			}
		} else {
			f("User has no payment method.")
		}
	},
}

var MeInfoCmd = &cobra.Command{
	Use:   "info",
	Short: `display your information`,
	Run: func(cmd *cobra.Command, args []string) {
		if client.Client == nil {
			log.Errorln("No available client")
			os.Exit(-1)
		}

		f("===== User Basic Information =====")
		resp, err := client.Client.DashboardOnBoarding()
		if err != nil {
			log.Errorln("Dashboard On boarding hits error", err)
			log.Debug(client.Client.LastResponse())
			os.Exit(-2)
		}
		f("Username:", resp.User.Name)
		f("UserID:", resp.User.ID)
		f("Email:", resp.User.Email)
		f("SessionToken:", resp.User.Session.SensitiveID)
	},
}
