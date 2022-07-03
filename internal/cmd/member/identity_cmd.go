package membercmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/SeeDAO-OpenSource/sgn/internal/member"
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/spf13/cobra"
)

type dataType struct {
	DiscordName string `json:"discordName"`
	Address     string `json:"address"`
	Contact     string `json:"contact"`
	Guilds      string `json:"guilds"`
	Projects    string `json:"projects"`
	Description string `json:"description"`
	TokenId     int64  `json:"tokenId"`
	TokenUrl    string `json:"tokenUrl"`
}

func NewIdentityCmd(builder *app.AppBuilder) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identity",
		Short: "identity",
		Long:  "identity",
	}
	cmd.AddCommand(newImportDataCmd(builder))
	return cmd
}

func newImportDataCmd(builder *app.AppBuilder) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "导入identity json数据",
		Long:  "导入identity json数据",
		RunE:  func(cmd *cobra.Command, args []string) error { return importData(cmd) },
	}
	cmd.PersistentFlags().StringP("file", "f", "", "数据文件")
	cmd.MarkFlagRequired("file")
	return (cmd)
}

func importData(cmd *cobra.Command) error {
	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	data := make([]dataType, 0)
	if err := json.Unmarshal(content, &data); err != nil {
		return err
	}
	members := make([]member.Member, len(data))
	for i, item := range data {
		members[i] = toMember(item)
	}
	service := di.Get[member.MemberService]()
	if service == nil {
		return errors.New("member service is nil")
	}
	if err := service.InsertManay(members); err != nil {
		return err
	}
	return nil
}

func toMember(item dataType) member.Member {
	m := member.Member{
		Address:     item.Address,
		Discord:     item.DiscordName,
		Description: item.Description,
	}
	if item.Projects != "" {
		m.Projects = strings.Split(item.Projects, ",")
	}
	if item.Contact != "" {
		m.Contact = strings.Split(item.Contact, ",")
	}
	if m.Nickname == "" {
		if m.Discord != "" {
			m.Nickname = strings.Split(m.Discord, "#")[0]
		}
		if m.Nickname == "" {
			m.Nickname = m.Address
		}
	}
	return m
}
