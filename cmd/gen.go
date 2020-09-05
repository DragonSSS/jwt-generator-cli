package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	jwt "github.com/dgrijalva/jwt-go"

	log "github.com/sirupsen/logrus"
)

var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "generate jwt toke with provided claims json and private key pem",
		Long:  "generate jwt toke with provided claims json and private key pem",
		Run:   gen,
	}
)

func genToken(pathPem string, pathClaims string, keyID string) string {

	jsonFile, err := os.Open(pathClaims)
	if err != nil {
		log.WithError(err).Fatalf("Failed to open claim file from the path %s", pathClaims)
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.WithError(err).Fatalf("Failed to read claim file from the path %s", pathClaims)
	}

	var claimsJSON map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &claimsJSON)
	if err != nil {
		log.WithError(err).Fatalf("Failed to deserialize claim file from the path %s", pathClaims)
	}

	signBytes, err := ioutil.ReadFile(pathPem)
	if err != nil {
		log.WithError(err).Fatalf("Failed to read pem file from the path %s", pathPem)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.WithError(err).Fatalf("Failed to convert pem file to rsa privatw key from the path %s", pathPem)
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), jwt.MapClaims(claimsJSON))
	t.Header["kid"] = keyID

	token, err := t.SignedString(signKey)
	if err != nil {
		log.WithError(err).Fatal("Failed to Generate the signing token")
	}

	return token
}

func gen(cmd *cobra.Command, args []string) {
	pemPath := cmd.Flag("privatePem").Value.String()
	jsonPath := cmd.Flag("claimJson").Value.String()
	keyID := cmd.Flag("keyID").Value.String()

	token := genToken(pemPath, jsonPath, keyID)
	log.WithFields(log.Fields{
		"token": token,
	}).Info("Jwt token generated successfully")
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringP("privatePem", "p", "", "path of private key pem file")
	err := genCmd.MarkFlagRequired("privatePem")
	if err != nil {
		log.WithError(err).WithField("parameter", "privatePem").Fatal("Failed to parse parameter")
	}

	genCmd.Flags().StringP("claimJson", "c", "", "path of claim json file")
	err = genCmd.MarkFlagRequired("claimJson")
	if err != nil {
		log.WithError(err).WithField("parameter", "claimJson").Fatal("Failed to parse parameter")
	}

	genCmd.Flags().StringP("keyID", "k", "", "id of private key")
	err = genCmd.MarkFlagRequired("keyID")
	if err != nil {
		log.WithError(err).WithField("parameter", "keyID").Fatal("Failed to parse parameter")
	}
}
