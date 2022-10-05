package cmd

import (
	"log"
	"net/http"

	"github.com/enrichman/generic-client/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Execute() {
	rootCmd := newRootCmd()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Host    string
	Verbose bool
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "genctl",
		Short:             "A brief description of your application",
		PersistentPreRunE: initializeConfigPreRun,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	config := &Config{}

	// init flags
	bindRootFlags(rootCmd.Flags(), rootCmd.PersistentFlags(), config)

	ep := client.NewClient(http.DefaultClient, "http://localhost")

	// add commands
	rootCmd.AddCommand(newUserCmd(config, ep.Users))

	return rootCmd
}

func bindRootFlags(flags, persistentFlags *pflag.FlagSet, config *Config) {
	flags.BoolP("verbose", "v", false, "enable verbose")
	persistentFlags.StringVar(&config.Host, "host", "", "Set your name")
}
