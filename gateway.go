package main
import (
	"flag"
	"github.com/zhaoweikid/gobase"
)

type Config struct {
	ip		string
	port	int
	logfile string
}

func main() {
	var conffile = flag.String("config", "gateway.conf", "config file")
	cfp := gobase.NewConfigParser(*conffile)
	if cfp == nil {
		gobase.Warn("not found config file at: %s", *conffile)
		return
	}
	sec := cfp.Section
	log := gobase.NewLog(sec["default"]["logfile"].AsString(), gobase.DEBUG)	
	
	log.Warn("haha, go~~")

}

