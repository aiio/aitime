package aitime

import "time"

// GetDate 获取当前时间的字符串格式
func GetDate() string {
	//当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetUnix 获取当前时间 单位:s
func GetUnix() int64 {
	return time.Now().Unix()
}

// GetUnixNano 获取当前时间 单位:纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// Unix2Date 时间戳转时间字符串
func Unix2Date(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
}

// Date2Unix 时间字符串转时间戳
func Date2Unix(formatTimeStr string) (int64, error) {
	formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	if err != nil {
		return 0, err
	}
	return formatTime.Unix(), nil
}
