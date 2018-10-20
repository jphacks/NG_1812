package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type Follower struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

func fetch(username string) ([]Follower, error) {
	url := "https://api.github.com/users/" + username + "/followers"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	var articles []Follower
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		c.Next()
	}
}

func KusaHandler(c *gin.Context) {
	name := c.Param("name")
	res, err := http.Get("https://github.com/users/"+name+"/contributions")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	html,err :=doc.Find("svg").Parent().Html()
	if err != nil {
		log.Fatal(err)
	}
	c.Writer.Header().Set("Content-Type", "image/svg+xml")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.String(200,html)
	return
}

func main() {
	router := gin.Default()
	//CORS対策
	router.Use(CORSMiddleware())

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		follower, err := fetch(name)
		if err != nil {
			//落ちると困るのでJSONでエラーを吐く
			//log.Fatalf("Error!: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}

		//構造体をJSONに戻す
		var buf bytes.Buffer
		b, _ := json.Marshal(follower)
		buf.Write(b)
		c.String(http.StatusOK, buf.String())
	})
	router.GET("/user/:name/kusa",KusaHandler)
	router.Run("localhost:8080")
}
