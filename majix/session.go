package majix

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	. "io"
	"net/http"
	"sync"
	"time"
)

var ctx = context.Background()
var mutex = &sync.Mutex{}

const (
	NAME   = "SID"
	MaxAge = 0 * time.Second
)

type SessionInterface interface {
	Set(key string, Value interface{})
	Get(key string) interface{}
	Delete()
}

type Session struct {
	Name     string
	Value    string
	MaxAge   time.Duration
	Content  map[string]interface{}
	Provider Provider
}

func (s *Session) Set(key string, Value interface{}) {
	mutex.Lock()
	s.Content[key] = Value
	s.Provider.Set(s.Value, s.Content)
	mutex.Unlock()
}

func (s *Session) Get(key string) interface{} {
	content, err := s.Provider.Get(s.Value)
	if err != nil {
		return nil
	}
	if Value, ok := content[key]; !ok {
		return nil
	} else {
		return Value
	}
}

func (s *Session) Delete() {
	s.Provider.Delete(s.Value)
}

func makeSID() string {
	b := make([]byte, 32)
	if _, err := ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)

}

func sessionInit(w http.ResponseWriter, provider Provider, setCookie bool) SessionInterface {
	session := &Session{
		Name:     NAME,
		Value:    makeSID(),
		MaxAge:   MaxAge,
		Content:  make(map[string]interface{}),
		Provider: provider,
	}
	if setCookie {
		cookie := http.Cookie{Name: session.Name, Value: session.Value, MaxAge: int(session.MaxAge)}
		http.SetCookie(w, &cookie)
	}
	return session
}

type SessionManager struct {
	w        http.ResponseWriter
	r        *http.Request
	Provider Provider
}

func NewSessionManager(provider Provider) *SessionManager {
	return &SessionManager{
		Provider: provider,
	}
}

func (sm *SessionManager) Session(w http.ResponseWriter, r *http.Request) SessionInterface {
	sm.r = r
	sm.w = w

	if sid, err := sm.r.Cookie("SID"); err == nil {
		content, err := sm.Provider.Get(sid.Value)

		if err != nil {
			return sessionInit(w, sm.Provider, true)
		}

		session := sessionInit(w, sm.Provider, false).(*Session)
		session.Value = sid.Value
		session.Content = content
		return session
	}

	return sessionInit(w, sm.Provider, true)
}
