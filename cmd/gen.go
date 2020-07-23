package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	jwt "github.com/dgrijalva/jwt-go"
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
	fatal(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var claimsJSON map[string]interface{}
	json.Unmarshal([]byte(byteValue), &claimsJSON)

	signBytes, err := ioutil.ReadFile(pathPem)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), jwt.MapClaims(claimsJSON))
	t.Header["kid"] = keyID

	token, err := t.SignedString(signKey)
	fatal(err)

	return token
}

func gen(cmd *cobra.Command, args []string) {
	pemPath := cmd.Flag("privatePem").Value.String()
	jsonPath := cmd.Flag("claimJson").Value.String()
	keyID := cmd.Flag("keyID").Value.String()

	fmt.Println(genToken(pemPath, jsonPath, keyID))
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringP("privatePem", "p", "", "path of private key pem file")
	genCmd.Flags().StringP("claimJson", "c", "", "path of claim json file")
	genCmd.Flags().StringP("keyID", "k", "", "id of private key")
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
