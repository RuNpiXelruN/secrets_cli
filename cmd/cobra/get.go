package cobra

import (
	"fmt"

	secret "github.com/RuNpiXelruN/secrets-cli-app"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Decrypts and returns a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(encodingKey) == 0 {
			fmt.Print("\nYou need to pass your secure encoding key.\neg,\n'secrets get -k  <your enc key> some_api_key'\n\n")
			return
		}
		v := secret.NewFileVault(encodingKey, secretsPath())
		key := args[0]
		val, err := v.Get(key)
		if err != nil {
			fmt.Printf("Something went wrong retrieving your secret - %v\n\nPlease ensure you passed the correct encryption key.\neg,\n'secrets get -k  <your enc key> some_api_key'\n\n", err)
			return
		}

		fmt.Printf("\n\t%v: %v\n\n", key, val)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
