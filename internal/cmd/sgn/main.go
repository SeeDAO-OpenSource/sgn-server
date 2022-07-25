package sgn

import (
	"github.com/spf13/cobra"
)

func NewSgnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sgn",
		Short: "sgn相关功能",
		Long:  "sgn相关功能",
	}
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewCompareCmd())
	return cmd
}
