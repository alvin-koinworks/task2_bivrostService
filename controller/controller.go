package controller

import (
	"bivrost_task2/database"
	"bivrost_task2/models"
	"log"
	"net/http"

	bs "github.com/koinworks/asgard-bivrost/service"
)

func CreateOrder(ctx *bs.Context) bs.Result {
	log.Println("Connect")

	db := database.GetDB()
	order := models.Orders{}

	err := ctx.BodyJSONBind(&order)

	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bs.ResponseBody{
			Message: map[string]string{
				"error":   "Bad Request",
				"message": err.Error(),
			},
		})
	}

	errPost := db.Debug().Create(&order).Error

	log.Println("satu", &order)

	if errPost != nil {

		log.Println("Error ", errPost.Error())
		return ctx.JSONResponse(http.StatusBadRequest, bs.ResponseBody{
			Message: map[string]string{
				"error":   "Bad Request",
				"message": errPost.Error(),
			},
		})
	}

	log.Println("Finish")

	return ctx.JSONResponse(http.StatusOK, bs.ResponseBody{
		Message: map[string]string{
			"status":  "Success",
			"message": "Insert item success",
		},
		Data: order,
	})

}

func CreateItem(ctx *bs.Context) bs.Result {

	db := database.GetDB()
	item := models.Item{}

	err := ctx.BodyJSONBind(&item)

	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bs.ResponseBody{
			Message: map[string]string{
				"error":   "Bad Request",
				"message": err.Error(),
			},
		})
	}

	errPost := db.Debug().Create(&item).Error

	if errPost != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bs.ResponseBody{
			Message: map[string]string{
				"error":   "Bad Request",
				"message": err.Error(),
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bs.ResponseBody{
		Data: item,
	})
}

func GetItems(ctx *bs.Context) bs.Result {
	db := database.GetDB()
	item := models.Item{}

	err := db.Find(&item).Error

	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bs.ResponseBody{
			Message: map[string]string{
				"error":   "Bad Request",
				"message": err.Error(),
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bs.ResponseBody{
		Message: map[string]string{
			"status":  "Success",
			"message": "Insert item success",
		},
		Data: item,
	})
}

func GetOrders(ctx *bs.Context) bs.Result {
	db := database.GetDB()
	order := models.Orders{}

	err := db.Find(&order).Error

	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bs.ResponseBody{
			Message: map[string]string{
				"error":   "Bad Request",
				"message": err.Error(),
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bs.ResponseBody{
		Data: order,
	})
}
