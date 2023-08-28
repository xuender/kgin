package kgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/set"
)

// DirHandler 目录读取.
func DirHandler(urlPrefix string, dir string) gin.HandlerFunc {
	var (
		dirFS   = http.Dir(dir)
		handler = http.FileServer(dirFS)
		length  = len(urlPrefix)
		paths   = set.NewSet[string]()
	)

	if length > 0 {
		handler = http.StripPrefix(urlPrefix, handler)
	}

	return func(ctx *gin.Context) {
		if httpHas(ctx.Request.URL.Path[length:], paths, dirFS) {
			handler.ServeHTTP(ctx.Writer, ctx.Request)
			ctx.Abort()
		}
	}
}

func httpHas(path string, paths set.Set[string], dirFS http.FileSystem) bool {
	if path == "" || paths.Has(path) {
		return true
	}

	if file, err := dirFS.Open(path); err == nil {
		file.Close()
		paths.Add(path)

		return true
	}

	return false
}
