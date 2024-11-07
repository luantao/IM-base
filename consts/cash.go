package consts

type CashStatus int

const (
	CashStatusApply   CashStatus = iota + 1 // 申请中
	CashStatusSuccess                       // 申请成功
	CashStatusFail                          // 申请失败
)
