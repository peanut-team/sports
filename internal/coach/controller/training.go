package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sports/pkg/ctr"
	"sports/pkg/errs"
	"sports/pkg/notifier"
)

// @summary 获取实时训练数据
// @Description 实时训练数据
// @Success 200 {object} coach.AthleteTrainingList
// @Failure 500 {object} errs.BasicError
// @Router  /api/v1/coach/training [get]
func GetLiveTraining(c *gin.Context) {
	// 升级
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		err = fmt.Errorf("Failed upgrade: %v ", err)
		ctr.Err(c, err)
		return
	}
	conn, err := notifier.NewUserConnection(ws)
	if err != nil {
		ctr.Err(c, errs.ErrorParams.Params(err))
		return
	}
	conn.Listen()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
