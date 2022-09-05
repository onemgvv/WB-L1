package main

import (
	ej "encoding/json"
	ex "encoding/xml"
	"fmt"
	"log"
)

//  Реализовать паттерн адаптер

/*
	Кастомные типы JSON и XML
*/
type json string
type xml string

// Интерфейс для работы с xml данными
type XMLAnalytic interface {
	WorkWithXML() xml
}

// Интерфейс для работы с json данными
type JSONAnalytic interface {
	WorkWithJSON() json
}

// структура JSON и метод получения данных

type JSON struct {
	data json
}

func (j *JSON) JSONDataLoad() json {
	return j.data
}

// структура XML и метод получения данных

type XML struct {
	data xml
}

func (x *XML) XMLDataLoad() xml {
	return x.data
}

/*
	паттерн Адаптер позволяет работать объектам (стурктурам в нашем случае)
	с несовместимыми интерфейсами работать вместе.
	Другими словами он позволяет адаптировать существующий код под наши нужды
	не переписывая его
*/

type XMLAdapter struct {
	*JSON
}

type JSONAdapter struct {
	*XML
}

func NewXMLAdapter(json *JSON) XMLAnalytic {
	return &XMLAdapter{json}
}

func NewJSONAdapter(xml *XML) JSONAnalytic {
	return &JSONAdapter{xml}
}

func (x *XMLAdapter) WorkWithXML() xml {
	var jsonMapData map[string]string
	data := x.JSONDataLoad()
	if err := ej.Unmarshal([]byte(data), &jsonMapData); err != nil {
		log.Fatal("[WRONG JSON]: ", err.Error())
	}
	// some magic with convert
	newData := fmt.Sprintf("%s\n<%s>\n\t<%s>%s</%s>\n<%s>", `<?xml version="1.0" encoding="UTF-8"?>`, "root", "key", jsonMapData["key"], "key", "root")
	return xml(newData)
}

func (j *JSONAdapter) WorkWithJSON() json {
	var xmlMapData map[string]string
	data := j.XMLDataLoad()
	if err := ex.Unmarshal([]byte(data), &xmlMapData); err != nil {
		log.Fatal("[WRONG XML]: ", err.Error())
	}
	// some magic with convert
	newData := fmt.Sprintf("{\"%s\": \"%s\"}", "key", xmlMapData["key"])
	return json(newData)
}