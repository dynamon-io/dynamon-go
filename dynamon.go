package dynamon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type line struct {
	Title  string  `json:"title"`
	Points []point `json:"points"`
}

type widget struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	XLabel string `json:"xLabel"`
	YLabel string `json:"yLabel"`
	Lines  []line `json:"lines"`
}

type payload struct {
	Widgets []widget `json:"widgets"`
}

func api(path string, reqBody payload) {
	reqBodyReader := new(bytes.Buffer)
	json.NewEncoder(reqBodyReader).Encode(reqBody)
	resp, err := http.Post("https://dynamon.io/"+path, "application/json", reqBodyReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Dynamon failed to log.\n%v", err)
		return
	}
	resp.Body.Close()
}
