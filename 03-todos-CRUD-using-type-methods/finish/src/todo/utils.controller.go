package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	commonController "github.com/nivb52/go-branches-learning/02-todos-CRUD-using-type-methods/finish/src/common"
)

// controller utils
func getIdFromUrl(url string) string {
	matches := todoNamedFieldIdRegex.FindStringSubmatch(url)
	var entityId string
	for i, name := range todoNamedFieldIdRegex.SubexpNames() {
		if name == "ID" {
			entityId = matches[i]
		}
	}
	fmt.Printf("entityId: %s\n", entityId)

	return entityId
}

func makeSuccessResponse(w http.ResponseWriter, r *http.Request) func(*Todo) {
	return func(data *Todo) {
		jsonBytes, err := json.Marshal(&data)
		if err != nil {
			commonController.InternalServerError(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	}
}

func makeBatchSuccessResponse(w http.ResponseWriter, r *http.Request) func(*[]Todo) {
	return func(data *[]Todo) {
		jsonBytes, err := json.Marshal(&data)
		if err != nil {
			commonController.InternalServerError(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	}
}
