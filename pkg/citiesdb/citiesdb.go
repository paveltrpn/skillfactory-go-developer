package main

import (
	"container/list"
	"log"
	"os"
)

type CityInfo struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	Region     string `json:"Region"`
	District   string `json:"District"`
	Population int    `json:"Population"`
	Foundation int    `json:"Foundation"`
}

type CitiesDB struct {
	DbFileName string
	DbFilePtr  os.File

	// Only linear search
	Db list.List
}

func (db *CitiesDB) Init(fname string) {

}

func (db *CitiesDB) Close() {
	db.DbFilePtr.Close()
	db.Dump()
}

// Dump all data to disc
func (db *CitiesDB) Dump() {

}

// Find city info by Id
func (db *CitiesDB) GetById(id int) CityInfo {
	return CityInfo{}
}

func (db *CitiesDB) Add(infoStruct CityInfo) {
	db.Db.PushBack(infoStruct)
}

func (db *CitiesDB) Delete(id int) {

}

func (db *CitiesDB) UpdatePopulationInfo(id int) {
	log.Fatal(500)
}

func (db *CitiesDB) GetByRegion(region string) []CityInfo {
	return nil
}

func (db *CitiesDB) GetByDistrict(district string) []CityInfo {
	return nil
}

func (db *CitiesDB) GetByPopulation(population int) []CityInfo {
	return nil
}

func (db *CitiesDB) GetByFoundation(foundation int) []CityInfo {
	return nil
}
