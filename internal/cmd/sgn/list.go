package sgn

import (
	"os"
	"strconv"
	"strings"

	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/sgn"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "获取sgn列表",
		Long:  "获取sgn列表",
		RunE:  excuteList,
	}
	cmd.PersistentFlags().StringP("output", "o", "./tokens.csv", "输出文件")
	return cmd
}

func excuteList(cmd *cobra.Command, args []string) error {
	var service = di.Get[sgn.SgnService]()
	data, error := service.GetTransferLogs("0x23fDA8a873e9E46Dbe51c78754dddccFbC41CFE1")
	if error != nil {
		return error
	}
	var sb = strings.Builder{}
	sb.Write([]byte("tokenId,"))
	sb.Write([]byte("transHash,"))
	sb.Write([]byte("From,"))
	sb.Write([]byte("To,"))
	sb.Write([]byte("Time,"))
	for _, v := range data {
		sb.Write([]byte(strconv.FormatInt(v.TokenID, 10)))
		sb.Write([]byte(","))
		sb.Write([]byte(v.Hash))
		sb.Write([]byte(","))
		sb.Write([]byte(v.From))
		sb.Write([]byte(","))
		sb.Write([]byte(v.To))
		sb.Write([]byte(","))
		sb.Write([]byte(v.TimeStamp.Format("2006-01-02 15:04:05")))
		sb.Write([]byte("\n"))
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	os.WriteFile(output, []byte(sb.String()), 0644)
	return nil
}
