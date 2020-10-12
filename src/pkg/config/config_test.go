package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c, _ := NewConfig("../../src/config/dev.yaml")
	fmt.Println(c)
	fmt.Println(c.GetString("mysql.a"))
	//fmt.Println(c.Get("mysql").Decode())
}
