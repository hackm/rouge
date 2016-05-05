package main

import (
    "./routes"
    "github.com/labstack/echo/engine/standard"
)

// 初期化処理
func Init()  {

}

func main() {
    router := route.Init()
    router.Run(standard.New(":8080"))
}