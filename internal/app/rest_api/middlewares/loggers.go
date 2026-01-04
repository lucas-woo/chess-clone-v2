package middlewares

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
);

func Logger() {
	loggerPath := filepath.Join("logs", "gin.log")
	file, err := os.Create(loggerPath);
	if err != nil {
		log.Fatal("logger failed")
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func LoggerFunc() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s - %d", p.Method, p.Path, p.StatusCode)
	})
}