package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/spf13/viper"
)

func Load() *Config {
	path, err := filepath.Abs("conf.d/app.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println("get config path: ", path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	conf := &Config{}
	setDefault(conf, "")
	if err := v.UnmarshalKey("config", conf); err != nil {
		panic(err)
	}

	return conf
}

// parse nested structure `config` and set to default value by tag `default`
func setDefault(config any, rootKey string) {
	keyStack := []string{rootKey}
	pointer := reflect.ValueOf(config)
	if pointer.Kind() != reflect.Pointer {
		return
	}
	pointer = pointer.Elem()
	if pointer.Kind() != reflect.Struct {
		return
	}
	stack := []reflect.Value{pointer}

	for len(stack) > 0 {
		pointer = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		key := keyStack[len(keyStack)-1]
		keyStack = keyStack[:len(keyStack)-1]

		for i := 0; i < pointer.NumField(); i++ {
			field := pointer.Field(i)
			switch field.Kind() {
			case reflect.Struct:
				stack = append(stack, field)
				keyStack = append(keyStack, fmt.Sprintf("%s.%s", key, field.Type().Name()))
				break
			default:
				tag := pointer.Type().Field(i).Tag.Get("default")
				if tag == "" {
					continue
				}
				field.Set(reflect.ValueOf(tag))
			}
		}
	}
}
