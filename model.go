package main

import (
	"database/sql"
	"errors"
)

type users struct {
	phoneUser    string `json:"phoneUser"`
	nameUser     string `json:"nameUser"`
	infoUser     string `json:"infoUser"`
	photoUser    string `json:"photoUser"`
	lastseenUser string `json:"lastseenUser"`
}

func (p *users) getUsers(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *users) updateUsers(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *users) deleteUsers(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *users) createUsers(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getUsers(db *sql.DB, start, count int) ([]users, error) {
	return nil, errors.New("Not implemented")
}
