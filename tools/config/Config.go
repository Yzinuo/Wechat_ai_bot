package config

import(
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Conf struct{
	APP App `json:"app" yaml:"app"`
	Keys Key `json : "keys" yaml : "keys`
}

type App struct{
	Env string `json:"env" yaml :"env"`
}

type Key struct{
	Bot_name string `json:"bot_name" yaml:"bot_name"`
	Deepseek_api string `json:"deepseek" yaml:"deepseek"`
	ApiUrl	string `json:"apiUrl"  yaml:"apiUrl"`
}

func GetConf(confpath string) (conf *Conf,err error){
	var(
		yamlFile = make([]byte,0)
	)

	logrus.Infof("filepath:%s",confpath)
	yamlFile,err = os.ReadFile(confpath)
	if err != nil{
		logrus.WithError(err).Error("An error occurred")
		return conf,err
	}

	err = yaml.Unmarshal(yamlFile,&conf)
	if err != nil{
		logrus.WithError(err).Error("yaml Unmarshal error")
		return conf,err
	}

	return conf,nil
}