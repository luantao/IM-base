package consts

import (
	"errors"
	"strconv"
)

type V2TIMElem struct {
	NextElem map[string]interface{} `json:"nextElem"`
}

type V2TimVideoElem struct {
	V2TIMElem
	VideoPath        string `json:"videoPath"`
	UUID             string `json:"UUID"`
	VideoSize        int    `json:"videoSize"`
	Duration         int    `json:"duration"`
	SnapshotPath     string `json:"snapshotPath"`
	SnapshotUUID     string `json:"snapshotUUID"`
	SnapshotSize     int    `json:"snapshotSize"`
	SnapshotWidth    int    `json:"snapshotWidth"`
	SnapshotHeight   int    `json:"snapshotHeight"`
	VideoUrl         string `json:"videoUrl"`
	SnapshotUrl      string `json:"snapshotUrl"`
	LocalVideoUrl    string `json:"localVideoUrl"`
	LocalSnapshotUrl string `json:"localSnapshotUrl"`
}

type V2TimSoundElem struct {
	V2TIMElem
	Path     string `json:"path"`
	UUID     string `json:"UUID"`
	DataSize int    `json:"dataSize"`
	Duration int    `json:"duration"`
	URL      string `json:"url"`
	LocalURL string `json:"localUrl"`
}
type V2TimFaceElem struct {
	V2TIMElem
	Index int    `json:"index"`
	Data  string `json:"data"`
}

type V2TimLocationElem struct {
	V2TIMElem
	Desc      string  `json:"desc"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type V2TimFileElem struct {
	V2TIMElem
	Path     string `json:"path"`
	FileName string `json:"fileName"`
	UUID     string `json:"UUID"`
	URL      string `json:"url"`
	FileSize int    `json:"fileSize"`
	LocalURL string `json:"localUrl"`
}

type V2TimImageElem struct {
	V2TIMElem
	Path      string        `json:"path"`
	ImageList []*V2TimImage `json:"imageList"`
}

type V2TimImage struct {
	UUID     string `json:"uuid"`
	Type     int    `json:"type"` // 0:原图 1:大图 2:小图
	Size     int    `json:"size"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	URL      string `json:"url"`
	LocalURL string `json:"local_url"`
}
type V2TimCustomElem struct {
	V2TIMElem
	Data      string `json:"data"`
	Desc      string `json:"desc"`
	Extension string `json:"extension"`
}
type V2TimTextElem struct {
	V2TIMElem
	Text string `json:"text"`
}

type V2TimGiftElem struct {
	V2TIMElem
	GiftID     int    `json:"gift_id"`
	GiftName   string `json:"gift_name"`
	GiftURL    string `json:"gift_url"`
	GiftNum    int    `json:"gift_num"`
	GiftAction string `json:"gift_action"`
}

type MessageStatus int

const (
	V2TIM_MSG_STATUS_SENDING        MessageStatus = iota + 1 // 消息发送中
	V2TIM_MSG_STATUS_SEND_SUCC                               // 消息发送成功
	V2TIM_MSG_STATUS_SEND_FAIL                               // 消息发送失败
	V2TIM_MSG_STATUS_HAS_DELETED                             // 消息被删除
	V2TIM_MSG_STATUS_LOCAL_IMPORTED                          // 导入到本地的消息
	V2TIM_MSG_STATUS_LOCAL_REVOKED                           // 被撤销的消息
)

type ElemType int

// 1:文本消息,2:自定义消息,3:图片消息,4:语音消息,5:视频消息,6:文件消息,7:地理位置消息,8:表情消息,9:群 Tips 消息,10:合并消息
const (
	ElemTypeText        ElemType = 1
	ElemTypeCustom      ElemType = 2
	ElemTypeImage       ElemType = 3 // 发送图片时至少要有原图和缩略图
	ElemTypeSound       ElemType = 4
	ElemTypeVideo       ElemType = 5
	ElemTypeFile        ElemType = 6
	ElemTypeLocation    ElemType = 7
	ElemTypeFace        ElemType = 8
	ElemTypeGroupTip    ElemType = 9
	ElemTypeMerge       ElemType = 10
	ElemTypeTimeDivider ElemType = 11 // 时间分割线（APP端）
	ElemTypeGift        ElemType = 60 //新增的礼物消息
	ElemTypeSysMessage  ElemType = 101
)

var ElemTypeMap = map[ElemType]string{
	ElemTypeText:        "Text",
	ElemTypeCustom:      "Custom",
	ElemTypeImage:       "Image",
	ElemTypeSound:       "Sound",
	ElemTypeVideo:       "Video",
	ElemTypeFile:        "File",
	ElemTypeLocation:    "Location",
	ElemTypeFace:        "Face",
	ElemTypeGroupTip:    "GroupTip",
	ElemTypeMerge:       "Merge",
	ElemTypeTimeDivider: "TimeDivider",
	ElemTypeGift:        "Gift",
	ElemTypeSysMessage:  "SysMessage",
}

type V2TimMessage struct {
	ID                        string              `json:"id"`           //ConvID 客户端给的会话ID
	MsgID                     string              `json:"msgID"`        //消息ID:消息创建的时候为空，调用sendMessage的时候同步返回。
	Timestamp                 int64               `json:"timestamp"`    //消息时间戳,消息发送到服务端的时间。可用于消息排序。
	Progress                  int                 `json:"progress"`     //消息中文件上传进度 取值范围为0-100
	Sender                    string              `json:"sender"`       //消息发送者:客户自己设置，跟login时传入的userID一致。
	NickName                  string              `json:"nickName"`     //消息发送者昵称,客户自己设置。调用 setSelfInfo 设置及修改
	FriendRemark              string              `json:"friendRemark"` //消息发送者好友备注,接收方使用。例如 alice 给好友 bob 备注为 "bob01"。当 bob 给 alice 发消息，此时对于 alice 而言，消息中的 friendRemark 为 "bob01"。调用 setFriendInfo 设置。
	FaceUrl                   string              `json:"faceUrl"`      //消息发送者头像
	NameCard                  string              `json:"nameCard"`     //如果是群组消息，nameCard 为发送者的群名片,例如 alice 修改自己的群名片为 "doctorA"，那么 alice 往群里发送的消息，群成员收到的消息 nameCard 字段值为 "doctorA"。接收者可以将这个字段优先作为用户名称的显示。调用 setGroupMemberInfo 设置。
	GroupID                   string              `json:"groupID"`      //如果是群组消息，groupID 为会话群组 ID，否则为 nil
	UserID                    string              `json:"userID"`       //如果是单聊消息，userID 为会话用户 ID，否则为 nil， 假设自己和 userA 聊天，无论是自己发给 userA 的消息还是 userA 发给自己的消息，这里的 userID 均为 userA
	Status                    MessageStatus       `json:"status"`       //消息发送状态:1:消息发送中,2:消息发送成功,3:消息发送失败,4:消息被删除,5:导入到本地的消息,6:被撤回的消息
	StatusDesc                string              `json:"statusDesc"`   //消息发送状态描述，这里可以写具体原因
	ElemType                  ElemType            `json:"elemType"`     //消息类型:1:文本消息,2:自定义消息,3:图片消息,4:语音消息,5:视频消息,6:文件消息,7:地理位置消息,8:表情消息,9:群 Tips 消息,10:合并消息
	TextElem                  *V2TimTextElem      `json:"textElem"`     //普通的文字消息。
	CustomElem                *V2TimCustomElem    `json:"customElem"`   //一段二进制 buffer，通常用于传输您应用中的自定义信令。
	ImageElem                 *V2TimImageElem     `json:"imageElem"`    //SDK 会在发送原始图片的同时，自动生成两种不同尺寸的缩略图，三张图分别被称为原图、大图、微缩图。
	SoundElem                 *V2TimSoundElem     `json:"soundElem"`    //支持语音是否播放红点展示。
	VideoElem                 *V2TimVideoElem     `json:"videoElem"`    //一条视频消息包含一个视频文件和一张配套的缩略图。
	FileElem                  *V2TimFileElem      `json:"fileElem"`     //文件消息最大支持100MB。
	LocationElem              *V2TimLocationElem  `json:"locationElem"` //理位置消息由位置描述、经度（longitude ）和纬度（latitude）三个字段组成。
	FaceElem                  *V2TimFaceElem      `json:"faceElem"`
	GroupTipsElem             *V2TimGroupTipsElem `json:"groupTipsElem"`
	MergerElem                *V2TimMergerElem    `json:"mergerElem"`                //最大支持 300 条消息合并。
	GiftElem                  *V2TimGiftElem      `json:"giftElem"`                  //礼物类型
	LocalCustomData           string              `json:"localCustomData"`           //消息自定义数据（本地保存，不会发送到对端，程序卸载重装后失效）
	LocalCustomInt            int                 `json:"localCustomInt"`            //消息自定义数据,可以用来标记语音、视频消息是否已经播放（本地保存，不会发送到对端，程序卸载重装后失效）
	CloudCustomData           string              `json:"cloudCustomData"`           //消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）
	IsSelf                    bool                `json:"isSelf"`                    //消息发送者是否是自己
	IsRead                    bool                `json:"isRead"`                    //消息自己是否已读
	IsPeerRead                bool                `json:"isPeerRead"`                //消息对方是否已读（只有 C2C 消息有效） 该字段为true的条件是消息 timestamp <= 对端标记会话已读的时间
	Priority                  int                 `json:"priority"`                  //消息优先级（只有 onRecvNewMessage 收到的 V2TIMMessage 获取有效）
	OfflinePushInfo           *OfflinePushInfo    `json:"offlinePushInfo"`           //消息的离线推送信息
	GroupAtUserList           []string            `json:"groupAtUserList"`           //群消息中被 @ 的用户 UserID 列表（即该消息都 @ 了哪些人）
	Seq                       string              `json:"seq"`                       //群聊中的消息序列号云端生成，在群里是严格递增且唯一的, 单聊中的序列号是本地生成，不能保证严格递增且唯一。
	Random                    int                 `json:"random"`                    //消息随机码
	IsExcludedFromUnreadCount bool                `json:"isExcludedFromUnreadCount"` //消息是否不计入会话未读数：默认为 NO，表明需要计入会话未读数，设置为 YES，表明不需要计入会话未读数
	IsExcludedFromLastMessage bool                `json:"isExcludedFromLastMessage"` //消息是否不计入会话未读数：默认为 NO，表明需要计入会话未读数，设置为 YES，表明不需要计入会话未读数
	IsSupportMessageExtension bool                `json:"isSupportMessageExtension"` //消息是否不计入会话未读数：默认为 NO，表明需要计入会话未读数，设置为 YES，表明不需要计入会话未读数
	MessageFromWeb            string              `json:"messageFromWeb"`            //WEB端传递到flutter端的文本数据
	NeedReadReceipt           bool                `json:"needReadReceipt"`           //消息是否需要已读回执（需要您购买旗舰版套餐）群消息在使用该功能之前，需要先到 IM 控制台设置已读回执支持的群类型
	Score                     int                 `json:"score"`                     //消息积分
	Coin                      int                 `json:"coin"`                      //消息硬币
	SensitiveList             []string            `json:"sensitiveList“`             //敏感词
}

type OfflinePushInfo struct {
	Title                     string `json:"title"`
	Desc                      string `json:"desc"`
	Ext                       string `json:"ext"`
	DisablePush               bool   `json:"disablePush"`
	IOSSound                  string `json:"iOS_Sound"`
	IgnoreIOSBadge            bool   `json:"ignoreIOSBadge"`
	AndroidOPPOChannelID      string `json:"androidOPPOChannelID"`
	AndroidVIVOClassification int    `json:"androidVIVOClassification"`
	AndroidSound              string `json:"android_Sound"`
	AndroidFCMChannelID       string `json:"androidFCMChannelID"`
	AndroidXiaoMiChannelID    string `json:"androidXiaoMiChannelID"`
	IOSPushType               int    `json:"iOS_pushType"`
	AndroidHuaWeiCategory     string `json:"androidHuaWeiCategory"`
}

type V2TimMessageReceipt struct {
	UserID      string `json:"userID"`
	Timestamp   int    `json:"timestamp"`
	GroupID     string `json:"groupID"`
	MsgID       string `json:"msgID"`
	ReadCount   int    `json:"readCount"`
	UnreadCount int    `json:"unreadCount"`
}

type V2TimMessageExtension struct {
	ExtensionKey   string `json:"extensionKey"`
	ExtensionValue string `json:"extensionValue"`
}

type V2TimMessageDownloadProgress struct {
	IsFinish    bool   `json:"isFinish"`
	IsError     bool   `json:"isError"`
	MsgID       string `json:"msgID"`
	CurrentSize int    `json:"currentSize"`
	TotalSize   int    `json:"totalSize"`
	Type        int    `json:"type"`
	IsSnapshot  bool   `json:"isSnapshot"`
	Path        string `json:"path"`
	ErrorCode   int    `json:"errorCode"`
	ErrorDesc   string `json:"errorDesc"`
}

type V2TimGroupAtInfo struct {
	Seq    string `json:"seq"`
	AtType int    `json:"atType"`
}

func GetV2TimMessageClientTimestamp(msgID string) (int64, error) {
	msgIDBytes := []byte(msgID)
	if len(msgIDBytes) < 13 {
		return 0, errors.New("msgID is too short")
	}
	return strconv.ParseInt(string(msgIDBytes[:13]), 10, 64)
}
