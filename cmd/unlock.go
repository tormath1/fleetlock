package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/tormath1/fleetlock/pkg/client"
)

var (
	unlock = &cobra.Command{
		Use: "unlock-if-held",
		RunE: func(cmd *cobra.Command, args []string) error {
			httpClient := http.DefaultClient

			c, err := client.New(url, group, id, httpClient)
			if err != nil {
				return fmt.Errorf("building the client: %w", err)
			}

			if err := c.UnlockIfHeld(); err != nil {
				return fmt.Errorf("unlocking: %w", err)
			}

			return nil
		},
	}
)
