package services

import (
	"encoding/json"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionService struct {
	session sessions.Session
}

type SessionServiceInterface interface {
	Set(ctx *gin.Context, key string, v interface{}) error
	Get(ctx *gin.Context, key string, i interface{})
	Remove(ctx *gin.Context, key string)
}

func NewSessionService() SessionServiceInterface {
	return &SessionService{}
}

func (ss *SessionService) Set(ctx *gin.Context, key string, v interface{}) error {
	s := ss.initialize(ctx)

	bt, err := json.Marshal(v)

	if err != nil {
		return err
	}

	s.Set(key, string(bt))

	return s.Save()
}

func (ss *SessionService) Get(ctx *gin.Context, key string, i interface{}) {
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
	s := ss.initialize(ctx)

	s.Delete(key)
}

func (ss *SessionService) initialize(ctx *gin.Context) sessions.Session {
	return sessions.Default(ctx)
}
