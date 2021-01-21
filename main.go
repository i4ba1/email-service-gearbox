package main

import (
	"fmt"
	"github.com/gogearbox/gearbox"
)

func main() {
	//setup gearbox
	g := gearbox.New()

	//Define your handlers
	g.Group("/email", []*gearbox.Route{
		g.Post("/sendEmail", func(ctx gearbox.Context){
			fmt.Println(ctx.Body())
			ctx.SendString("Hello Cuy!!")
		}),

	})

	//Start service
	g.Start(":3445")
}
