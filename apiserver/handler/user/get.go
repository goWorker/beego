package user

import (
	"apiserver/model"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	. "apiserver/handler"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
