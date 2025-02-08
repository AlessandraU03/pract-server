package server_principal

import "github.com/gin-gonic/gin"

func StartServer1() {
	r := gin.Default()

	r.GET("/usuarios", getUsers)
	r.POST("/usuarios", createUser)
	r.DELETE("/usuarios/:id", deleteUser)  
	r.PUT("/usuarios/:id", updateUser)     

	r.GET("/longpoll", longPolling)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Servidor 1 activo"})
	})

	r.Run(":8080")
}