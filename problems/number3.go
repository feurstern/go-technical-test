package problems

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	router := gin.Default()

	// fmt.Printf("API_URL: %s", os.Getenv("API_URL"))
	router.GET("fetch-data", func(ctx *gin.Context) {
		res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

		if err != nil {
			ctx.JSON(500, gin.H{"message": "erorrr to read the response body "})
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			ctx.JSON(500, gin.H{"Message": "error to parse jso"})
			return
		}

		var data []Data

		if err := json.Unmarshal(body, &data); err != nil {
			ctx.JSON(500, gin.H{"message": "failed to parse the data"})
		}

		ctx.JSON(200, data)

	})

	router.Run(":8181")
}
