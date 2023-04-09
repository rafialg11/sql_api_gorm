package controllers

import (
	"net/http"
	"sql-api-gorm/database"
	"sql-api-gorm/models"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	db := database.GetDB()
	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data semua buku"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()
	var books models.Book
	if err := ctx.ShouldBindJSON(&books); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&books).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat data baru"})
		return
	}

	ctx.JSON(http.StatusCreated, books)
}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDB()
	books := models.Book{}

	if err := db.First(&books, ctx.Param("id")).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal Mendapatkan Data"})
		return
	}

	if err := ctx.ShouldBindJSON(&books); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&books).Where("id = ?", ctx.Param("id")).Updates(models.Book{Title: books.Title, Author: books.Author}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "gagal mengupdate data"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func GetBookByID(ctx *gin.Context) {
	db := database.GetDB()

	var books models.Book
	if err := db.First(&books, ctx.Param("id")).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal Mendapatkan Data"})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func DeleteBookById(ctx *gin.Context) {
	db := database.GetDB()
	var books models.Book
	if err := db.Where("id = ?", ctx.Param("id")).Delete(&books).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal Menghapus Data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil menghapus data !!!"})
}
