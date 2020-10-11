package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"github.com/sirupsen/logrus"

	vmixgo "github.com/FlowingSPDG/vmix-go"
	vmix "github.com/FlowingSPDG/vmix-go-TCP"
)

var (
	vm = &vmix.Vmix{}
	m  = &melody.Melody{}
)

func init() {
	logrus.Debugln("START...")

	var err error // Declare error to avoid "vm" var's shadowing
	vm, err = vmix.New("localhost")
	if err != nil {
		panic(err)
	}

	// Subscribe tally event
	_, err = vm.SUBSCRIBE("TALLY")
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	m = melody.New()

	entrypoint := "./static/index.html"
	r.GET("/", func(c *gin.Context) { c.File(entrypoint) })
	r.Use(static.Serve("/css", static.LocalFile("./static/css", false)))
	r.Use(static.Serve("/js", static.LocalFile("./static/js", false)))
	r.Use(static.Serve("/img", static.LocalFile("./static/img", false)))
	r.Use(static.Serve("/fonts", static.LocalFile("./static/fonts", false)))

	r.GET("/api/inputs", func(c *gin.Context) {
		body, err := vm.XML()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		v := vmixgo.Vmix{}
		if err := xml.Unmarshal([]byte(body), &v); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, v.Inputs.Input)
	})

	r.GET("/api/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		// Sync tally status
		// m.Broadcast(msg)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// DO NOTHING
	})

	// register callback
	vm.RegisterTallyCallback(func(res *vmix.TallyResponse) {
		msg := make(map[string]vmix.TallyStatus)
		body, err := vm.XML()
		if err != nil {
			logrus.Debugf("Failed to get XML : %v\n", err)
			return
		}
		v := vmixgo.Vmix{}
		err = xml.Unmarshal([]byte(body), &v)
		if err != nil {
			logrus.Debugf("Failed to unmarshal XML : %v\n", err)
			return
		}

		logrus.Debugln("TALLY STATUS :", res.Status)
		for i := 0; i < len(res.Tally); i++ {
			logrus.Debugf("TALLY [%d] STATUS : %s\n", i+1, res.Tally[i].String())
			for _, v := range v.Inputs.Input {
				if int(v.Number) == i+1 { // If input scene matches...
					msg[v.Key] = res.Tally[i]
				}
			}
		}
		b, err := json.Marshal(msg)
		if err != nil {
			logrus.Debugf("Failed to marshal XML : %v\n", err)
			return
		}
		if err := m.Broadcast(b); err != nil {
			logrus.Debugf("Failed to Broadcast data : %v\n", err)
			return
		}
	})

	if err := r.Run(":5000"); err != nil {
		panic(err)
	}
}
