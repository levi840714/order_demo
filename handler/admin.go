package handler

import (
	"order_demo/lib/compute"
	"order_demo/lib/logger"
	"order_demo/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const GetAllOrder = 0

type GoodStatus struct {
	Status string `form:"status" binding:"required"`
}

// @Summary 查看餐點
// @Tags Goods
// @version 1.0
// @Accpet json
// @Produce json
// @Param status query string true "查看餐點("0"=>"未上架", "1"=>"已上架")"
// @Success 200 {string}} json "{"code": 0, "msg": "", "data": "{餐點}""}"
// @Router /goods [get]
func GetGoods(c *gin.Context) {
	var request GoodStatus
	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	goods, err := model.GetGoods(request.Status)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Get goods failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": goods})
}

type Add struct {
	Goods  string  `json:"goods" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

// @Summary 新增餐點
// @Tags Admin
// @version 1.0
// @Accpet json
// @Produce json
// @Security ApiToken
// @Param payload body Add true "餐點資料"
// @Success 200 {string} json "{"code": 0, "msg": "", "data": "id": "{餐點ID}"}"
// @Router /admin/goods [post]
func AddGoods(c *gin.Context) {
	var request Add
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	id, err := model.AddGoods(request.Goods, request.Amount)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Add Goods failed", "data": ""})
		return
	}
	c.JSON(201, gin.H{"code": 0, "msg": "", "data": map[string]int{"id": id}})
}

type UpdateId struct {
	ID     int     `json:"id" binding:"required"`
	Goods  string  `json:"goods"`
	Amount float64 `json:"amount"`
}

// @Summary 更新餐點
// @Tags Admin
// @version 1.0
// @Accpet json
// @Produce json
// @Security ApiToken
// @Param payload body UpdateId true "餐點資料"
// @Success 200 {string} json "{"code": 0, "msg": "Update goods success", "data": ""}"
// @Router /admin/goods/ [put]
func UpdateGoods(c *gin.Context) {
	var request UpdateId
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	if err := model.UpdateGoods(request.ID, request.Goods, request.Amount); err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Update goods failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Update goods success", "data": ""})
}

// @Summary 刪除餐點
// @Tags Admin
// @version 1.0
// @Accpet json
// @Produce json
// @Security ApiToken
// @Param id path string true "餐點ID"
// @Success 200 {string} json "{"code": 0, "msg": "Delete success", "data": ""}"
// @Router /admin/goods/{id} [delete]
func DeleteGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	if err := model.DeleteGoods(id); err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Delete goods failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Delete success", "data": ""})
}

// @Summary 今日訂餐總結算
// @Tags Admin
// @version 1.0
// @Accpet json
// @Produce json
// @Security ApiToken
// @Success 200 {string} json "{"code": 0, "msg": "", "data": "點餐總計"}"
// @Router /admin/summaryList [get]
func GetTodaySummary(c *gin.Context) {
	var count, total float64
	var orderList = make(map[string]map[string]float64)

	orders, err := model.GetTodayOrder(GetAllOrder)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Get today order failed", "data": ""})
		return
	}

	for _, v := range *orders {
		count++
		total += v.Amount
		if orderList[v.Account] == nil {
			orderList[v.Account] = make(map[string]float64)
		}
		orderList[v.Account][v.Goods] = compute.Add(orderList[v.Account][v.Goods], v.Amount)
	}

	if orderList["summary"] == nil {
		orderList["summary"] = make(map[string]float64)
	}
	orderList["summary"]["count"] = count
	orderList["summary"]["total"] = total

	c.JSON(200, gin.H{"code": 0, "msg": "", "data": orderList})
}
