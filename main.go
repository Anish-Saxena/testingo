package Isip

import "strconv"

func CheckIP(ip string) bool {
	dot := 0
	ip = ip
	var dotp [3]int
	for i := 0; i < len(ip); i++ {
		if ip[i] == '.' {
			dotp[dot] = i
			dot = dot + 1
		}
	}
	if dot != 3 {
		return false
	}
	dot = 0
	start := 0
	for dot < 3 {
		nn, err := strconv.Atoi(ip[start:dotp[dot]])
		if nn <= 255 && nn >= 0 && err == nil {
			if ip[start] == '0' && start+1 != dotp[dot] {
				return false
			}
			dot = dot + 1
		} else {
			return false
		}
		start = dotp[dot-1] + 1
	}
	nn, err := strconv.Atoi(ip[dotp[2]+1:])
	if err == nil && nn <= 255 && nn >= 0 {
		if ip[start] == '0' && start+1 != len(ip) {
			return false
		}
	}
	return true
}
