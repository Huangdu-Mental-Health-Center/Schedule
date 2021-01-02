package driver

import (
	"Huangdu_HMC_Schedule/src/logger"
	"encoding/json"
	"io/ioutil"
)

func LoadFile(nameFile string, v interface{}) bool {
	dir := ASSESR_DIR + nameFile + ".json"
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		logger.Error.Printf("Open File %s Failed\n", dir)
		logger.Error.Println(err)
		return false
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		logger.Error.Printf("Unmarshal %s Failed\n", dir)
		logger.Error.Println(err)
		return false
	}
	return true
}
