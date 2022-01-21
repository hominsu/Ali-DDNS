package service

import (
	"Ali-DDNS/app/server/service/internal/biz"
	"Ali-DDNS/pkg"
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type Device struct {
	UUID string `json:"uuid"`
}

type Domain struct {
	DomainName string `json:"domain_name"`
}

type User struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func setSession(context *gin.Context, key interface{}, val interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, val)
	return session.Save()
}

func getSession(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}

func delSession(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	session.Delete(key)
	return session.Save()
}

// SessionMiddleWare session middleware, ensure that handlers after are checked by the middleware to see if they are logged in
func (s *DomainTaskService) SessionMiddleWare(c *gin.Context) {
	sess := getSession(c, "userKey")
	if sess == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}
}

// RegisterPost register a user
func (s *DomainTaskService) RegisterPost(c *gin.Context) {
	// get the username and password from http request header
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	username := user.UserName
	password := user.PassWord

	// check whether the username already exist
	exists, err := s.domainUserUsecase.IsUserExists(context.TODO(), &biz.DomainUser{Username: username})
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if exists {
		c.String(http.StatusBadRequest, "User is registered")
		return
	} else {
		// add the user to data repo
		_, err := s.domainUserUsecase.AddUser(context.TODO(), &biz.DomainUser{
			Username: username,
			Password: password,
		})
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// returns the message and redirects to the login page
		c.JSON(http.StatusOK, gin.H{
			"Register": true,
		})
		c.Redirect(http.StatusFound, "/login")
		return
	}
}

// RegisterDelete delete a user
func (s *DomainTaskService) RegisterDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
	return
}

// ------------------------- login、logout api -------------------------

// LoginGet login interface
func (s *DomainTaskService) LoginGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"LoginGet": "OK",
	})
}

// LoginPost post the username and password to login in
func (s *DomainTaskService) LoginPost(c *gin.Context) {
	// get the username and password from http request header
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	username := user.UserName
	password := user.PassWord

	// check whether the username is null, return if null
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.String(http.StatusBadRequest, "username or password should not be empty")
		return
	}

	// check whether the user already exist, return if not
	userExists, err := s.domainUserUsecase.IsUserExists(context.TODO(), &biz.DomainUser{Username: username})
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if !userExists {
		c.String(http.StatusUnauthorized, "User is not registered")
		return
	}

	// check if the user and password matches, return if it does not exist
	userPassword, err := s.domainUserUsecase.GetUserPassword(context.TODO(), &biz.DomainUser{Username: username})
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if password != userPassword {
		c.String(http.StatusUnauthorized, "Authorized failed")
		return
	}

	// add into session
	if err = setSession(c, "userKey", username); err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Login": true,
	})
}

// Logout logout interface
func (s *DomainTaskService) Logout(c *gin.Context) {
	// delete session
	if err := delSession(c, "userKey"); err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

// ------------------------- homepage -------------------------

// Home homepage
func (s *DomainTaskService) Home(c *gin.Context) {
	userName := c.Param("user_name")
	sessName := getSession(c, "userKey").(string)

	// check whether the username in req header is same as that in session repo
	if userName == sessName {
		c.JSON(http.StatusOK, gin.H{
			"status":    true,
			"user_name": userName,
		})
	} else {
		c.Redirect(http.StatusFound, "/home/"+sessName)
	}
}

// ------------------------- domain api -------------------------

// DomainNameGet obtain the domain name of the current user
func (s *DomainTaskService) DomainNameGet(c *gin.Context) {
	userName := c.Param("user_name")
	sessName := getSession(c, "userKey").(string)
	if userName != sessName {
		c.Redirect(http.StatusFound, "/DomainName/"+sessName)
		return
	}

	// obtain the domain name of the current user
	domainNames, err := s.domainUserUsecase.GetDomainName(context.TODO(), &biz.DomainUser{Username: userName})
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"domain_names": domainNames,
	})
}

// DomainNamePost is the interface for domain name post
func (s *DomainTaskService) DomainNamePost(c *gin.Context) {
	userName := c.Param("user_name")
	var domain Domain

	// get the domain info from req header
	if err := c.BindJSON(&domain); err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// add the domain name into date repo
	if _, err := s.domainUserUsecase.AddDomainName(context.TODO(), &biz.DomainUser{
		Username:   userName,
		DomainName: domain.DomainName,
	}); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

// DomainNameDel is interface for delete domain name
func (s *DomainTaskService) DomainNameDel(c *gin.Context) {
	userName := c.Param("user_name")
	var domain Domain

	// get the domain info from req header
	if err := c.BindJSON(&domain); err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// delete the domain name from data repo
	if _, err := s.domainUserUsecase.DelDomainName(context.TODO(), &biz.DomainUser{
		Username:   userName,
		DomainName: domain.DomainName,
	}); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

// ------------------------- device api -------------------------

// DeviceGet obtain the devices of the current user
func (s *DomainTaskService) DeviceGet(c *gin.Context) {
	userName := c.Param("user_name")
	sessName := getSession(c, "userKey").(string)
	if userName != sessName {
		c.Redirect(http.StatusFound, "/Device/"+sessName)
		return
	}

	// obtain all devices of the current user from data repo
	devices, err := s.domainUserUsecase.GetDevice(context.TODO(), &biz.DomainUser{Username: userName})
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"devices": devices,
	})
}

// DevicePost is the interface for add new device
func (s *DomainTaskService) DevicePost(c *gin.Context) {
	userName := c.Param("user_name")

	// generate a uuid
	uuid, err := pkg.NewUUID()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// save the uuid into data repo
	if _, err := s.domainUserUsecase.AddDevice(context.TODO(), &biz.DomainUser{
		Username: userName,
		UUID:     uuid,
	}); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

// DeviceDel is the interface for delete device
func (s *DomainTaskService) DeviceDel(c *gin.Context) {
	userName := c.Param("user_name")
	var device Device

	// get the device uuid from req header
	if err := c.BindJSON(&device); err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// delete this device from data repo
	if _, err := s.domainUserUsecase.DelDevice(context.TODO(), &biz.DomainUser{
		Username: userName,
		UUID:     device.UUID,
	}); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
