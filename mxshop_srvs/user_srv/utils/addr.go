/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-24 22:45:16
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-24 22:45:17
 */
package utils

import (
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
