package kgin

import (
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// StaticHandler fs.
func StaticHandler(urlPrefix string, fsys fs.FS, dirs ...string) gin.HandlerFunc {
	if len(dirs) > 0 {
		fsys, _ = fs.Sub(fsys, filepath.Join(dirs...))
	}

	handler := http.FileServer(http.FS(fsys))
	if urlPrefix != "" {
		handler = http.StripPrefix(urlPrefix, handler)
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path[1:]
		if _, err := fsys.Open(path); path == "" || err == nil {
			handler.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
