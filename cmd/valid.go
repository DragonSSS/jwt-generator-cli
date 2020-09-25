package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	valCmd = &cobra.Command{
		Use:   "val",
		Short: "validate jwt toke with provided public key pem",
		Long:  "validate jwt toke with provided public key pem",
		Run:   validate,
	}
)

func validateToken(pemPath, token string) {

}

func validate(cmd *cobra.Command, args []string) {
	pemPath := cmd.Flag("publicPem").Value.String()
	token := cmd.Flag("token").Value.String()

	validateToken(pemPath, token)
}

func init() {
	rootCmd.AddCommand(valCmd)

	valCmd.Flags().StringP("publicPem", "p", "", "path of public key pem file")
	err := valCmd.MarkFlagRequired("publicPem")
	if err != nil {
		log.WithError(err).WithField("parameter", "publicPem").Fatal("Failed to parse parameter")
	}

	valCmd.Flags().StringP("token", "t", "", "string of jwt token")
	err = valCmd.MarkFlagRequired("token")
	if err != nil {
		log.WithError(err).WithField("parameter", "token").Fatal("Failed to parse parameter")
	}
}
