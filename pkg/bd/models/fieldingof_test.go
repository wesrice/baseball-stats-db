package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameFieldingOF(t *testing.T) {
  out := FieldingOF{}
  expectedValue := "fieldingof"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameFieldingOF(t *testing.T) {
  out := FieldingOF{}
  expectedValue := "FieldingOF.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathFieldingOF(t *testing.T) {
  out := FieldingOF{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/FieldingOF.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVFieldingOFForError(t *testing.T) {
  out := FieldingOF{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling FieldingOF.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
