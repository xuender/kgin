package kgin

import (
	"io/fs"
	"net/http"
	"path/filepath"
)

func FileSystem(fsys fs.FS, dirs ...string) http.FileSystem {
	if len(dirs) > 0 {
		fsys, _ = fs.Sub(fsys, filepath.Join(dirs...))
	}

	return http.FS(fsys)
}

// func Root(fsys fs.FS, dirs ...string) gin.HandlerFunc {
// 	eng.Use(func(ctx *gin.Context) {
// 		if ctx.Request.Method == http.MethodGet {
// 			if url := ctx.Request.URL.String(); url != "" && url != "/" {
// 				ctx.Header(CacheControl, MaxAge1y)
// 			}
// 		}
// 	})

// 	if len(dirs) > 0 {
// 		fsys, _ = fs.Sub(fsys, filepath.Join(dirs...))
// 	}

// 	eng.StaticFS("", http.FS(fsys))
// }
