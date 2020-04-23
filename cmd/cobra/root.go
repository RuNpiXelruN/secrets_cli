package cobra

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// RootCmd var
var RootCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Secrets is an encrypted key/value password manager",
}

// EncodingKey var
var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "Your encryption key passed used when encoding and decoding secrets")
	RootCmd.MarkFlagRequired("key")
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
