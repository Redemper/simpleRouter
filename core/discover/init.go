package discover

import (
	"fmt"
	"simpleRouter/core/conf"
)

func init() {
	lbConf := conf.GetLbConf()
	// TODO 增加其他discovery
	fmt.Println(lbConf)
	initNacos()
}
