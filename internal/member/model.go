package member

type MemberMetadata struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

// Member model info
// @Description User account information
type Member struct {
	// 钱包地址
	Address string `json:"address" bson:"address"`
	// 昵称
	Nickname string `json:"nickname" bson:"nickname"`
	// discord 账号
	Discord string `json:"discord" bson:"discord"`
	// 描述
	Description string `json:"description" bson:"description"`
	Email       string `json:"email" bson:"email"`
	// 社交账号
	Contact []string `json:"contact" bson:"contact"`
	// SeeDao中的身份或角色
	Identity []string `json:"identity" bson:"identity"`
	// 参与项目
	Projects []string `json:"projects" bson:"projects"`
	// 其他信息
	Metadata []MemberMetadata `json:"metadata" bson:"metadata"`

	CreatedAt int64 `json:"created_at" bson:"created_at"`

	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
	// 从discord中获取
	// Guilds      string `json:"guilds" bson:"guilds"`
}
