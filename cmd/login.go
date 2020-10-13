package cmd

import (
	"github.com/hazelcast/hazelcast-cloud-cli/internal"
	"github.com/hazelcast/hazelcast-cloud-cli/util"
	"github.com/spf13/cobra"
)

// Flags
var (
	apiKey    string
	apiSecret string
)


// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"login"},
	Short:   "This command logins you to Hazelcast Cloud with api-key and api-secret.",
	Run: func(cmd *cobra.Command, args []string) {
		configService := internal.NewConfigService()
		configService.Set(internal.ApiKey, apiKey)
		configService.Set(internal.ApiSecret, apiSecret)
	},
}

func init() {
	if util.IsCloudShell() {
		return
	}
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVar(&apiKey, "api-key", "", "api key of your accunt")
	loginCmd.Flags().StringVar(&apiKey, "api-secret", "", "api secret of your account")
	loginCmd.MarkFlagRequired("api-key")
	loginCmd.MarkFlagRequired("api-secret")
}
