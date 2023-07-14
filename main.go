package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	initLogs()

	// 加载配置文件
	cfg, _ := loadAPPConfig()

	// 接待所有的文件并监听变化
	for _, section := range cfg.SectionStrings()[1:] {
		projectDir := getCurrentAbPathByCaller()
		var configDir string = filepath.Join(projectDir, "config")

		groupIni := cfg.Section(section).Key("group").String()
		dataidIni := cfg.Section(section).Key("dataid").String()
		configPath := cfg.Section(section).Key("configpath").String()

		dataIDs := strings.Split(dataidIni, ",")
		group := groupIni

		appConfigDir := ""
		if configPath == "" {
			appConfigDir = filepath.Join(configDir, group)
			//os.Mkdir(appConfigDir, 0777)
		} else {
			appConfigDir = configPath
		}

		os.Mkdir(appConfigDir, 0777)

		if group != "common" {
			commonDir := filepath.Join(configDir, "common")
			fileslist, _ := listFilesInDirectory(commonDir)
			for _, file := range fileslist {
				linkFromfile := filepath.Join(commonDir, file)
				linkFromTo := filepath.Join(appConfigDir, file)
				os.Symlink(linkFromfile, linkFromTo)
			}

		}

		err := runNacosConfig(appConfigDir, group, dataIDs)
		if err != nil {
			log.Println(err.Error())
		}

	}
	select {}
}
