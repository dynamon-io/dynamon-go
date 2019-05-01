package dynamon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Opts specifies dynamon plot options.
type Opts struct {
	Title   string
	XLabel  string
	YLabel  string
	Legend  string   // Used for PlotLine
	Legends []string // Used for PlotLines
}

type widget struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	XLabel string `json:"xLabel"`
	YLabel string `json:"yLabel"`
	Lines  []Line `json:"lines"`
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
