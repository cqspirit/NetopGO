package main

import (
	_ "github.com/cqspirit/NetopGO/routers"

	"github.com/cqspirit/NetopGO/models"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strings"
	//"time"
)

func init() {
	models.RegisterDB()
	orm.RunSyncdb("default", false, false)
	//orm.RunSyncdb("default", false, true)
}

func main() {
	//orm.Debug = true
	beego.Run()
}
