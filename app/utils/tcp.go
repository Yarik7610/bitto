package utils

import "net"

func GetRemoteAddrString(conn net.Conn) string {
	return conn.RemoteAddr().String()
}
