package conf

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)



func ParseConfInfo() GlobalConfig {
	var g GlobalConfig
	err := yaml.Unmarshal([]byte(Configure),&g)
	if err != nil {
		panic("Global Config Error!")
	}
	path,err := GetDir(g.GatewayConfPath)
	if err != nil {
		panic("Gateway Config Path Error!")
	}
	fmt.Println(path)
	gatewayList := getGatewayList(path)
	g.GatewayList = gatewayList
	return g
}

func getGatewayList(path []string) []GatewayInfo {
	gatewayList := make([]GatewayInfo,0)
	PthSep := string(os.PathSeparator)
	for _,p := range path {
		var gateway GatewayInfo
		c,err := ioutil.ReadFile(p + PthSep + "gateway.conf")
		if err != nil {

			panic("Gateway Config Path Error! Error path: " + p)
		}
		err = yaml.Unmarshal(c,&gateway)
		if err != nil {
			panic("Gateway Config Error! Error path: " + p)
		}
		if gateway.GatewayStatus != "on" {
			continue
		}
		gateway.ApiList = getApiList(gateway.ApiConfPath)
		gateway.StrategyList = getStrategyList(gateway.StrategyConfPath)
		gateway.BackendList = getBackendList(gateway.BackendConfPath)
		gatewayList = append(gatewayList,gateway)
	}
	return gatewayList
}

func getStrategyList(path string) Strategy {
	var strategy Strategy
	c,err := ioutil.ReadFile(path)
	if err != nil {
		panic("Strategy Config Path Error! Error path: " + path)
	}
	err = yaml.Unmarshal(c,&strategy)
	if err != nil {
		panic("Strategy Config Error! Error path: " + path)
	}
	return strategy
}

func getApiList(path string) Api {
	var api Api
	c,err := ioutil.ReadFile(path)
	if err != nil {
		panic("Api Config Path Error! Error path: " + path)
	}
	err = yaml.Unmarshal(c,&api)
	if err != nil {
		panic("Api Config Error! Error path: " + path)
	}
	return api
}

func getBackendList(path string) Backend {
	var backend Backend
	c,err := ioutil.ReadFile(path)
	if err != nil {
		panic("Backend Config Path Error! Error path: " + path)
	}
	err = yaml.Unmarshal(c,&backend)
	if err != nil {
		panic("Error backend config! Error path: " + path)
	}
	return backend
}
