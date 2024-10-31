package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/cocktails", getCocktails)
    router.GET("/cocktails/:letter", getCocktailsByLetter)

	router.Run(":8080")
}

func getCocktails(context *gin.Context) {
	// Call the external API
    response, err := http.Get("https://www.thecocktaildb.com/api/json/v1/1/search.php?f=a")
    if err != nil {
        log.Println("Error calling the API:", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
        return
	}
	defer response.Body.Close()

    // Read the response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Println("Error reading response body:", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data"})
        return
    }

    // Send the response to the client
    context.Data(http.StatusOK, "application/json", body)
}

func getCocktailsByLetter(context *gin.Context) {
    // Get the letter from te URL parameter
    letter := context.Param("letter")

    url := fmt.Sprintf("https://www.thecocktaildb.com/api/json/v1/1/search.php?f=%s", letter)

    // Call the external API
    response, err := http.Get(url)
    if err != nil {
        log.Println("Error calling the API:", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
        return
	}
	defer response.Body.Close()

    // Read the response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Println("Error reading response body:", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data"})
        return
    }

    // Send the response to the client
    context.Data(http.StatusOK, "application/json", body)
}
