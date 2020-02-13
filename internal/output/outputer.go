package output

import (
	"encoding/json"
	"fmt"
	"log"
)

type Outputer interface {
	Dump(payload interface{})
}

type JSONOutput struct {
}

func (o *JSONOutput) Dump(payload interface{}) {
	res, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Fatalf("Failed to dump payload %v: %s", payload, err)
	}
	fmt.Println(string(res))
}
