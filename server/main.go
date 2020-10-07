package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"

	"github.com/FlowingSPDG/vmix-go"
	vmixtcp "github.com/FlowingSPDG/vmix-go-TCP"
)

var (
	vm = &vmixtcp.Vmix{}
	m  = &melody.Melody{}
)

func init() {
	logrus.Debugln("START...")

	// Set GIN Production mode
	gin.SetMode(gin.ReleaseMode)

	var err error // Declare error to avoid "vm" var's shadowing
	vm, err = vmixtcp.New("localhost")
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

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/api/inputs", func(c *gin.Context) {
		_, body, err := vm.XML()
		if err != nil {
			// handle error
		}
		v := vmixgo.Vmix{}
		err = xml.Unmarshal([]byte(body), &v)
		if err != nil {
			// handle error
		}
		c.JSON(http.StatusOK, v.Inputs.Input)
	})

	r.GET("/api/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// m.Broadcast(msg)
	})

	// register callback
	vm.RegisterTallyCallback(func(res *vmixtcp.TallyResponse) {
		msg := make(map[string]vmixtcp.TallyStatus)
		_, body, err := vm.XML()
		if err != nil {
			// handle error
		}
		v := vmixgo.Vmix{}
		err = xml.Unmarshal([]byte(body), &v)
		if err != nil {
			// handle error
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
			// handle error
		}
		m.Broadcast(b)
	})

	r.Run(":5000")
}
