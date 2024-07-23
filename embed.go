package kgin

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Embed(group *gin.RouterGroup, fsys embed.FS, dir string) error {
	files, err := fs.Sub(fsys, dir)
	if err != nil {
		return err
	}

	sys := http.FileServer(http.FS(files))
	hander := func(ctx *gin.Context) {
		if url := strings.Trim(ctx.Request.URL.Path, "/"); url != "" {
			ctx.Header(CacheControl, MaxAge1y)
		}

		sys.ServeHTTP(ctx.Writer, ctx.Request)
	}

	group.GET("", hander)
	group.HEAD("", hander)

	var register func(string) error
	register = func(sub string) error {
		sub = sanitize(sub)

		items, err := fsys.ReadDir(sub)
		if err != nil {
			return err
		}

		for _, item := range items {
			if item.IsDir() {
				if err := register(filepath.Join(sub, item.Name())); err != nil {
					return err
				}

				continue
			}

			url := filepath.Join(sub, item.Name())[len(dir):]

			group.GET(url, hander)
			group.HEAD(url, hander)
		}

		return nil
	}

	return register(dir)
}

func sanitize(embedPath string) string {
	return strings.ReplaceAll(embedPath, "\\", "/")
}
