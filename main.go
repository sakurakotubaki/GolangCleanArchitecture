package main

import (
	"2/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	return db, err
}

func main() {
	db, err := initDB()
	// エラーならプロセスを終了する
	if err != nil {
		log.Fatal(err)
	}
	// dbをlogに出力する
	fmt.Println(db)

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