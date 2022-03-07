package server

func Init() {
	router := initRouter()

	router.Run("localhost:8080")
}
