package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"xProcessBackend/model"
)

type Configer struct {
	Database         Database `json:"database"`
	Server           string   `json:"server"`
	FilePath         string   `json:"file_path"`
	ExportPath       string   `json:"export_path"`
	UsvName          string   `json:"usv_name"`
	ClientName       string   `json:"client_name"`
	ActivityAliasNum int      `json:"activity_alias_num"`
}
type Database struct {
	MysqlDsn string `json:"mysql_dsn"`
	redisUrl string `json:"redis_url"`
}

// Init 初始化配置项,返回端口号
func Init() {
	conf := GetConfig()
	// 连接数据库
	model.Database(conf.Database.MysqlDsn)
	model.RedisPool(conf.Database.redisUrl, "")

	// 返回配置文件中服务端口号

}

func GetConfig() Configer {
	var config Configer
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Get Config Error :", err)
	}
	return config
}
