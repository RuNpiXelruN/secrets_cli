package cobra

import (
	"fmt"

	secret "github.com/RuNpiXelruN/secrets-cli-app"
	"github.com/spf13/cobra"
)

var listKeysCmd = &cobra.Command{
	Use:   "listKeys",
	Short: "Returns list of stored keys",
	Run: func(cmd *cobra.Command, args []string) {
		if len(encodingKey) == 0 {
			fmt.Print("\nYou need to pass your secure encoding key.\neg,\n'secrets getKeys -k <your enc key>'\n\n")
			return
		}
		v := secret.NewFileVault(encodingKey, secretsPath())
		keys, err := v.ListKeys()
		if err != nil {
			fmt.Printf("Something went wrong retrieving your stored keys - %v\n\nPlease ensure you passed the correct encryption key.\neg,\n'secrets getKeys -k <your enc key>'\n\n", err)
			return
		}

		if len(keys) == 0 {
			fmt.Println("You have no secrets stored.")
			return
		}

		fmt.Println("")
		for _, k := range keys {
			fmt.Printf("\t%v\n", k)
		}
		fmt.Println("")
		return
	},
}

func init() {
	RootCmd.AddCommand(listKeysCmd)
}
