package main

import "github.com/beevik/ntp"
import (
	"fmt"
	"time"
	"os"
	"io"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		io.WriteString(os.Stderr, err.Error())
		io.WriteString(os.Stderr, "\n")
		os.Exit(1)
	}

	time := time.Now().Add(response.ClockOffset)

	fmt.Println(time)
}
