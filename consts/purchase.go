package consts

const (
	PurchaseTypeWX    = 1
	PurchaseTypeQQ    = 2
	PurchaseTypePhone = 3
)

const (
	CertifyNo   = 0
	CertifyOk   = 1
	CertifyFail = 2
	CertifyWait = 3
)

var PurchaseTypeMap = map[int]string{
	PurchaseTypeWX:    "微信",
	PurchaseTypeQQ:    "QQ",
	PurchaseTypePhone: "手机",
}

var PurcharseList = []int{
	PurchaseTypeWX,
	PurchaseTypeQQ,
	PurchaseTypePhone,
}

const (
	SwitchOn  = 1
	SwitchOff = 2
)
