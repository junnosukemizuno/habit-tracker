package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Habit struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func CreateHabit(c *gin.Context) {
	var habit Habit

	// JSONバインド（バリデーションも自動）
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ※今回はDB保存はせず、JSONを返すだけ
	c.JSON(http.StatusOK, gin.H{
		"message": "習慣を登録しました！",
		"habit":   habit,
	})
}
