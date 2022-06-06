package cmd

import (
	"github.com/spf13/cobra"
)

type userService interface {
	userLister
}

func newUserCmd(config *Config, userService userService) *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "A brief description of your application",
	}

	userCmd.AddCommand(newUserListCmd(config, userService))

	return userCmd
}
