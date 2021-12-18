package discover

import (
	"fmt"
	"log"
	"simpleRouter/core/conf"
	"simpleRouter/core/pojo"
)

func init() {
	lbConf := pojo.LbConf{}
	err := conf.YamlConf.LbConf
	if err != nil {
		log.Fatal("can load discover from yaml")
	}
	// TODO 增加其他discovery
	fmt.Println(lbConf)
	// TODO 使用yaml中的配置
	initNacos()
}
