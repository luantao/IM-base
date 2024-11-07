package consts

type CountType int

const (
	DbStatusApplay = 1 // 申请中
	DbStatusAccept = 2 // 已接受
	DbStatusReject = 3 // 已拒绝

	DbEnableNormal  = 1 // 正常
	DbEnableDisable = 2 // 禁用

	DbDelNormal = 1 // 正常
	DbDelDelete = 2 // 删除

	DBAuditIng = 0 //审核中
	DBAuditOk  = 1 //审核通过
	DBAuditNo  = 2 //审核拒绝

	CountTypeIncrease CountType = 1 // 增加
	CountTypeReduce   CountType = 2 // 减少

	GenderNode   = 0
	GenderMale   = 1
	GenderFemale = 2

	ListFillterNone  = 0x00 //不过滤
	ListFilterDelete = 0x01 //过滤删除的
	ListFilterAudit  = 0x02 //过滤未审核通过的
)
