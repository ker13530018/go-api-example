package user

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

// Model is struct of user
type Model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tel  string `json:"tel"`
}

// List of user
type List map[int]Model

// Context self gin.Context
type Context struct {
	*gin.Context
}

// Users is struct of users
type Users struct {
	All          List
	LastInsertID int
}

// GetAll get user all data
func (u *Users) GetAll(c *gin.Context) {
	context := Context{c}

	list, err := u.findAll()
	if err != nil {
		context.ErrorInternalServer(errors.Wrap(err, "find user error"))
		return
	}

	if len(*list) <= 0 {
		context.ErrorNotFound(errors.New("users is empty"))
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
		"data":    list,
	})
}

// Get get user data
func (u *Users) Get(c *gin.Context) {
	context := Context{c}

	uid := context.Param("uid")
	if uid == "" {
		context.ErrorBadRequest(errors.New("uid can not be null or empty"))
		return
	}

	uidInt32, err := strconv.Atoi(uid)
	if err != nil {
		context.ErrorInternalServer(errors.Wrap(err, "convert uid to int32"))
		return
	}

	model, err := u.find(uidInt32)
	if err != nil {
		context.ErrorInternalServer(errors.Wrap(err, "find user error"))
		return
	}

	if model == nil {
		context.ErrorNotFound(errors.New("user not found"))
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
		"data":    model,
	})
}

// Post store user data
func (u *Users) Post(c *gin.Context) {
	context := Context{c}

	model := Model{}
	if model.Name = c.PostForm("name"); model.Name == "" {
		context.ErrorBadRequest(errors.New("Name is required"))
		return
	}

	if model.Tel = c.PostForm("tel"); model.Tel == "" {
		context.ErrorBadRequest(errors.New("Tel is required"))
		return
	}

	if err := u.store(&model); err != nil {
		context.ErrorInternalServer(errors.Wrap(err, "store user data error"))
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
		"data":    model,
	})

}

func (u *Users) findAll() (*List, error) {
	if u.All == nil {
		u.All = make(map[int]Model)
	}

	return &u.All, nil
}

func (u *Users) find(id int) (*Model, error) {
	if u.All == nil {
		u.All = make(map[int]Model)
	}

	model := u.All[id]
	if model.ID <= 0 {
		return nil, nil
	}
	return &model, nil
}

func (u *Users) store(model *Model) error {
	if u.All == nil {
		u.All = make(map[int]Model)
	}

	idx := u.LastInsertID + 1
	model.ID = idx
	u.All[idx] = *model
	u.LastInsertID = idx
	return nil
}

// ErrorBadRequest self gin context
func (c *Context) ErrorBadRequest(err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "bad_request",
			"detail":  err.Error(),
		})
		return
	}
}

// ErrorInternalServer self gin context
func (c *Context) ErrorInternalServer(err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "internal_server_error",
			"detail":  err.Error(),
		})
		return
	}
}

// ErrorNotFound self gin context
func (c *Context) ErrorNotFound(err error) {
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "not_found",
			"detail":  err.Error(),
		})
		return
	}
}
