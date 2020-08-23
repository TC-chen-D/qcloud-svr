/*
 * Created by Chen on Fri.Aug.2020
 */
package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/tcchend/qcloud-svr/models"
)

type Conf struct {
	models.InstancesRequest
	models.LogConf
	models.ApiUrl
}

func (c *Conf) GetConf() (*Conf, error) {
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(yamlFile, &c); err != nil {
		return nil, err
	}
	return c, nil
}



