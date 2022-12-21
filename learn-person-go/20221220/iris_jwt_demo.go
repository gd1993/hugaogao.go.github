package main

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	j := jwt.New(jwt.Config{
		Extractor: jwt.FromParameter("token"),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	app.Get("/", getTokenHandler)
	app.Get("/secured", j.Serve, myAuthenticatedHandler)
	app.Run(iris.Addr(":8080"))
}

func myAuthenticatedHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")

	foobar := user.Claims.(jwt.MapClaims)
	for key, value := range foobar {
		ctx.Writef("%s = %s", key, value)
	}
}

func getTokenHandler(ctx iris.Context) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})
	tokenString, _ := token.SignedString([]byte("My Secret"))
	ctx.HTML(`Token: ` + tokenString + `<br/><br/>
    <a href="/secured?token=` + tokenString + `">/secured?token=` + tokenString + `</a>`)
}
