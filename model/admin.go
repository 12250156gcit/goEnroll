package model

import "myapp/dataStore/postgres"

type Admin struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

const queryInsertAdmin = "insert into admin(firstname, lastname, email, password) values ($1,$2,$3,$4);"

func (adm *Admin) Create() error {
	_, err := postgres.Db.Exec(queryInsertAdmin, adm.FirstName, adm.LastName, adm.Email, adm.Password)
	return err
}

const queryGetAdmin = "select email, password from admin where email= $1 and password=$2;"

func (adm *Admin) Get() error {
	row := postgres.Db.QueryRow(queryGetAdmin, adm.Email, adm.Password)
	return row.Scan(&adm.Email, &adm.Password)
}
