package Server

import (
	"github.com/polariseye/PolarServer/Server/WebServer"
)

var (
	WebServerObj *WebServer.WebServerStruct
)

func InitWebServer() {
	handle4UrlItem := WebServer.NewHandle4Url()
	WebServerObj = WebServer.NewWebServer(0, handle4UrlItem)
}
