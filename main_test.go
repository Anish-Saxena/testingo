package Isip

import "testing"

func TestCheckIP(t *testing.T) {
	testip := [10]string{
		"0.0.0.0", "08.08.08.08", "8.8.8.8", "256.243.13.1", "123.121.111.1", "192.168.0.1", "111.111.111.111", "255.255.255.255", "-1.-1.-1.-1", "0.1.1.0"}
	testres := [10]bool{
		true, false, true, false, true, true, true, true, false, true}
	for i := 0; i < 10; i++ {
		if testres[i] != CheckIP(testip[i]) {
			t.Error("CheckIP Test ", i+1, " failed!")
		}
	}
}

func TestRemAlphaNum(t *testing.T) {
	testcase := [10]string{
		"H##eL@!)(L10", "123)(**", "AXNAOxx899", "/0", "a1b2c3A1B2C3", "!@#$%^&a#$%^&*", "~~)))(((Q)A", "lop lll lll lll", "qwerty keypad", "0"}
	testres := [10]string{
		"HeLL10", "123", "AXNAOxx899", "0", "a1b2c3A1B2C3", "a", "QA", "loplllllllll", "qwertykeypad", "0"}
	for i := 0; i < 10; i++ {
		if RemAlphaNum(testcase[i]) != testres[i] {
			t.Error("RemAlphaNum Test ", i+1, " failed!")
		}
	}
}
