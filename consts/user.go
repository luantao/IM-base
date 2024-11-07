package consts

type V2TimUserInfo struct {
	UserID   string `json:"userID"`
	NickName string `json:"nickName"`
	FaceUrl  string `json:"faceUrl"`
}

const (
	UserTypeReal = 0
	UserTypeFake = 1
)

const AdminTypeSuper = 1
