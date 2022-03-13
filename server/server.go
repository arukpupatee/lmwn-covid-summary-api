package server

func Init() {
	router := InitRouter()

	router.Run("localhost:8080")
}
