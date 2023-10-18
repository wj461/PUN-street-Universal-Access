package delivery

import (
	"strconv"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CartHandler struct {
	CartUsecase domain.CartUsecase
}

func NewCartHandler(e *gin.Engine, cartUsecase domain.CartUsecase) {
	handler := &CartHandler{
		CartUsecase: cartUsecase,
	}
	e.GET("/api/v1/cart/:customerID/:productID/:storeID", handler.GetCartById)
	e.POST("/api/v1/cart", handler.PostCart)
}

func (s *CartHandler) GetCartById(c *gin.Context) {
	customerID := c.Param("customerID")
	productID := c.Param("productID")
	storeID := c.Param("storeID")

	cart, err := s.CartUsecase.GetByID(c, customerID, productID, storeID)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :(",
		})
		return
	}

	c.JSON(200, &swagger.CartInfo{
		CustomerID:      cart.CustomerID,
		ProductID:       cart.ProductID,
		StoreID:         cart.StoreID,
		ProductQuantity: cart.ProductQuantity,
	})
}

func (s *CartHandler) PostCart(c *gin.Context) {
	var cart domain.Cart

	if err := c.BindJSON(&cart); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :( ",
		})
		return
	}

	cart.CustomerID = sub(cart.CustomerID, 1)
	cart.ProductID = sub(cart.ProductID, 1)
	cart.StoreID = sub(cart.StoreID, 1)

	_cart, err := s.CartUsecase.PostCart(c, &cart)

	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :( ",
		})
		return
	}

	c.JSON(200, &swagger.CartInfo{
		CustomerID:      _cart.CustomerID,
		ProductID:       _cart.ProductID,
		StoreID:         _cart.StoreID,
		ProductQuantity: _cart.ProductQuantity,
	})
}

func sub(s string, n int) string {
	i, err := strconv.Atoi(s)
	if err != nil {
		return s
	}
	return strconv.Itoa(i - n)
}
