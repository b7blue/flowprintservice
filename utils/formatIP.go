package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatIP(rawip uint32) string {
	var low uint32 = 255
	return fmt.Sprintf("%d.%d.%d.%d", rawip>>24&low, rawip>>16&low, rawip>>8&low, rawip&low)
}

func IP2uint32(ipstring string) uint32 {
	ipslice := strings.Split(ipstring, ".")
	var ip uint32 = 0
	for i := range ipslice {
		s, _ := strconv.Atoi(ipslice[i])
		ip = (ip << 8) | uint32(s)
	}
	return ip
}
