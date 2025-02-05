package container

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ContainerHandler struct {
}

func NewContainerHandler() *ContainerHandler {
	return &ContainerHandler{}
}

type Container struct {
	Addr            string    `json:"addr"`
	Alive           bool      `json:"alive"`
	LastPing        time.Time `json:"time"`
	LastSuccessPing time.Time
}

func (ch *ContainerHandler) ReceiveData(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body error: error text [%s], path [%s], method [%s]\n", err.Error(), r.URL.Path, r.Method)
		return
	}

	container := &Container{}
	err = json.Unmarshal(body, container)
	if err != nil {
		log.Printf("unmarshal body error: error text [%s], path [%s], method [%s], body [%s]\n", err.Error(), r.URL.Path, r.Method, string(body))
		return
	}

	log.Printf("%+v", container)

}
