package kgin

import (
	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
)

func Default() *gin.Engine {
	if oss.IsRelease() {
		logs.SetLevel(logs.Info)
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
	} else {
		logs.SetLevel(logs.Debug)
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()
	// 设置真实IP的header
	// engine.TrustedPlatform = "Client-IP"
	// 设置安全的代理
	// engine.SetTrustedProxies([]string{})
	engine.Use(LogHandler, gin.Recovery())

	return engine
}
