package services

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionService struct {
}

type SessionServiceInterface interface {
	Set(ctx *gin.Context, key string, v any) error
	Get(ctx *gin.Context, key string, i any)
	Remove(ctx *gin.Context, key string)
}

func NewSessionService() SessionServiceInterface {
	return &SessionService{}
}

func (ss *SessionService) Set(ctx *gin.Context, key string, v interface{}) error {

	if ctx == nil {
		return errors.New("ctx is required but is missing in current request")
	}

	s := ss.initialize(ctx)

	bt, err := json.Marshal(v)

	if err != nil {
		return err
	}

	s.Set(key, string(bt))

	return s.Save()
}

func (ss *SessionService) Get(ctx *gin.Context, key string, i interface{}) {

	if ctx == nil {
		return
	}

	s := ss.initialize(ctx)

	r := s.Get(key)

	if r == nil {
		return
	}

	err := json.Unmarshal([]byte(r.(string)), i)

	if err != nil {
		log.Printf("Error on Unmarshal record %s", err.Error())
	}

}

func (ss *SessionService) Remove(ctx *gin.Context, key string) {

	if ctx == nil {
		return
	}

	s := ss.initialize(ctx)

	s.Delete(key)
}

func (ss *SessionService) initialize(ctx *gin.Context) sessions.Session {
	return sessions.Default(ctx)
}
