package utils

import (
	"bytes"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

func ResUtil(c *gin.Context, resp *http.Response) {

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	c.JSON(http.StatusOK, body)
}

func Struct2Map(obj interface{}) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		data[objT.Field(i).Name] = objV.Field(i).Interface()
	}
	err = nil
	return
}

// Struct2MapString 转换为string类型
func Struct2MapString(obj interface{}) (data map[string]string, err error) {
	data = make(map[string]string)
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		var val string
		switch objV.Field(i).Type().String() {
		case "int", "int64":
			val = strconv.FormatInt(objV.Field(i).Int(), 10)
		case "string":
			val = objV.Field(i).String()
		}
		tagStr := objT.Field(i).Tag.Get("json")
		if len(tagStr) > 0 {
			data[tagStr] = val
		} else {
			data[objT.Field(i).Name] = val
		}
	}
	err = nil
	return
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
