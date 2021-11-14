package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sports/pkg/common/base"
	"sports/pkg/logger"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Port int    `json:"port"`
	Env  string `json:"env"`
}

type GinEngine struct {
	g      *gin.Engine
	config Config
}

type Option func(engine *gin.Engine)

func AddMiddleware(middleware ...gin.HandlerFunc) Option {
	return func(engine *gin.Engine) {
		engine.Use(middleware...)
	}
}

func NewGinEngine(config *Config, options ...Option) *GinEngine {
	g := gin.New()
	engine := &GinEngine{g: g, config: *config}
	engine.useConfig()
	for _, option := range options {
		option(engine.g)
	}
	return engine
}

func (e *GinEngine) useConfig() {
	if e.config.Env != base.EnvProduction {
		e.g.Use(gin.Logger(), gin.Recovery())
	}
	if e.config.Port == 0 {
		e.config.Port = base.ServerPort
	}
}

// 启动服务
func (e *GinEngine) Start() {
	server := &http.Server{
		Addr:           ":" + strconv.Itoa(e.config.Port),
		Handler:        e.g,
		ReadTimeout:    750 * time.Second,
		WriteTimeout:   750 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error("服务监听异常", err)
			panic("服务监听异常：" + err.Error())
		}
	}()
	gracefulExitWeb(server)
}

// nolint
// 平滑退出服务
func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("got a signal", sig)

	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		// 处理退出失败
		fmt.Println("err", err)
	}
	// 看看实际退出所耗费的时间
	fmt.Println("延时", time.Since(now), "退出")
}

// 配置解析
func ParseConfig(config interface{}, configPath string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&config, func(c *mapstructure.DecoderConfig) {
		c.TagName = "json"
	})
}
