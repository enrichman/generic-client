package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/enrichman/epinio-client-go/pkg/client"
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
	config := NewConfig()

	rootCmd := &cobra.Command{
		Use:   "epinioctl",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		PersistentPreRunE: initializeConfigPreRun,
		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()
			fmt.Fprintln(out, "Your name is:", config.Name)
		},
	}

	// init flags
	bindRootFlags(rootCmd.Flags(), rootCmd.PersistentFlags(), &config)

	ep := client.NewClient(http.DefaultClient, "http://localhost")

	// add commands
	rootCmd.AddCommand(newNamespaceCmd(&config, ep.Namespace))

	return rootCmd
}

func bindRootFlags(flags, persistentFlags *pflag.FlagSet, config *Config) {
	flags.BoolP("toggle", "t", false, "Help message for toggle")
	persistentFlags.StringVarP(&config.Name, "name", "n", "default_name", "Set your name")
}
