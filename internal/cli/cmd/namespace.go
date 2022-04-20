package cmd

import (
	"context"
	"fmt"

	"github.com/enrichman/epinio-client-go/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type namespaceService interface {
	namespaceLister
}

func newNamespaceCmd(config *Config, namespaceService namespaceService) *cobra.Command {
	namespaceCmd := &cobra.Command{
		Use:   "namespace",
		Short: "A brief description of your application",
	}

	namespaceCmd.AddCommand(newNamespaceListCmd(config, namespaceService))

	return namespaceCmd
}

type namespaceLister interface {
	List(ctx context.Context) ([]client.Namespace, error)
}

func newNamespaceListCmd(config *Config, namespaceLister namespaceLister) *cobra.Command {
	namespaceListConfig := config.NamespaceConfig.NamespaceListConfig

	namespaceListCmd := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			namespaceLister.List(context.Background())

			out := cmd.OutOrStdout()
			fmt.Fprintln(out, "namespaceListConfig : ", namespaceListConfig.Name)
			fmt.Fprintln(out, "config : ", config.Name)
		},
	}

	bindNamespaceListConfig(namespaceListCmd.Flags(), namespaceListCmd.PersistentFlags(), &namespaceListConfig)

	return namespaceListCmd
}

func bindNamespaceListConfig(flags, persistentFlags *pflag.FlagSet, config *NamespaceListConfig) {
	flags.StringVar(&config.Name, "name2", "default_name_2", "Set your name")
}
