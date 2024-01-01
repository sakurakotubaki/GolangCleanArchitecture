# Controller層を作成する
GETやPOSTするfunctionsを別ファイルに作成して、呼び出す時に使います。

controllerディレクトリを作成して、`task_controller.go`を作成します。
このファイルから、functionをmain.goで呼び出すには、`package controller`を定義する必要があります。

```go
package controller

import(
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct {}

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
}

func (t *TaskController) Get(c echo.Context) error {
	// tasks, err := usecase.GetTasks()
	return c.JSON(http.StatusOK, nil)
}

func (t *TaskController) Create(c echo.Context) error {
	var task Task
	err := c.Bind(&task)
	if err != nil {
			return err
	}
	// ここで task を使用して何かを行う
	return nil
}
```

`main.go`で`controller`を呼び出すには、`import`を追加します。

```go
package main

import (
	"2/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())// Loggerはリクエストとレスポンスの情報を出力する
	e.Use(middleware.Recover())// パニックが発生した場合に500エラーを返す

	// コントローラーを使うには初期化する必要がある
	taskController := controller.TaskController{}

	e.GET("/tasks", taskController.Get)
	e.POST("/tasks", taskController.Create)

	// サーバーを起動する
	e.Start(":8080")
}
```