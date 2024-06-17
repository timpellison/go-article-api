package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"goservice/config"
	"net"
	"net/http"
	"strconv"
)

func main() {
	config := config.NewConfiguration("./config.yml")
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		addr, err := net.LookupHost(config.Database.Cluster)
		w.WriteHeader(http.StatusOK)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(addr[0]))
		}
	})

	router.HandleFunc("/health/db", func(w http.ResponseWriter, r *http.Request) {
		dsn := "host=" + config.Database.Cluster +
			" user=" + config.Database.UserName +
			" port=" + strconv.FormatInt(int64(config.Database.Port), 10) +
			" dbname=" + config.Database.DatabaseName +
			" password=" + config.Database.Password +
			" sslmode=" + config.Database.SslMode +
			" TimeZone=UTC connect_timeout=8"

		_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			fmt.Fprint(w, "The database is ready &s", config.Database.Cluster)
		}
	})

	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}
