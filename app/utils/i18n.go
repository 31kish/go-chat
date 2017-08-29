package utils

import (
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type util_i18n interface {
	Translate(key string) string
}

var I18n util_i18n = i18n{}

type i18n struct{}

func (i18n) Translate(key string) string {
	messages := loadMessageFile()
	keys := strings.Split(key, ".")

	for i := 0; i < len(keys)-1; i++ {
		if val, ok := messages[keys[i]]; !ok || reflect.TypeOf(messages[keys[i]]).Kind() != reflect.Map {
			log.Panicf("translation missing key: %#v val: %#v", keys[i], val)
		}
		messages = messages[keys[i]].(map[interface{}]interface{})
	}

	if val, ok := messages[keys[len(keys)-1]]; !ok {
		log.Panicf("translation missing key: %#v val: %#v", keys[len(keys)-1], val)
	}

	return messages[keys[len(keys)-1]].(string)
}

func loadMessageFile() map[interface{}]interface{} {
	buf, err := ioutil.ReadFile("messages/ja.yaml")
	if err != nil {
		panic(err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	return m
}
