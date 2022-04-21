package command

import (
	"github.com/MinseokOh/data-warehouse/transformer"
	"github.com/spf13/cobra"
)

func NewTransformerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transformer",
		Short: "transformer commands",
	}

	cmd.AddCommand(RunTransformer())

	return cmd
}

func RunTransformer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run producer application",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := cmd.Flags().GetString(flagConfig)
			if err != nil {
				return err
			}

			trans := transformer.NewTransformer(config)
			trans.Run()

			return nil
		},
	}

	cmd.Flags().String(flagConfig, "", "set transformer config")

	return cmd
}
