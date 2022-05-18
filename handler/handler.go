package handler

import (
	"fmt"
	"net/http"
	"project/todomvc/main.go/db"
	"project/todomvc/main.go/model"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var a model.TodomvcAdd
	c.ShouldBindJSON(&a)
	db.DB.Create(&model.Todomvc{Item: a.Item})
	c.JSON(http.StatusOK, nil)
}
func Delete(c *gin.Context) {
	var d model.TodomvcDel
	c.ShouldBindJSON(&d)
	db.DB.Where("id = ?", d.Id).Delete(&model.Todomvc{})
	c.JSON(http.StatusOK, nil)
}
func Update(c *gin.Context) {
	var u []model.TodomvcUpdate
	c.ShouldBindJSON(&u)

	for err, t := range u {
		fmt.Println(err)
		db.DB.Model(&model.Todomvc{}).Where("id = ?", t.Id).Select("Item", "Status").Updates(map[string]interface{}{
			"Item":   t.Item,
			"Status": t.Status,
		})

	}
	c.JSON(http.StatusOK, nil)
}
func Find(c *gin.Context) {
	var f model.TodomvcFind
	c.ShouldBindJSON(&f)

	var t []model.Todomvc

	if f.Item == "" {
		db.DB.Model(&model.Todomvc{}).Where("Status = ?", f.Status).Find(&t)
	}
	if f.Item != "" {
		db.DB.Model(&model.Todomvc{}).Where("Item LIKE ?", "%"+f.Item+"%").Find(&t)
	}
	c.JSON(http.StatusOK, &t)

}
