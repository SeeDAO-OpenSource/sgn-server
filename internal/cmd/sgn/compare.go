package sgn

import (
	"log"
	"os"
	"strings"

	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/sgn"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
	"github.com/spf13/cobra"
)

func NewCompareCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compare",
		Short: "未领取的地址列表",
		Long:  "比较白名单列表，找出未领取的地址",
		RunE:  excuteCompare,
	}
	cmd.PersistentFlags().StringP("whitelist", "i", "./whitelist.csv", "白名单申请列表文件路径")
	cmd.PersistentFlags().StringP("output", "o", "./result.csv", "未领取的地址列表保存路径")
	return cmd
}

func excuteCompare(cmd *cobra.Command, args []string) error {
	whitelist, err := cmd.Flags().GetString("whitelist")
	if err != nil {
		return err
	}
	lines, err := utils.ReadAllLines(whitelist)
	if err != nil {
		return err
	}
	log.Println(len(lines))
	whitelistSet := make(map[string]string)
	for _, line := range lines {
		addr := strings.Split(line, ",")[0]
		whitelistSet[strings.ToLower(addr)] = line
	}

	var service = di.Get[sgn.SgnService]()
	data, error := service.GetTransferLogs("0x23fDA8a873e9E46Dbe51c78754dddccFbC41CFE1")
	if error != nil {
		return error
	}
	for _, v := range data {
		delete(whitelistSet, strings.ToLower(v.To))
	}
	var sb = strings.Builder{}
	for k := range whitelistSet {
		sb.Write([]byte(k))
		sb.Write([]byte("\n"))
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	os.WriteFile(output, []byte(sb.String()), 0644)
	return nil
}
