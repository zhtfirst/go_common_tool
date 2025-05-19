package tool

import "encoding/json"

// StructToMapByJson
//
//	@Description: 结构体转map
//	@param s
//	@return map[string]any
//	@return error
func StructToMapByJson(s any) (map[string]any, error) {
	sByte, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	m := make(map[string]any)
	err = json.Unmarshal(sByte, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
