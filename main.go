package main

import (
	"net/http"

	"github.com/nguyenvu/go-restapi-gin/controllers/albumcontroller"

	"github.com/nguyenvu/go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/albums", albumcontroller.GetAlbums)
	r.GET("/api/albums/:id", albumcontroller.GetAlbumByID)
	r.POST("/api/albums", albumcontroller.CreateAlbum)
	r.PUT("/api/albums/:id", albumcontroller.UpdateAlbum)
	r.DELETE("/api/albums/:id", albumcontroller.DeleteAlbum)

	//----------------upload file----------------
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.POST("/", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Cannot upload image",
			})
			return
		}
		//save the file
		err = c.SaveUploadedFile(file, "assets/upload/"+file.Filename)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Cannot upload image",
			})
			return
		}
		//render page
		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": "/assets/upload/" + file.Filename,
		})
	})
	r.Run()

}
