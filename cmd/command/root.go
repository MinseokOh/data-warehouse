package command

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "data-warehouse",
		Short: "data-warehouse commands",
	}

	rootCmd.AddCommand(NewProducerCmd())
	rootCmd.AddCommand(NewTransformerCmd())

	return rootCmd
}
