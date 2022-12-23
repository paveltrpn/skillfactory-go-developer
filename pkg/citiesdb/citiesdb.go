package citiesdb

import (
	"container/list"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
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

	Db *list.List
}

func (info *CityInfo) updatePopulation(count int) {
	info.Population = count
}

func (db *CitiesDB) Init(fname string) {
	db.Db = list.New()
}

func (db *CitiesDB) Close() {
	db.DbFilePtr.Close()
}

// Dump all data to disc
func (db *CitiesDB) Dump() {

}

func (db *CitiesDB) GetById(id int) (CityInfo, error) {
	var rt CityInfo

	for e := db.Db.Front(); e != nil; e = e.Next() {
		if e.Value.(*CityInfo).Id == id {
			rt = *e.Value.(*CityInfo)
			return rt, nil
		}
	}
	return CityInfo{}, errors.New("not found")
}

func (db *CitiesDB) Add(infoStruct CityInfo) {
	// Make city id from two first bytes of md5 hash from string
	// formed by Name, District and Region fields of CityInfo struct
	toHash := fmt.Sprintf("%v%v%v", infoStruct.Name, infoStruct.District, infoStruct.Region)
	md5Hash := md5.Sum([]byte(toHash))
	infoStruct.Id = int(binary.BigEndian.Uint16(md5Hash[0:2]))

	db.Db.PushFront(&infoStruct)
}

func (db *CitiesDB) Delete(id int) error {
	for e := db.Db.Front(); e != nil; e = e.Next() {
		if e.Value.(*CityInfo).Id == id {
			db.Db.Remove(e)
			return nil
		}
	}

	return errors.New("not found")
}

func (db *CitiesDB) UpdatePopulationInfo(id, newPopulation int) error {
	for e := db.Db.Front(); e != nil; e = e.Next() {
		if e.Value.(*CityInfo).Id == id {
			e.Value.(*CityInfo).Population = newPopulation
			return nil
		}
	}

	return errors.New("not found")
}

func (db *CitiesDB) GetByRegion(region string) []CityInfo {
	return nil
}

func (db *CitiesDB) GetByDistrict(district string) []CityInfo {
	return nil
}

func (db *CitiesDB) GetByPopulation(from, to int) []CityInfo {
	return nil
}

func (db *CitiesDB) GetByFoundation(from, to int) []CityInfo {
	return nil
}
