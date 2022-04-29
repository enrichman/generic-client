package cmd

import (
	"context"
	"fmt"

	"github.com/enrichman/epinio-client-go/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type namespaceLister interface {
	List(ctx context.Context) ([]client.Namespace, error)
}

type NamespaceListConfig struct {
	Name string
}

var DefaultNamespaceListConfig = NamespaceListConfig{
	Name: "default_namespace_list_name",
}

func newNamespaceListCmd(rootConfig *RootConfig, namespaceLister namespaceLister) *cobra.Command {
	namespaceListConfig := DefaultNamespaceListConfig
	rootConfig.NamespaceConfig.NamespaceListConfig = namespaceListConfig

	namespaceListCmd := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			namespaceLister.List(context.Background())

			out := cmd.OutOrStdout()
			fmt.Fprintln(out, "namespaceListConfig : ", namespaceListConfig.Name)
			fmt.Fprintln(out, "config : ", rootConfig.Name)
		},
	}

	bindNamespaceListConfig(namespaceListCmd.Flags(), namespaceListCmd.PersistentFlags(), &namespaceListConfig)

	return namespaceListCmd
}

func bindNamespaceListConfig(flags, persistentFlags *pflag.FlagSet, config *NamespaceListConfig) {
	flags.StringVar(&config.Name, "name2", config.Name, "Set your name")
}
