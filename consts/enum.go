package consts

type MessageElemType int

const (
	// 没有元素
	V2TIM_ELEM_TYPE_NONE MessageElemType = iota
	// 文本消息
	V2TIM_ELEM_TYPE_TEXT
	// 自定义消息
	V2TIM_ELEM_TYPE_CUSTOM
	// 图片消息
	V2TIM_ELEM_TYPE_IMAGE
	// 语音消息
	V2TIM_ELEM_TYPE_SOUND
	// 视频消息
	V2TIM_ELEM_TYPE_VIDEO
	// 文件消息
	V2TIM_ELEM_TYPE_FILE
	// 地理位置消息
	V2TIM_ELEM_TYPE_LOCATION
	// 表情消息
	V2TIM_ELEM_TYPE_FACE
	// 群 Tips 消息（存消息列表）
	V2TIM_ELEM_TYPE_GROUP_TIPS
	// 合并消息
	V2TIM_ELEM_TYPE_MERGER
)

type V2TIMUserStatusType int

const (
	// 未知状态
	V2TIM_USER_STATUS_UNKNOWN V2TIMUserStatusType = iota
	// 在线状态
	V2TIM_USER_STATUS_ONLINE
	// 离线状态
	V2TIM_USER_STATUS_OFFLINE
	// 未登录（如主动调用 logout 接口，或者账号注册后还未登录）
	V2TIM_USER_STATUS_UNLOGINED
)

type V2TIMGender int

const (
	// 未知性别
	V2TIM_GENDER_UNKNOWN V2TIMGender = iota
	// 男性
	V2TIM_GENDER_MALE
	// 女性
	V2TIM_GENDER_FEMALE
)

type V2TIMFriendAllowType int

const (
	// 同意任何用户加好友
	V2TIM_FRIEND_ALLOW_ANY V2TIMFriendAllowType = iota
	// 需要验证
	V2TIM_FRIEND_NEED_CONFIRM
	// 拒绝任何人加好友
	V2TIM_FRIEND_DENY_ANY
)

type V2TIMFriendApplicationType int

const (
	// 别人发给我的
	V2TIM_FRIEND_APPLICATION_COME_IN V2TIMFriendApplicationType = iota + 1
	// 我发给别人的
	V2TIM_FRIEND_APPLICATION_SEND_OUT
	// 别人发给我的 和 我发给别人的。仅拉取时有效
	V2TIM_FRIEND_APPLICATION_BOTH
)

type V2TIMFriendRelationType int

const (
	// 不是好友
	V2TIM_FRIEND_RELATION_TYPE_NONE V2TIMFriendRelationType = 0x0
	// 对方在我的好友列表中
	V2TIM_FRIEND_RELATION_TYPE_IN_MY_FRIEND_LIST V2TIMFriendRelationType = 0x1
	// 我在对方的好友列表中
	V2TIM_FRIEND_RELATION_TYPE_IN_OTHER_FRIEND_LIST V2TIMFriendRelationType = 0x2
	// 互为好友
	V2TIM_FRIEND_RELATION_TYPE_BOTH_WAY V2TIMFriendRelationType = 0x3
)

type V2TIMFriendAcceptType int

const (
	// 接受加好友（建立单向好友）
	V2TIM_FRIEND_ACCEPT_AGREE V2TIMFriendAcceptType = iota
	// 接受加好友并加对方为好友（建立双向好友）
	V2TIM_FRIEND_ACCEPT_AGREE_AND_ADD
)

// 用户资料修改标记
type V2TIMUserInfoModifyFlag int

const (
	// 未定义
	V2TIM_USER_INFO_MODIFY_FLAG_UNKNOWN V2TIMUserInfoModifyFlag = 0
	// 昵称
	V2TIM_USER_INFO_MODIFY_FLAG_NICK V2TIMUserInfoModifyFlag = 1
	// 头像
	V2TIM_USER_INFO_MODIFY_FLAG_FACE_URL V2TIMUserInfoModifyFlag = 2
	// 性别
	V2TIM_USER_INFO_MODIFY_FLAG_GENDER V2TIMUserInfoModifyFlag = 3
	// 生日
	V2TIM_USER_INFO_MODIFY_FLAG_BIRTHDAY V2TIMUserInfoModifyFlag = 4
	// 修改签名
	V2TIM_USER_INFO_MODIFY_FLAG_SELF_SIGNATURE V2TIMUserInfoModifyFlag = 7
	// 等级
	V2TIM_USER_INFO_MODIFY_FLAG_LEVEL V2TIMUserInfoModifyFlag = 8
	// 角色
	V2TIM_USER_INFO_MODIFY_FLAG_ROLE V2TIMUserInfoModifyFlag = 9
	// 好友验证方式
	V2TIM_USER_INFO_MODIFY_FLAG_ALLOW_TYPE V2TIMUserInfoModifyFlag = 10
	// 自定义字段
	V2TIM_USER_INFO_MODIFY_FLAG_CUSTOM V2TIMUserInfoModifyFlag = 11
)

// 好友资料修改标记
type V2TIMFriendInfoModifyFlag int

const (
	// 未定义
	V2TIM_FRIEND_INFO_MODIFY_FLAG_UNKNOWN V2TIMFriendInfoModifyFlag = 0
	// 好友备注
	V2TIM_FRIEND_INFO_MODIFY_FLAG_REMARK V2TIMFriendInfoModifyFlag = 1
	// 好友自定义字段
	V2TIM_FRIEND_INFO_MODIFY_FLAG_CUSTOM V2TIMFriendInfoModifyFlag = 2
)
