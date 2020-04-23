package cobra

import (
	"fmt"

	secret "github.com/RuNpiXelruN/secrets-cli-app"
	"github.com/spf13/cobra"
)

var removeKeyCmd = &cobra.Command{
	Use:   "removeKey",
	Short: "Delete a key/value from your secret store",
	Run: func(cmd *cobra.Command, args []string) {
		if len(encodingKey) == 0 {
			fmt.Print("\nYou need to pass your secure encoding key.\neg,\n'secrets removeKey -k <your enc key>'\n\n")
			return
		}

		v := secret.NewFileVault(encodingKey, secretsPath())
		key := args[0]
		err := v.RemoveKey(key)
		if err != nil {
			fmt.Printf("Something went wrong removing your secret key/value - %v\n\nPlease ensure you passed the correct encryption key.\neg,\n'secrets removeKey -k <your enc key>'\n\n", err)
			return
		}

		fmt.Printf("\n%v was successfully removed.\n", key)
	},
}

func init() {
	RootCmd.AddCommand(removeKeyCmd)
}
