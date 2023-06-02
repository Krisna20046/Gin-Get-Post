package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Definisikan struktur data untuk sebuah postingan
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Daftar postingan yang sudah ada
var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// SetupRouter mengatur rute-rute HTTP dan mengembalikan mesin Gin yang telah dikonfigurasi
func SetupRouter() *gin.Engine {
	r := gin.Default() 

	// Rute untuk mendapatkan semua postingan
	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"posts": Posts}) 
	})

	// Rute untuk mendapatkan postingan berdasarkan ID
	r.GET("/posts/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		for _, post := range Posts {
			if post.ID == id {
				c.JSON(http.StatusOK, gin.H{"post": post})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
	})

	// Rute untuk menambahkan postingan baru
	r.POST("/posts", func(c *gin.Context) {
		var newPost Post
		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		newPost.ID = len(Posts) + 1
		newPost.CreatedAt = time.Now()
		newPost.UpdatedAt = time.Now()

		Posts = append(Posts, newPost)

		c.JSON(http.StatusCreated, gin.H{"message": "Postingan berhasil ditambahkan", "post": newPost})
	})

	return r
}

func main() {
	r := SetupRouter()

	// Jalankan server HTTP pada port 8080
	r.Run(":8080")
}
