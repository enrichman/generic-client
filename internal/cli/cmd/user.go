package cmd

import (
	"github.com/enrichman/generic-client/internal/gencli"
	"github.com/spf13/cobra"
)

func newUserCmd(config *Config, userService gencli.UserService) *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "A brief description of your application",
	}

	userCmd.AddCommand(newUserListCmd(config, userService))
	userCmd.AddCommand(newUserCreateCmd(config, userService))

	return userCmd
}
