package consts

type Options struct {
	Key     string           `json:"key"`
	Value   string           `json:"value"`
	Options []SelectListItem `json:"options"`
}

type SelectListItem struct {
	Key int    `json:"key"`
	Val string `json:"value"`
}

const (
	// 未知
	EducationUnknown = iota
	// 小学
	EducationPrimary
	// 初中
	EducationMiddle
	// 中专
	EducationSecondaryVocational
	// 高中
	EducationHigh
	//大专
	EducationCollege
	//本科
	EducationUniversity
	// 研究生
	EducationGraduate
	// 博士
	EducationDoctor
	// 博士后
	EducationMaster
)

const (
	AreaProvinceType = 3
	AreaCityType     = 4
	AreaParentId     = 8 //表示中国区域
)

var IncomeMap = map[int]string{
	0:  "不限",
	1:  "1万以下",
	2:  "1-3万",
	3:  "3-5万",
	4:  "5-10万",
	5:  "10-20万",
	6:  "20-50万",
	7:  "50-100万",
	8:  "100-200万",
	9:  "200-500万",
	10: "500-1000万",
}

var IncomeList = []SelectListItem{
	{0, "不限"},
	{1, "1万以下"},
	{2, "1-3万"},
	{3, "3-5万"},
	{4, "5-10万"},
	{5, "10-20万"},
	{6, "20-50万"},
	{7, "50-100万"},
	{8, "100-200万"},
	{9, "200-500万"},
	{10, "500-1000万"},
}
var GenderMap = map[int]string{
	0: "不限",
	1: "男",
	2: "女",
}
var GenderList = []SelectListItem{
	{0, "不限"},
	{1, "男"},
	{2, "女"},
}

var CertifyMap = map[int]string{
	0: "不限",
	1: "真人认证",
	2: "认证不通过",
}

var CertifyList = []SelectListItem{
	{0, "不限"},
	{1, "真人认证"},
	{2, "认证不通过"},
}

//ToDo::地区，需要调用获取地区的方法
//ToDo::年龄，需要循环生成年龄

var EducationMap = map[int]string{
	EducationUnknown:             "不限",
	EducationPrimary:             "小学",
	EducationMiddle:              "初中",
	EducationSecondaryVocational: "中专",
	EducationHigh:                "高中",
	EducationCollege:             "大专",
	EducationUniversity:          "大学",
	EducationGraduate:            "研究生",
	EducationDoctor:              "博士",
	EducationMaster:              "博士后",
}

var EducationList = []SelectListItem{
	{EducationUnknown, "不限"},
	{EducationPrimary, "小学及以上"},
	{EducationMiddle, "初中及以上"},
	{EducationSecondaryVocational, "中专及以上"},
	{EducationHigh, "高中及以上"},
	{EducationCollege, "大专及以上"},
	{EducationUniversity, "大学及以上"},
	{EducationGraduate, "研究生及以上"},
	{EducationDoctor, "博士及以上"},
	{EducationMaster, "博士后"},
}

var CarMap = map[int]string{
	0: "不限",
	1: "已购车",
}

var CarList = []SelectListItem{
	{0, "不限"},
	{1, "已购车"},
}

var HouseMap = map[int]string{
	0: "不限",
	1: "已购房",
}

var HouseList = []SelectListItem{
	{0, "不限"},
	{1, "已购房"},
}

var HeightRange = map[int]string{
	0:  "不限",
	1:  "150-160cm",
	2:  "160-170cm",
	3:  "170-180cm",
	4:  "180-190cm",
	5:  "190-200cm",
	6:  "200-210cm",
	7:  "210-220cm",
	8:  "220-230cm",
	9:  "230-240cm",
	10: "240-250cm",
	11: "250-260cm",
	12: "260-270cm",
	13: "270-280cm",
	14: "280-290cm",
	15: "290-300cm",
}

var HeightList = []SelectListItem{
	{0, "不限"},
	{1, "150-160cm"},
	{2, "160-170cm"},
	{3, "170-180cm"},
	{4, "180-190cm"},
	{5, "190-200cm"},
	{6, "200-210cm"},
	{7, "210-220cm"},
	{8, "220-230cm"},
	{9, "230-240cm"},
	{10, "240-250cm"},
	{11, "250-260cm"},
	{12, "260-270cm"},
	{13, "270-280cm"},
	{14, "280-290cm"},
	{15, "290-300cm"},
}

var AgesRange = map[int]string{
	0:  "不限",
	1:  "18-22岁",
	2:  "23-27岁",
	3:  "28-32岁",
	4:  "33-37岁",
	5:  "38-42岁",
	6:  "43-47岁",
	7:  "48-52岁",
	8:  "53-57岁",
	9:  "58-62岁",
	10: "63-67岁",
	11: "68-72岁",
	12: "73-77岁",
	13: "78-82岁",
	14: "83-87岁",
	15: "88-92岁",
}

var AgesList = []SelectListItem{
	{0, "不限"},
	{1, "18-25岁"},
	{2, "26-30岁"},
	{3, "31-35岁"},
	{4, "36-40岁"},
	{5, "41-45岁"},
	{6, "46-50岁"},
	{7, "51-55岁"},
	{8, "56-60岁"},
	{9, "61-65岁"},
	{10, "66-70岁"},
	{11, "71-75岁"},
	{12, "76-80岁"},
	{13, "81-85岁"},
	{14, "86-90岁"},
}
