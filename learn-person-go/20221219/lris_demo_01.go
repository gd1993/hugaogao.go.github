package main

import (
	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello world!</h1>")
	})

	//app.Run(iris.Addr(":8900"), iris.WithTunneling)

	app.Run(iris.Addr(":8000"), iris.WithConfiguration(
		iris.Configuration{
			Tunneling: iris.TunnelingConfiguration{
				AuthToken:    "my-ngrok-auth-client-token",
				Bin:          "/bin/path/for/ngrok",
				Region:       "eu",
				WebInterface: "127.0.0.1:4040",
				Tunnels: []iris.Tunnel{
					{
						Name: "MyApp",
						Addr: ":8080",
					},
				},
			},
		}))
}
