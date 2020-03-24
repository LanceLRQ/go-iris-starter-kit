package utils

import (
	"bytes"
	"encoding/json"
	"time"
)


// 对象转JSON字符串
func ObjectToJSONString(obj interface{}, format bool) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return "{}"
	} else {
		if format {
			var out bytes.Buffer
			err = json.Indent(&out, b, "", "    ")
			if err != nil {
				return "{}"
			}
			return out.String()
		}
		return string(b)
	}
}

// JSON字符串转对象
func JSONStringToObject(jsonStr string, obj interface{}) bool {
	return JSONStringByteToObject([]byte(jsonStr), obj)
}

// JSON字节流转对象
func JSONStringByteToObject(jsonStrByte []byte, obj interface{}) bool {
	err := json.Unmarshal(jsonStrByte, &obj)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 获取当前系统时间并按照中国的格式标准表达
func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05.999")
}
