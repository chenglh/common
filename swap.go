package common

import "encoding/json"

// 通过 json tag 进行结构体赋值
func SwapTo(request, target interface{}) (err error) {
	var dataByte []byte
	if dataByte, err = json.Marshal(request); err != nil {
		return
	}

	err = json.Unmarshal(dataByte, target)
	return
}
