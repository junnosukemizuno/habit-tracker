package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mjunn/habit-tracker/handlers"
)

func main() {
	r := gin.Default()

	// エンドポイント登録
	r.POST("/habits", handlers.CreateHabit)

	r.Run(":8080") // localhost:8080 で起動
}
