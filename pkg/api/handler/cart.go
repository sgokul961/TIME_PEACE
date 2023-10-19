package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	response "gokul.go/pkg/utils/response"
)

type CartHandler struct {
	usecase usecaseInterfaces.CartUseCase
}

func NewCartHandler(usecase usecaseInterfaces.CartUseCase) *CartHandler {
	return &CartHandler{
		usecase: usecase,
	}
}

func (i *CartHandler) AddToCart(c *gin.Context) {
	var model models.AddToCart

	if err := c.BindJSON(&model); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := i.usecase.AddToCart(model.UserID, model.InventoryID); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added to cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// func (i *CartHandler) CheckOut(c *gin.Context) {

// 	id, err := strconv.Atoi(c.Query("id"))
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "user_id not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	products, err := i.usecase.CheckOut(id)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}
// 	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
// 	c.JSON(http.StatusOK, successRes)
// }

//------mind checkout while doing ---------------commanding//