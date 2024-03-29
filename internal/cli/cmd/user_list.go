package cmd

import (
	"github.com/enrichman/generic-client/internal/cli/ui"
	"github.com/enrichman/generic-client/internal/gencli"
	"github.com/spf13/cobra"
)

func newUserListCmd(config *Config, userService gencli.UserService) *cobra.Command {
	return &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			u, err := userService.List(cmd.Context())
			ui.List(cmd.OutOrStdout(), u, err)
		},
	}
}
