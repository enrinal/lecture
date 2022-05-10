package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type StudentRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func handlerStudent(w http.ResponseWriter, r *http.Request) {
	// decode student request body
	var student StudentRequest
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		// set error response code
		w.WriteHeader(http.StatusBadRequest)
		// set error response body
		errResponse := ErrorResponse{
			Error:   "Bad Request",
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}
		err := json.NewEncoder(w).Encode(errResponse)
		if err != nil {
			panic(err)
		}
		return
	}

	// validate student request body
	if student.Name == "" {
		// set error response code
		w.WriteHeader(http.StatusBadRequest)
		// set error response body
		errResponse := ErrorResponse{
			Error:   "Bad Request",
			Code:    http.StatusBadRequest,
			Message: "Name is required",
		}
		err := json.NewEncoder(w).Encode(errResponse)
		if err != nil {
			panic(err)
		}
		return
	}
	if student.Age <= 0 {
		// set error response code
		w.WriteHeader(http.StatusBadRequest)
		// set error response body
		errResponse := ErrorResponse{
			Error:   "Bad Request",
			Code:    http.StatusBadRequest,
			Message: "Age must be greater than 0",
		}
		err := json.NewEncoder(w).Encode(errResponse)
		if err != nil {
			panic(err)
		}
		return
	}

	// hit func createStudent to create student
	err = createStudent(false)
	// check if error is nil, if nil return status code 200, else return error status code 500
	if err != nil {
		// set error response code
		w.WriteHeader(http.StatusInternalServerError)
		// set error response body
		returnErr := ErrorResponse{
			Error:   "Internal Server Error",
			Code:    http.StatusInternalServerError,
			Message: "Error when creating student",
		}
		// encode error response body
		err := json.NewEncoder(w).Encode(returnErr)
		if err != nil {
			panic(err)
		}
		return
	}

	// set response code
	w.WriteHeader(http.StatusOK)
	// set response body
	returnSuccess := SuccessResponse{
		Code:    http.StatusOK,
		Message: "Student created",
	}
	// encode response body
	err = json.NewEncoder(w).Encode(returnSuccess)
	if err != nil {
		panic(err)
	}

	return
}

// example of error handling
func createStudent(isError bool) error {
	if isError {
		returnErr := errors.New("error creating student")
		return returnErr
	}
	return nil
}

func main() {
	http.HandleFunc("/student", handlerStudent)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
