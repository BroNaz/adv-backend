package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/BroNaz/adv-backend/internal/app/store/sqlstore"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

/*
func (s *APIServer) apiHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//идея засунуть сюда свегер и выводить доступные методы
		file, err := ioutil.ReadFile("static/apiHelper.html")
		if err != nil {
			s.logger.Error("Error apiHelper")
			rw.Write([]byte("501 - Not Implemented"))
			return
		}
		rw.Write(file)
	}
}
*/
