package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando archivo .env")
		return
	}

	api_key := os.Getenv("API_KEY")
	if api_key == "" {
		fmt.Println("API_KEY no está configurada")
		return
	}
	testUrl := fmt.Sprintf("https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/Kipi/LAN?api_key=%s", api_key)

	resp, err := http.Get(testUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
