package main

import (
	"go-package/demo/interface/productrepo"
)

func main() {
	// 选择"aliCloud"
	env := "aliCloud"
	// 根据env新建一个repo
	repo := productrepo.New(env)
	// 在新建的repo上进行存储
	repo.StoreProduct("HuaWei mate 40", 105)
}
