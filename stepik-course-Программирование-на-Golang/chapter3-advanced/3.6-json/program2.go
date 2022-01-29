package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PropertieInner struct {
	Descriptions string `json:"description"`
	Type         string `json:"type"`
}
type Propertie struct {
	SysObjId      PropertieInner `json:"system_object_id,omitempty"`
	Kod           PropertieInner `json:"Kod,omitempty"`
	IsDeleted     PropertieInner `json:"is_deleted,omitempty"`
	SignatureDate PropertieInner `json:"signature_date,omitempty"`
	NomdeScr      PropertieInner `json:"Nomdescr,omitempty"`
	GlobalId      PropertieInner `json:"global_id"`
	Idx           PropertieInner `json:"Idx,omitempty"`
	Razdel        PropertieInner `json:"Razdel,omitempty"`
	Name          PropertieInner `json:"Name,omitempty"`
}
type Item struct {
	Descriptions string      `json:"description"`
	Type         string      `json:"type"`
	Properties   []Propertie `json:"properties"`
	Required     []string    `json:"required"`
}
type JsonStruct struct {
	MinItems     int    `json:"minItems"`
	Schema       string `json:"$schema"`
	Descriptions string `json:"descriptions"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	Items        []Item `json:"items"`
}

func Program2() {
	items := []Item{}
	f, err := os.Open("data-20190514T0100.json")
	if err != nil {
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	//декодируем данные в перем items
	if err := json.Unmarshal(data, &items); err != nil {
		fmt.Println(err)
	}
	fmt.Println(items)
}
