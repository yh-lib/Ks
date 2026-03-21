package utils

import "encoding/json"

func Struct2Map(s any) map[string]string {
	j, _ := json.Marshal(s)
	m := make(map[string]string)
	json.Unmarshal(j, &m)
	return m
}
