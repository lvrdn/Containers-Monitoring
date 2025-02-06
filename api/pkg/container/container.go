package container

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ContainerHandler struct {
	Storage Storage
}

func NewContainerHandler(storage Storage) *ContainerHandler {
	return &ContainerHandler{
		Storage: storage,
	}
}

type Container struct {
	Addr            string     `json:"addr"`
	Alive           *bool      `json:"alive,omitempty"`
	LastPing        *time.Time `json:"last_ping_time"`
	LastSuccessPing *time.Time `json:"last_alive_time"`
}

type Storage interface {
	AddNewContainerRecord(address string) error
	UpdateContainerRecord(address string, time time.Time, alive bool) error
	ShowContainerRecords() ([]*Container, error)
}

func (ch *ContainerHandler) InitData(addresses []string) error {

	for _, address := range addresses {
		err := ch.Storage.AddNewContainerRecord(address)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ch *ContainerHandler) UpdateData(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body error: error text [%s], path [%s], method [%s]\n", err.Error(), r.URL.Path, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	container := &Container{}
	err = json.Unmarshal(body, container)
	if err != nil {
		log.Printf("unmarshal body error: error text [%s], path [%s], method [%s], body [%s]\n", err.Error(), r.URL.Path, r.Method, string(body))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if container.Alive == nil {
		log.Printf("alive value nil: address [%s]\n", container.Addr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if container.LastPing == nil {
		log.Printf("last ping time is nil: address [%s]\n", container.Addr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ch.Storage.UpdateContainerRecord(container.Addr, *container.LastPing, *container.Alive)
	if err != nil {
		format := "update container record error: error text [%s], addr [%s], last ping [%s], alive [%t]\n"
		log.Printf(format, err.Error(), container.Addr, container.LastPing.String(), container.Alive)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ch *ContainerHandler) ShowData(w http.ResponseWriter, r *http.Request) {

	containers, err := ch.Storage.ShowContainerRecords()
	if err != nil {
		log.Printf("get container records error: error text [%s]\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dataToSend, err := json.Marshal(containers)
	if err != nil {
		log.Printf("marshal containers records error: error text [%s]\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(dataToSend)
}
