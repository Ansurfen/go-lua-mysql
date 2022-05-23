package main

import (
	_ "github.com/go-sql-driver/mysql"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	Lvm := lua.NewState()
	defer Lvm.Close()
	Lvm.PreloadModule("lua-mysql", load)
	if err := Lvm.DoFile("test.lua"); err != nil {
		panic(err)
	}
}
