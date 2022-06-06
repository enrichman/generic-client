package cmd

import (
	"context"

	"github.com/enrichman/generic-client/internal/cli/ui"
	"github.com/enrichman/generic-client/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type userLister interface {
	List(ctx context.Context) ([]client.User, error)
}

func newUserListCmd(config *Config, userLister userLister) *cobra.Command {
	userListConfig := config.UserConfig.UserListConfig

	userListCmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.List(cmd.OutOrStdout(), userLister)
		},
	}

	bindUserListConfig(userListCmd.Flags(), userListCmd.PersistentFlags(), userListConfig)

	return userListCmd
}

func bindUserListConfig(flags, persistentFlags *pflag.FlagSet, config *UserListConfig) {
	flags.StringVar(&config.Name, "name2", config.Name, "Set your name")
}
