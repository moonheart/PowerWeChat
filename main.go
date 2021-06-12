package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/work"
	"os"
)

func GetConfig() *object.HashMap {
	return &object.HashMap{
		"corp_id": os.Getenv("corp_id"),
		"secret":  os.Getenv("secret"),
	}
}

func main() {

	fmt.Printf("hello Wechat! \n")

	config := GetConfig()

	app := work.NewWork(config)
	//fmt2.Dump(app)
	//fmt2.Dump(app.GetConfig())

	//token := app.AccessToken.GetToken()
	//fmt2.Dump(token)

	//cType := reflect.TypeOf((*app.Components)["base"].(*base.Client))
	//fmt.Printf("kind %s \n", cType.Kind())
	//fmt.Printf("type %v \n", cType)

	//ips := app.Base.GetCallbackIp()
	//fmt2.Dump(ips)
	//domainIps := app.Base.GetAPIDomainIP()
	//fmt2.Dump(domainIps)

	fmt2.Dump(app.OAuth.Config)



}
