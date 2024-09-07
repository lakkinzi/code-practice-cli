package runners

import (
	"practice/cmd/config"
)

var runners = map[config.ProgLang]func(string){
	config.Go: runGo,
}

func Run(conf *config.Config, path string) {
	runners[conf.ProgLang](path)
}
