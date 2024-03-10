package routes

import (
	"github.com/golangast/endrulats/src/handler/get/home"
	"github.com/golangast/endrulats/src/handler/get/loginemail"
	"github.com/golangast/endrulats/src/handler/get/profile"
	"github.com/golangast/endrulats/src/handler/post/createuser"
	"github.com/golangast/endrulats/src/handler/post/userinput"
	"github.com/golangast/endrulats/src/handler/restful/post"
	"github.com/labstack/echo/v4"
	// importroute
)

func Routes(e *echo.Echo) {
	//get
	e.GET("/", home.Home)
	e.GET("/usercreate", profile.Profile)
	e.GET("/loginemail/:email/:sitetoken", loginemail.LoginEmail)
	//getroute
	//post
	e.POST("/usercreate", createuser.Createuser)
	e.POST("/userinput", userinput.UserInput)
	e.POST("/p", post.Posts)
	//postroute
}
