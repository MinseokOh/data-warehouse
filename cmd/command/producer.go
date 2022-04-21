package command

import (
	"github.com/MinseokOh/data-warehouse/producer"
	"github.com/spf13/cobra"
)

func NewProducerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "producer",
		Short: "producer commands",
	}

	cmd.AddCommand(RunProducer())

	return cmd
}

func RunProducer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run producer application",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := cmd.Flags().GetString(flagConfig)
			if err != nil {
				return err
			}

			prod := producer.NewProducer(config)
			prod.Run()

			return nil
		},
	}

	//cmd.Flags().String(flagData, "json1", "set producer data format(xml1|,xml2|json1|json2)")
	cmd.Flags().String(flagConfig, "", "set producer config")
	return cmd
}
