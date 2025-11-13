package gui

import (
	"embed"
	"github.com/lincaiyong/gui/js"
	"github.com/lincaiyong/gui/utils"
	"github.com/lincaiyong/log"
	"io/fs"
	"path/filepath"
)

//go:embed com/**/*.js
var allJs embed.FS

func init() {
	err := fs.WalkDir(allJs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		b, err := allJs.ReadFile(path)
		if err != nil {
			return err
		}
		comName := filepath.Base(filepath.Dir(path))
		comName = utils.PascalCase(comName)
		js.Set(comName, string(b))
		return nil
	})
	if err != nil {
		log.FatalLog("fail to walk: %v", err)
	}
}
