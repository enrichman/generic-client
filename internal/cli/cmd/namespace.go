package cmd

import (
	"github.com/spf13/cobra"
)

var DefaultNamespaceConfig = NamespaceConfig{
	NamespaceListConfig: DefaultNamespaceListConfig,
}

type namespaceService interface {
	namespaceLister
}

type NamespaceConfig struct {
	NamespaceListConfig NamespaceListConfig
}

func newNamespaceCmd(rootConfig *RootConfig, namespaceService namespaceService) *cobra.Command {
	namespaceConfig := DefaultNamespaceConfig
	rootConfig.NamespaceConfig = namespaceConfig

	namespaceCmd := &cobra.Command{
		Use:   "namespace",
		Short: "A brief description of your application",
	}

	namespaceCmd.AddCommand(newNamespaceListCmd(rootConfig, namespaceService))

	return namespaceCmd
}
