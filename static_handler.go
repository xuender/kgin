package kgin

import (
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/set"
)

// StaticHandler fs.
func StaticHandler(urlPrefix string, fsys fs.FS, dirs ...string) gin.HandlerFunc {
	if len(dirs) > 0 {
		fsys, _ = fs.Sub(fsys, filepath.Join(dirs...))
	}

	var (
		handler = http.FileServer(http.FS(fsys))
		length  = len(urlPrefix)
		paths   = set.NewSet[string]()
	)

	if length > 0 {
		handler = http.StripPrefix(urlPrefix, handler)
	}

	return func(c *gin.Context) {
		if fsHas(c.Request.URL.Path[length:], paths, fsys) {
			handler.ServeHTTP(c.Writer, c.Request)
			c.Abort()
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
