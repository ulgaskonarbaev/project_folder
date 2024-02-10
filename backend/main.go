package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация маршрутизатора Gin
	router := gin.Default()

	// Обработчики запросов
	router.GET("/api/users", GetUsers)
	router.GET("/api/users/:id", GetUser)
	router.POST("/api/users", CreateUser)
	router.PUT("/api/users/:id", UpdateUser)
	router.DELETE("/api/users/:id", DeleteUser)

	// Запуск сервера на порту 8000
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

// Пример обработчиков запросов
func GetUsers(c *gin.Context) {
	// Логика получения списка пользователей
}

func GetUser(c *gin.Context) {
	// Логика получения конкретного пользователя
}

func CreateUser(c *gin.Context) {
	// Логика создания нового пользователя
}

func UpdateUser(c *gin.Context) {
	// Логика обновления информации о пользователе
}

func DeleteUser(c *gin.Context) {
	// Логика удаления пользователя
}