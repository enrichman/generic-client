package cmd

import (
	"context"
	"net/http"
	"os"

	"github.com/enrichman/generic-client/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Execute() {
	rootCmd := newRootCmd()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	config := newConfig()

	rootCmd := &cobra.Command{
		Use:               "epinioctl",
		Short:             "A brief description of your application",
		PersistentPreRunE: initializeConfigPreRun,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	// init flags
	bindRootFlags(rootCmd.Flags(), rootCmd.PersistentFlags(), config)

	ep := client.NewClient(http.DefaultClient, "http://localhost")
	ep.Users.List(context.Background())

	// add commands
	rootCmd.AddCommand(newUserCmd(config, ep.Users))

	return rootCmd
}

func bindRootFlags(flags, persistentFlags *pflag.FlagSet, config *Config) {
	flags.BoolP("toggle", "t", false, "Help message for toggle")
	persistentFlags.StringVarP(&config.Name, "name", "n", "default_name", "Set your name")
}
