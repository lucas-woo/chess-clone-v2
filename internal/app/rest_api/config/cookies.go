package config

import (
	"log"
	"os"
)

var (
	CookieSessionIDString = "session_id"
)

func InitCookiesEnv() {
	val, found := os.LookupEnv("COOKIE_SESSION_ID_STRING")
	if !found {
		log.Fatal("error getting redis env")
	}
	CookieSessionIDString = val;
}