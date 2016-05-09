package main

import (
    "os"
    "os/exec"
    "runtime"
    "fmt"
    "flag"
    "text/template"
    "./routes"
    "./engine"
    "github.com/labstack/echo/engine/standard"
)

var (
    serve = flag.Bool("s", false, "serve")
    config *engine.Config
)

// 初期化処理
func init()  {
    config = engine.LoadConfig()
}

func main() {
    flag.Parse()
    if *serve {
        fmt.Println("start serve...")
        router := route.Init()
        router.Run(standard.New(":8080"))
    } else {
        generateRouter()
        startServe()
    }
}

func generateRouter() {
    fmt.Println("generate router.go")
    file, err := os.Create("routes/router.go")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    tmpl := template.Must(template.ParseFiles("routes/router.go.tmpl"))
    if err := tmpl.Execute(file, config.Plugins); err != nil {
        panic(err)
    }
}
func startServe() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("powershell", "go run ./main.go -s")
    } else {
        cmd = exec.Command("go", "run ./main.go -s")        
    }
    cmd.Stdout = os.Stdout
    if err := cmd.Start(); err != nil {
        panic(err)
    }
    cmd.Wait()
}