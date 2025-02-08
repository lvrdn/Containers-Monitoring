package handler

import (
	"api/pkg/container"
	"api/pkg/storage"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ContainerHandler struct {
	Storage storage.Storage
}

func NewContainerHandler(storage storage.Storage) *ContainerHandler {
	return &ContainerHandler{
		Storage: storage,
	}
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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body error: error text [%s], path [%s], method [%s]\n", err.Error(), r.URL.Path, r.Method)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	container := &container.Container{}
	err = json.Unmarshal(body, container)
	if err != nil {
		log.Printf("unmarshal body error: error text [%s], path [%s], method [%s], body [%s]\n", err.Error(), r.URL.Path, r.Method, string(body))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if container.Alive == nil {
		log.Printf("alive is nil: address [%s]\n", container.Addr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if container.LastPing == nil {
		log.Printf("last ping time is nil: address [%s]\n", container.Addr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ch.Storage.UpdateContainerRecord(container)
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
