package todo

import (
	"fmt"
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
