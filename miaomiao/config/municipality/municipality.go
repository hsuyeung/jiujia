package municipality

// Municipality 直辖市
type Municipality struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Municipalities 存储直辖市数据
var Municipalities = []Municipality{
	{Code: "1101", Name: "北京市"},
	{Code: "1201", Name: "天津市"},
	{Code: "3101", Name: "上海市"},
	{Code: "5001", Name: "重庆市"},
}
