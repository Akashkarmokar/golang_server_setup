package handler

import "net/http"

func PostNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle the POST request logic here
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func GetAllNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func GetNewsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func UpdateNewsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func DeleteNewsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
