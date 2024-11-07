package consts

type V2TimGroupMemberInfo struct {
	UserID       string `json:"userID"`
	NickName     string `json:"nickName"`
	NameCard     string `json:"nameCard"`
	FriendRemark string `json:"friendRemark"`
	FaceUrl      string `json:"faceUrl"`
}

type V2TimGroupTipsElem struct {
	GroupID              string                        `json:"groupID"`
	Type                 int                           `json:"type"`
	OpMember             *V2TimGroupMemberInfo         `json:"opMember"` //操作者群成员资料
	MemberList           []*V2TimGroupMemberInfo       `json:"memberList"`
	GroupChangeInfoList  []*V2TimGroupChangeInfo       `json:"groupChangeInfoList"`  //群信息变更（type = V2TIM_GROUP_TIPS_TYPE_INFO_CHANGE 时有效）
	MemberChangeInfoList []*V2TimGroupMemberChangeInfo `json:"memberChangeInfoList"` //成员变更（type = V2TIM_GROUP_TIPS_TYPE_MEMBER_INFO_CHANGE 时有效）
	MemberCount          *int                          `json:"memberCount"`
}
type V2TimGroupChangeInfo struct {
	Type      *int    `json:"type"`
	Value     *string `json:"value"`
	Key       *string `json:"key"`
	BoolValue *bool   `json:"boolValue"`
}

type V2TimGroupMemberChangeInfo struct {
	UserID   *string `json:"userID"`
	MuteTime *int    `json:"muteTime"`
}
