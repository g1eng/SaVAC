package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func EncodeHttpError(r *http.Response, err error) error {
	if err == nil {
		log.Fatalln("No error reported but tried to encode it to a new error")
	}
	if r == nil {
		return fmt.Errorf("no response but raising an error %w", err)
	}
	msg, _ := io.ReadAll(r.Body)
	buf := bytes.NewBufferString("")
	jerr := json.Indent(buf, msg, "", "\t")
	if jerr != nil {
		return fmt.Errorf("invalid json payload: %w: %w", jerr, err)
	}
	return fmt.Errorf("%s", string(msg))
}
