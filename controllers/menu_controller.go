package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/wahyuharyo/prakerja/config"

	"github.com/wahyuharyo/prakerja/models"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var pesanan []models.Menu

	config.DB.Find(&pesanan)
	c.JSON(http.StatusOK, gin.H{"Orderan": pesanan})


}
func Show(c *gin.Context) {

	var pesanan models.Menu
	id := c.Param("id")

	if err := config.DB.First(&pesanan, id).Error; err != nil {
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"Orderan": pesanan})

}
func Create(c *gin.Context) { 

	var pesanan models.Menu

	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Pesanan tidak dapat dibuat"})
		return
	}

	config.DB.Create(&pesanan)
	c.JSON(http.StatusOK, gin.H{
		"message" : "Pesanan berhasil dibuat"})


}
func Update(c *gin.Context) {

	var pesanan models.Menu
	id := c.Param("id")

	if config.DB.Model(&pesanan).Where("id = ?", id).Updates(&pesanan).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Tidak dapat mengupdate pesanan"})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pesanan berhasil diupdate"})

}
func Delete(c *gin.Context) {

	var pesanan models.Menu

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if config.DB.Delete(&pesanan, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Tidak dapat menghapus pesanan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pesanan berhasil dihapus"})

}