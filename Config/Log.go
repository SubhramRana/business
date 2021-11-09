package Config
import(
	log "github.com/sirupsen/logrus"
)
func LogConfig(){
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
