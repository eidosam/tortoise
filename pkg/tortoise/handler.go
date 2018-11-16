package tortoise

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const maxRandomDelay = 30000

// Handler :: HTTP Handler
func Handler(w http.ResponseWriter, r *http.Request) {
	delay, err := strconv.ParseInt(r.URL.Query().Get("delay"), 10, 64)
	if err != nil {
		delay = rand.Int63n(maxRandomDelay)
	}

	time.Sleep(time.Millisecond * time.Duration(delay))
	fmt.Fprintf(w, "Delayed %d ms", delay)
}
