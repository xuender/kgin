package kgin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/oss"
)

func Default() *gin.Engine {
	if oss.IsRelease() {
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		gin.SetMode(gin.TestMode)
	}

	engine := gin.New()
	engine.ContextWithFallback = true
	// 设置真实IP的header
	// engine.TrustedPlatform = "Client-IP"
	// 设置安全的代理
	// engine.SetTrustedProxies([]string{})
	engine.Use(LogHandler, gin.Recovery())

	return engine
}
