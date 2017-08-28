package utils

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// I18n -
var I18n interface {
	Translate(string) string
}

type i18n struct {
}

func (i i18n) Translate(key string) string {
	buf, err := ioutil.ReadFile("messages/ja.yaml")
	if err != nil {
		return "translation missing"
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	log.Printf("%s", m["a"])
	return ""
}
