package main

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"

	"go.uber.org/dig"

	"github.com/studtool/common/utils"

	"github.com/studtool/vk-auth-service/api"
	"github.com/studtool/vk-auth-service/beans"
)

func main() {
	c := dig.New()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	utils.AssertOk(c.Provide(api.NewServer))
	utils.AssertOk(c.Invoke(func(srv *api.Server) {
		if err := srv.Run(); err != nil {
			beans.Logger.Fatal(err)
		}
	}))
	defer func() {
		utils.AssertOk(c.Invoke(func(srv *api.Server) {
			if err := srv.Run(); err != nil {
				beans.Logger.Fatal(err)
			}
		}))
	}()

	<-ch

	_ = &oauth2.Config{
		Endpoint: vk.Endpoint,
	}
}
