package common

import "fmt"

func GetKey(n int) []byte {
	return []byte("key_" + fmt.Sprintf("%07d", n))
}

func GeyValue64B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func GeyValue128B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func GeyValue256B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func GeyValue512B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv" + "valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}
