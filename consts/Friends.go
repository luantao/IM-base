package consts

type V2TimFriendInfo struct {
	UserID           string             `json:"userID"`           //好友 ID
	FriendRemark     *string            `json:"friendRemark"`     //备注长度最长不得超过 96 个字节;
	FriendGroups     []string           `json:"friendGroups"`     //好友所在分组列表
	FriendCustomInfo map[string]string  `json:"friendCustomInfo"` //好友自定义字段
	UserProfile      *V2TimUserFullInfo `json:"userProfile"`      //好友个人资料
}

type V2TimUserFullInfo struct {
	UserID        *string           `json:"userID"`
	NickName      *string           `json:"nickName"`
	FaceUrl       *string           `json:"faceUrl"`
	SelfSignature *string           `json:"selfSignature"`
	Gender        *int              `json:"gender"`
	AllowType     *int              `json:"allowType"`  //用户好友验证方式
	CustomInfo    map[string]string `json:"customInfo"` //用户自定义字段
	Role          *int              `json:"role"`
	Level         *int              `json:"level"`
	Birthday      *int64            `json:"birthday"`
}

type V2TimFriendApplication struct {
	UserID     string  `json:"userID"` //用户标识
	NickName   *string `json:"nickname"`
	FaceUrl    *string `json:"faceUrl"`
	AddTime    *int    `json:"addTime"`
	AddSource  *string `json:"addSource"`  //来源
	AddWording *string `json:"addWording"` //加好友附言
	Type       int     `json:"type"`
}
