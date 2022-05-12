// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wire-example/internal/config"
	"wire-example/internal/db"
)

// Injectors from wire.go:

// 调用wire.Build方法传入所有的依赖对象以及构建最终对象的函数得到目标对象
func InitApp() (*App, error) {
	string2 := _wireStringValue
	configConfig, err := config.New(string2)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.NewDb(configConfig)
	if err != nil {
		return nil, err
	}
	app := NewApp(sqlDB)
	return app, nil
}

var (
	_wireStringValue = "wo来了"
)
