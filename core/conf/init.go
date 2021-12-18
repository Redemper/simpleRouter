package conf

import "simpleRouter/core/pojo"

func init() {
	yc := new(pojo.YamlConf)
	ReadYamlFromDefaultPath(yc)
	YamlConf = yc
}
