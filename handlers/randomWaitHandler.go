package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// RandomWaitHandler waits random amount of time
type RandomWaitHandler struct {
}

var (
	r   *rand.Rand
	one sync.Once
)

func (rcv RandomWaitHandler) ServeHTTP(w http.ResponseWriter, rq *http.Request) {

	one.Do(func() {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	})

	wait := r.Intn(100)
	time.Sleep(time.Duration(wait) * time.Millisecond)

	log.Info("received web-request to /random-wait, waited " + strconv.Itoa(wait) + " ms")

	w.WriteHeader(http.StatusOK)

}
