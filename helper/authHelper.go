package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// CheckUserType function
func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("you are not authorized to access this route")
		return err;
	}
	return err
}

// MathUserTypeToUid function
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("you are not authorized to access this route")
		return err;
 	}

	err = CheckUserType(c, userType);
	return err;
}