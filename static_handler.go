package kgin

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/set"
)

// StaticHandler fs.
func StaticHandler(fsys fs.FS, dirs ...string) gin.HandlerFunc {
	if len(dirs) > 0 {
		fsys, _ = fs.Sub(fsys, filepath.Join(dirs...))
	}

	var (
		handler = http.FileServer(http.FS(fsys))
		paths   = set.NewSet[string]()
	)

	return func(ctx *gin.Context) {
		if url := strings.Trim(ctx.Request.URL.Path, "/"); fsHas(url, paths, fsys) {
			if url != "" {
				ctx.Header(CacheControl, MaxAge1y)
			}

			handler.ServeHTTP(ctx.Writer, ctx.Request)
			ctx.Abort()
		}
	}
}

func fsHas(path string, paths set.Set[string], fsys fs.FS) bool {
	if path == "" || paths.Has(path) {
		return true
	}

	if file, err := fsys.Open(path); err == nil {
		file.Close()
		paths.Add(path)

		return true
	}

	return false
}
