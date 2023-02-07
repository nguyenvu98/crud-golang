package albumcontroller

import (
	"net/http"

	"github.com/nguyenvu/go-restapi-gin/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {

	var albums []models.Album

	models.DB.Find(&albums)
	c.JSON(http.StatusOK, gin.H{"albums": albums})

}

func GetAlbumByID(c *gin.Context) {
	var album models.Album
	id := c.Param("id")

	if err := models.DB.First(&album, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Can't find album!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"album": album})
}

func CreateAlbum(c *gin.Context) {

	var album models.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&album)
	c.JSON(http.StatusOK, gin.H{"album": album})
}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	id := c.Param("id")

	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&album).Where("id = ?", id).Updates(&album).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't update album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Albums updated successfully!"})

}

func DeleteAlbum(c *gin.Context) {
	var album models.Album
	id := c.Param("id")

	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&album).Where("id = ?", id).Delete(&album).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot delete album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album has been deleted"})
}
