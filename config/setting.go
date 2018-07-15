package config

import (
	"io/ioutil"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"bytes"
	"os"
	"strconv"
)

var envirnoment=map[string]string{
	"tests":         "config/config.json",
}
func SetConfig(config map[string]string){
	envirnoment=config
}
var bInited = false
var settings map[string]interface{}
var env = "tests"

type ConfigValue struct {
	Value interface{}
}
func Init()error{
	if bInited {
		return nil
	}
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "tests"
	}
	fmt.Println("环境[" + env + "]")
	var configMap map[string]interface{}
	err := LoadSettingByLocalEnv(env, &configMap)

	if err != nil {
		return err
	}
	settings = configMap
	bInited = true
	return nil
}
func (self *ConfigValue) ToString() string {

	switch v := self.Value.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	}

	return fmt.Sprintf("%s", self.Value)
}

func (self *ConfigValue) ToInt() int {
	switch v := self.Value.(type) {
	case int:
		return v
	case string:
		k, _ := strconv.Atoi(v)
		return k
	case float32:

		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		fmt.Println(v)
		//util.CheckErr(errors.New("不能转换为int类型111"))

		return self.Value.(int)

	}

	return 0
}

func (self *ConfigValue) ToFloat() float64 {
	switch v := self.Value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	case int:

		return float64(v)
	case string:
		f, _ := strconv.ParseFloat(v, 20)

		return float64(f)
	}



	return 0
}


func LoadSettingByLocalEnv(env string,resultMap *map[string]interface{})error{
	return LocalSettingsByFile(envirnoment[env],resultMap)
}

func LocalSettingsByFile(file string,resultMap *map[string]interface{})error{
	content,err:=ioutil.ReadFile(file)
	if err!=nil {
		fmt.Println("Error",err)
		if err != nil {
			panic(err)
		}
	}
	mdz := json.NewDecoder(bytes.NewBuffer([]byte(content)))

	mdz.UseNumber()
	jsonErr := mdz.Decode(resultMap)

	return jsonErr
}
func GetValue(key string) *ConfigValue {

	if settings[key] == nil {
		valueStr := os.Getenv(key)
		if valueStr != "" {
			value := &ConfigValue{valueStr}
			return value
		}
	} else {
		value := &ConfigValue{settings[key]}
		return value
	}
	value := &ConfigValue{""}
	return value
}