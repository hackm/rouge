package main

import (
    "./route"
)

// 初期化処理
func Init()  {

}

func main() {
    router := route.Init()
    router.Run(":8080")
}