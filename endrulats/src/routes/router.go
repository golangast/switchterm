package routes

import (
	"github.com/golangast/endrulats/src/handler/get/ch"
	"github.com/golangast/endrulats/src/handler/get/dogger"
	"github.com/golangast/endrulats/src/handler/get/home"
	"github.com/golangast/endrulats/src/handler/get/loginemail"
	"github.com/golangast/endrulats/src/handler/get/profile"
	"github.com/golangast/endrulats/src/handler/get/yet"
	"github.com/golangast/endrulats/src/handler/post/createuser"
	"github.com/golangast/endrulats/src/handler/post/userinput"
	"github.com/golangast/endrulats/src/handler/restful/post"
	"github.com/labstack/echo/v4"
	"github.com/golangast/endrulats/src/handler/get/gar"
"github.com/golangast/endrulats/src/handler/get/mo"
"github.com/golangast/endrulats/src/handler/get/yes"
"github.com/golangast/endrulats/src/handler/get/golang"
// importroute
)

func Routes(e *echo.Echo) {
	//get
	e.GET("/", home.Home)
	e.GET("/usercreate", profile.Profile)
	e.GET("/loginemail/:email/:sitetoken", loginemail.LoginEmail)
	e.GET("/ch", ch.Ch)
	e.GET("/dogger", dogger.Dogger)
	e.GET("/yet", yet.Yet)
	e.GET("/gar", gar.Gar)
e.GET("/mo", mo.Mo)
e.GET("/yes", yes.Yes)
e.GET("/golang", golang.Golang)
//getroute
	//post
	e.POST("/usercreate", createuser.Createuser)
	e.POST("/userinput", userinput.UserInput)
	e.POST("/p", post.Posts)
	//postroute
}
