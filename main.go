package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type User struct {
    Username string `json:"username"`
    Gender string `json:"gender"`
}

func getPingHandler(context *gin.Context) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "message": "pong",
        },
    )
}

func postUsersHandler(context *gin.Context) {
    var user User
    context.ShouldBind(&user)
    context.JSON(
        http.StatusOK,
        user,
    )
}

func main() {
    engine := gin.Default()
    engine.GET("/ping", getPingHandler)
    engine.POST("/users", postUsersHandler)
    engine.Run()
}
