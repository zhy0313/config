package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//New new a Config
func New(file string) Config {
	return Config{file: file}
}

type Config struct {
	file string
	maps map[string]interface{}
}

//Get name pattern key or key.key.key
//
func (c *Config) Get(name string) interface{} {
	if c.maps == nil {
		c.read()
	}

	// app.view.path
	keys := strings.Split(name, ".")

	if len(keys) == 1 {
		return c.maps[name]
	}

	var ret interface{}
	for i := 0; i < len(keys); i++ {
		if ret == nil {
			ret = c.maps[keys[i]]
		} else {
			ret = ret.(map[string]interface{})[keys[i]]
		}
	}
	return ret
}

func (c *Config) read() {
	if !filepath.IsAbs(c.file) {
		file, err := filepath.Abs(c.file)
		if err != nil {
			panic(err)
		}
		c.file = file
	}

	bts, err := ioutil.ReadFile(c.file)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bts, &c.maps)

	if err != nil {
		panic(err)
	}
}
