package consts

type V2TimTopicInfo struct {
	TopicID         string
	TopicName       string
	TopicFaceUrl    string
	Introduction    string
	Notification    string
	IsAllMute       bool
	SelfMuteTime    int
	CustomString    string
	RecvOpt         int
	DraftText       string
	UnreadCount     int
	LastMessage     V2TimMessage
	GroupAtInfoList []V2TimGroupAtInfo
}
