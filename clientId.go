package xnet

import (
	"strconv"
	"strings"
)

func GetClientId(addr string) string {
	tmp := strings.Split(addr, ":")
	if len(tmp) != 2 {
		panic("zsab3f73py GetClientId  len(tmp) != 2")
	}
	ip := tmp[0]
	port := tmp[1]
	var result string
	ipTmp := strings.Split(ip, ".")
	for _, item := range ipTmp {
		partInt, err := strconv.Atoi(item)
		if err != nil {
			partInt = 0
		}
		part := strconv.FormatInt(int64(partInt), 16)
		if len(part) == 1 {
			part = "0" + part
		}
		result += part
	}
	for len(port) < 4 {
		port = "0" + port
	}
	result += port
	return result
}
