package miaomiao

// Api .
type Api struct {
	Url    string
	Method string
}

// Apis 需要用到的接口列表
type Apis struct {
	ChildRegions Api // 获取区域
	Vaccines     Api // 疫苗列表
	Member       Api // 成员列表
	St           Api // 加密参数 st
	Subscribe    Api // 秒杀
}

// Response 秒苗响应信息返回格式
type Response struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	NotOk bool   `json:"notOk"`
	Ok    bool   `json:"ok"`
	Data  interface{}
}

// Member 接种人信息
type Member struct {
	Id     int32
	Name   string
	IdCard string
}

// Vaccine 医院/疫苗列表
type Vaccine struct {
	Address     string `json:"address"`
	Id          int64  `json:"id"`
	ImgUrl      string `json:"imgUrl"`
	Name        string `json:"name"`
	StartTime   string `json:"startTime"`
	Stock       int64  `json:"stock"`
	VaccineCode string `json:"vaccineCode"` // 8802-四价 8803-九价
	VaccineName string `json:"vaccineName"`
}

// Region 区域
type Region struct {
	Code         string   `json:"code"`
	Name         string   `json:"name"`
	ChildRegions []Region `json:"cities"`
}
