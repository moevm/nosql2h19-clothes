package utils

import "fmt"

var SuccessMessage = "success"
var IdKey = "id"

func EntityWithIdNotExistMessage(title string) string {
	return fmt.Sprintf("%s with id not exist", title)
}

func EntityNotUpdatedMessage(title string) string {
	return fmt.Sprintf("%s not updated", title)
}
func EntityNotExistMessage(title string) string {
	return fmt.Sprintf("%s not exist", title)
}
func EntityNotDeletedMessage(title string) string {
	return fmt.Sprintf("%s not deleted", title)
}
func EntityExistMessage(title string) string {
	return fmt.Sprintf("%s with params exist", title)
}
