package container

import (
	"Goshop/global/my_errors"
	"Goshop/global/variable"
	"strings"
	"sync"
)

var sMap sync.Map

func CreateContainersFactory() *containers {
	return &containers{}
}

type containers struct {
}

// 注册到容器
func (c *containers) Set(key string, value interface{}) (res bool) {
	if _, exists := c.KeyIsExists(key); exists == false {
		sMap.Store(key, value)
		res = true
	} else {
		variable.ZapLog.Warn(my_errors.ErrorsContainerKeyAlreadyExists + ", 相关键：" + key)
	}
	return
}

//  删除键值对
func (c *containers) Delete(key string) {
	sMap.Delete(key)
}

//  传递键，从容器获取值
func (c *containers) Get(key string) interface{} {
	if value, exists := c.KeyIsExists(key); exists {
		return value
	}
	return nil
}

//  判断键是否被注册
func (c *containers) KeyIsExists(key string) (interface{}, bool) {
	return sMap.Load(key)
}

// 按照键的前缀模糊删除容器中注册的内容
func (c *containers) FuzzyDelete(keyPre string) {
	sMap.Range(func(key, value interface{}) bool {
		if keyname, ok := key.(string); ok {
			if strings.HasPrefix(keyname, keyPre) {
				sMap.Delete(keyname)
			}
		}
		return true
	})
}
