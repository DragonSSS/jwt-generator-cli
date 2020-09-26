package cmd

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
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
	publicBytes, err := ioutil.ReadFile(pemPath)
	if err != nil {
		log.WithError(err).Fatalf("Failed to read pem file from the path %s", pemPath)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.WithError(err).Fatalf("Failed to convert pem file to rsa public key from the path %s", pemPath)
	}

	validatedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return publicKey, nil
	})

	switch err.(type) {
	case nil:
		if !validatedToken.Valid {
			return
		}
	case *jwt.ValidationError: // something was wrong during the validation
		vErr := err.(*jwt.ValidationError)
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return
		default:
			return
		}
	default: // something else went wrong
		return
	}
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
