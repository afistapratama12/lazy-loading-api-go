package service

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

var (
	fake = faker.New()
)

type handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{db}
}

func (h *handler) GetAllData(c *gin.Context) {
	var products []Product

	if err := h.db.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "error internal server",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

func (h *handler) GetQuery(c *gin.Context) {
	var products []Product

	var sql = findAll

	if s := c.Query("s"); s != "" {
		sql = fmt.Sprintf(queryWhere, sql, s, s)
	}

	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf(querySort, sql, sort)
	}

	page, _ := strconv.Atoi(c.Query("page"))
	perPage := 9
	var total int64

	h.db.Raw(findAll).Count(&total)

	sql = fmt.Sprintf(queryPage, sql, perPage, (page-1)*perPage)

	h.db.Raw(sql).Scan(&products)

	c.JSON(200, gin.H{
		"data":      products,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total / int64(perPage))),
	})
}

func (h *handler) InsertData(c *gin.Context) {
	for i := 1; i <= 50; i++ {
		h.db.Create(&Product{
			Title:       fake.Lorem().Word(),
			Description: fake.Lorem().Paragraph(15),
			Image:       fmt.Sprintf("http://lorempixel.com/200/200?%s", fake.UUID().V4()),
			Price:       rand.Intn(90) + 10,
		})
	}

	c.JSON(201, gin.H{
		"message": "success",
	})
}
