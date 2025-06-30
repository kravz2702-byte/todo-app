package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kravz2702-byte/todo-app/pkg/entity"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

}

func (h *Handler) signIn(c *gin.Context) {

}
