package main

import "fmt"

func main() {
	var var1 string
	var var2, var3 int
	var (
		var4 string
		var5 int
	)
	var var6 string = "var6"
	var (
		var7 int    = 7
		var8 string = "var8"
	)
	var (
		var9  = 9
		var10 = "var10"
	)
	var (
		var11, var12, var13 = 11, 12, "var13"
	)
	var14 := 14
	var15 := "var15"
	fmt.Println(var1, var2, var3, var4, var5, var6, var7, var8, var9, var10, var11, var12, var13, var14, var15)

	const HTTPStatusOK = 200
	const (
		StatusOK         = 0
		StatusReset      = 1
		StatusOtherError = 2
	)
	fmt.Println(HTTPStatusOK, StatusOK, StatusReset, StatusOtherError)
}
