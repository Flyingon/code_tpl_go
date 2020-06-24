package util

import "net"

// GetLocalIp 获取当前服务器IP
func GetLocalIP() string {
    addrSlice, err := net.InterfaceAddrs()
    if nil != err {
        return "localhost"
    }
    for _, addr := range addrSlice {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if nil != ipnet.IP.To4() {
                return ipnet.IP.String()
            }
        }
    }
    return "localhost"
}
