package main

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
)

func runNacosConfig(configDir, group string, dataIDs []string) error {
	initLogs()
	serverInfo, _ := loadNacosConfig()
	serverPort, _ := strconv.ParseUint(serverInfo["port"], 10, 64)
	sc := []constant.ServerConfig{
		{
			IpAddr: serverInfo["ipAddr"],
			Port:   serverPort,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         serverInfo["namespaceId"],
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "warn",
		Username:            serverInfo["username"],
		Password:            serverInfo["password"],
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		return err
	}

	for _, dataID := range dataIDs {
		content, err := configClient.GetConfig(vo.ConfigParam{
			DataId: dataID,
			Group:  group,
		})
		if err != nil {
			log.Printf("获取配置 %s 失败：%s\n", dataID, err.Error())
			continue
		}

		filename := filepath.Join(configDir, dataID)
		err = ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			log.Printf("写入配置文件 %s 失败：%s\n", filename, err.Error())
			continue
		}

		log.Printf("成功写入配置文件：%s\n", filename)

		func() {
			err := configClient.ListenConfig(vo.ConfigParam{
				DataId: dataID,
				Group:  group,
				OnChange: func(namespace, group, dataID, data string) {
					//fmt.Println("配置文件发生了变化...")
					//fmt.Println("group:" + group + ", dataID:" + dataID + ", data:" + data)
					log.Printf("配置文件有修改，成功重新写入配置文件：%s\n", filename)
					err = ioutil.WriteFile(filename, []byte(data), 0644)
					if err != nil {
						log.Printf("写入配置文件 %s 失败：%s\n", filename, err.Error())
					}
				},
			})
			if err != nil {
				log.Println("监听配置变化失败:", err.Error())
			}
		}()
	}

	return nil
}
