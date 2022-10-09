package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		logrus.Error(err)
	}
}
