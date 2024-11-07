package consts

import (
	"errors"
	"strings"
)

type ConvType int

const (
	ConvTypeC2C   ConvType = 1
	ConvTypeGroup ConvType = 2
)

const (
	ConvTypeC2CStr   string = "c2c"
	ConvTypeGroupStr string = "group"
)

type V2TimConversation struct {
	ConversationID  string              `json:"conversationID"` //会话唯一 ID，如果是 C2C 单聊，组成方式为 c2c_userID，如果是群聊，组成方式为 group_groupID
	Type            ConvType            `json:"type"`           //会话类型 1:个人 2:群组
	UserID          string              `json:"userID"`         //如果会话类型为 C2C 单聊，userID 会存储对方的用户ID，否则为 nil
	GroupID         string              `json:"groupID"`        //如果会话类型为群聊，groupID 会存储当前群的群 ID，否则为 nil
	ShowName        string              `json:"showName"`       //会话展示名称（群组：群名称 >> 群 ID；C2C：对方好友备注 >> 对方昵称 >> 对方的 userID）
	FaceUrl         string              `json:"faceUrl"`
	GroupType       string              `json:"groupType"` //如果会话类型为群聊，groupType 为当前群类型，否则为 nil
	UnreadCount     int                 `json:"unreadCount"`
	LastMessage     *V2TimMessage       `json:"lastMessage"`
	DraftText       string              `json:"draftText"`       //草稿信息，设置草稿信息请调用 setConversationDraft() 接口
	DraftTimestamp  int                 `json:"draftTimestamp"`  //草稿编辑时间，草稿设置的时候自动生成
	IsPinned        bool                `json:"isPinned"`        //是否置顶
	RecvOpt         int                 `json:"recvOpt"`         //消息接收选项（接收 | 接收但不提醒 | 不接收）
	GroupAtInfoList []*V2TimGroupAtInfo `json:"groupAtInfoList"` //群会话 @ 信息列表，用于展示 “有人@我” 或 “@所有人” 这两种提醒状态
	OrderKey        int                 `json:"orderkey"`        //排序字段
	MarkList        []int               `json:"markList"`        //会话标记列表，取值详见 @V2TIMConversationMarkType
	//V2TIM_CONVERSATION_MARK_TYPE_STAR 会话标星
	//V2TIM_CONVERSATION_MARK_TYPE_UNREAD 会话标记未读（重要会话）
	//V2TIM_CONVERSATION_MARK_TYPE_FOLD 会话折叠
	//V2TIM_CONVERSATION_MARK_TYPE_HIDE 会话隐藏
	CustomData            string   `json:"customData"`            //会话自定义数据
	ConversationGroupList []string `json:"conversationGroupList"` //会话所属分组列表
}

type V2TimMergerElem struct {
	V2TIMElem
	IsLayersOverLimit bool     `json:"isLayersOverLimit"`
	Title             string   `json:"title"`
	AbstractList      []string `json:"abstractList"`
}

func GetConvTypeAndStr(conversationID string) (convType ConvType, userID string, err error) {
	getInfo := strings.Split(conversationID, "_")
	if len(getInfo) != 2 {
		err = errors.New("conversationID 格式不正确")
		return
	}
	switch getInfo[0] {
	case ConvTypeC2CStr:
		convType = ConvTypeC2C
	case ConvTypeGroupStr:
		convType = ConvTypeGroup
	}
	userID = getInfo[1]
	return
}
