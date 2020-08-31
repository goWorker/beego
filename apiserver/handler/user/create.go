package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c,errno.ErrBind,nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	//admin2 := c.Param("username")
	//log.Infof("URL username: %s",admin2)
	//
	//desc := c.Query("desc")
	//log.Infof("URL key param desc: %s",desc)
	//
	//contentType := c.GetHeader("Content-Type")
	//log.Infof("Header Content-Type: %s",contentType)
	//
	//log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateReponse{
		Username: r.Username,
	}
	SendResponse(c,nil,rsp)

}

func (r *CreateRequest) checkParam() error{
	if r.Username == ""{
		return errno.New(errno.ErrValidation,nil).Add("username is empty")
	}
	if r.Password == "" {
		return errno.New(errno.ErrValidation,nil).Add("password is empty.")
	}
	return nil
}