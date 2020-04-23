package cobra

import (
	"fmt"

	secret "github.com/RuNpiXelruN/secrets-cli-app"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Encrypts and writes a secret to your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(encodingKey) == 0 {
			fmt.Print("\nYou need to pass your secure encoding key.\neg,\n'secrets set -k  <your enc key> some_api_key some_value'\n\n")
			return
		}
		v := secret.NewFileVault(encodingKey, secretsPath())
		key, val := args[0], args[1]
		err := v.Set(key, val)
		if err != nil {
			fmt.Printf("Something went wrong writing your secret - %v\n\nPlease ensure you passed the correct encryption key.\neg,\n'secrets set -k  <your enc key> some_api_key some_value'\n\n", err)
			fmt.Printf("Unable to write to secret storage %v\n", err)
			return
		}
		fmt.Println("Secret stored successfully.")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
