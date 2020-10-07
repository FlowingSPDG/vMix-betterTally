package main

import (
<<<<<<< Updated upstream
=======
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
>>>>>>> Stashed changes
	"github.com/sirupsen/logrus"

	vmix "github.com/FlowingSPDG/vmix-go-TCP"
)

var (
	vm = *&vmix.Vmix{}
)

func init() {
	logrus.Debugln("START...")

	var err error // Declare error to avoid "vm" var's shadowing
	vm, err = vmix.New()
	if err != nil {
		panic(err)
	}
}

func main() {
<<<<<<< Updated upstream
=======
	r := gin.Default()
	m = melody.New()

	entrypoint := "./static/index.html"
	r.GET("/", func(c *gin.Context) { c.File(entrypoint) })
	r.Use(static.Serve("/css", static.LocalFile("./static/css", false)))
	r.Use(static.Serve("/js", static.LocalFile("./static/js", false)))
	r.Use(static.Serve("/img", static.LocalFile("./static/img", false)))
	r.Use(static.Serve("/fonts", static.LocalFile("./static/fonts", false)))

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

	m.HandleConnect(func(s *melody.Session) {
		// m.Broadcast(msg)
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
>>>>>>> Stashed changes

}
