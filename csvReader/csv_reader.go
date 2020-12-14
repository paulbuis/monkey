package csvReader

// untested package !!!
// TODO: Test this

import (
	"bufio"
	"encoding/csv"
	"monkey/object"
	objectArray "monkey/object/array"
	objectError "monkey/object/error"
	objectHash "monkey/object/hash"
	objectString "monkey/object/string"
	"os"
)

func New(fileName string) object.Object {
	reader, errOpen := os.OpenFile(fileName, os.O_RDONLY, 0)
	if errOpen != nil {
		return objectError.New("error opening CSV file: %s", errOpen)
	}

	bufferedReader := bufio.NewReader(reader)
	csvReader := csv.NewReader(bufferedReader)
	headers, headerErr := csvReader.Read()
	if headerErr != nil {
		return objectError.New("error reading CSV header: %s", headerErr)
	}

	monkeyHeaders := make([]object.Object, 0, len(headers))
	for _, header := range headers {
		monkeyString := objectString.New(header)
		monkeyHeaders = append(monkeyHeaders, monkeyString)
		//monkeyKeys = append(monkeyKeys, monkeyString.HashKey() )
	}

	body, bodyErr := csvReader.ReadAll()
	if bodyErr != nil {
		return objectError.New("error reading CSV file body: %s", bodyErr)
	}

	elements := make([]object.Object, 0, len(body))
	for _, line := range body {
		fields := make([]object.Object, 0, len(headers))
		for _, field := range line {
			fields = append(fields, objectString.New(field))
		}
		hash := objectHash.NewHash(monkeyHeaders, fields)
		elements = append(elements, hash)
	}
	return objectArray.New(elements)
}
