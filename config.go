package main

import (
	"gopkg.in/ini.v1"
	"log"
)

func loadAPPConfig() (*ini.File, error) {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("加载配置文件错误: ", err)
	}
	//fmt.Println(cfg)
	return cfg, nil
}

func loadNacosConfig() (map[string]string, error) {
	cfg, err := ini.Load("config/nacos.ini")
	if err != nil {
		log.Fatal("加载配置文件错误: ", err)
	}

	serverInfo := map[string]string{
		"ipAddr":      cfg.Section("nacos").Key("ipAddr").String(),
		"port":        cfg.Section("nacos").Key("port").String(),
		"namespaceId": cfg.Section("nacos").Key("namespaceId").String(),
		"username":    cfg.Section("nacos").Key("username").String(),
		"password":    cfg.Section("nacos").Key("password").String(),
	}
	return serverInfo, nil
}
