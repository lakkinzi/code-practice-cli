package editors

import (
	"practice/cmd/config"
)

var constructors = map[config.Editor]func(string){
	config.NeoVim: openNvim,
	config.VSCode: openVsCode,
}

func Open(conf *config.Config, path string) {
	constructors[conf.Editor](path)
}
