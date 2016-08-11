package main

import "github.com/kataras/iris"

func main() {
    api := iris.New()
    api.Get("/", hello)
    api.Get("/temp", temp)
    api.Listen(":8080")
}

func temp(ctx *iris.Context) {
	temp, err := ctx.URLParamInt("temp")
	if err != nil {
		ctx.JSON(200, iris.Map{"error":err})
	}else if temp == 0 {
		ctx.JSON(200, iris.Map{"error":"please send temperature"})	
	}else {
		if temp > 30 {
			ctx.JSON(200, iris.Map{"fan":"on","temperature":temp})	
		}else {
			ctx.JSON(200, iris.Map{"fan":"off","temperature":temp})	
		}
	}
	
}

func hello(ctx *iris.Context) {
	ctx.Write("Hello Everyone!")
}
