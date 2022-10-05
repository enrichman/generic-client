package ui

import (
	"fmt"
	"io"

	"github.com/enrichman/generic-client/internal/gencli"
)

func List(writer io.Writer, users []gencli.User, err error) {
	fmt.Fprintln(writer, "Users:")
	for _, u := range users {
		fmt.Fprintf(writer, "- %s\n", u.Name)
	}
}

func Create(writer io.Writer, user gencli.User, err error) {
	fmt.Fprintln(writer, "User:")
	fmt.Fprintf(writer, "- %s\n", user.Name)
}
