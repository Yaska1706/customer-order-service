package cmd

import (
	"net/http"
	"time"
)

// Run initializes executables for the app
func Run() error {
	if err := setupDatabase(); err != nil {
		return err
	}

	s := &http.Server{
		Addr:              ":5000",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	s.ListenAndServe()
	return nil
}
