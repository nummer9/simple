package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/exp/slog"
)

// RandomWaitHandler waits random amount of time
type RandomWaitHandler struct{}

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

	slog.Info("received web-request and waited "+strconv.Itoa(wait)+" ms", slog.String("route", "/random-wait"))

	fmt.Fprintf(w, "waited %d ms", wait)
}
