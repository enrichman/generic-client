package cmd

import (
	"github.com/enrichman/generic-client/internal/cli/ui"
	"github.com/enrichman/generic-client/internal/gencli"
	"github.com/spf13/cobra"
)

func newUserCreateCmd(config *Config, userService gencli.UserService) *cobra.Command {
	return &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			u, err := userService.Create(cmd.Context(), "")
			ui.Create(cmd.OutOrStdout(), u, err)
		},
	}
}
