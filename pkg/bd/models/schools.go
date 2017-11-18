package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// Schools is a model that maps the CSV to a DB Table
type Schools struct {
   Schoolid   string `json:"schoolid"  csv:"schoolID"  db:"schoolID"`
   Namefull   string `json:"namefull"  csv:"name_full"  db:"name_full"`
   City   string `json:"city"  csv:"city"  db:"city"`
   State   string `json:"state"  csv:"state"  db:"state"`
   Country   string `json:"country"  csv:"country"  db:"country"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Schools) GetTableName() string {
  return "schools"
}

// GetFileName returns the name of the source file the model was created from
func (m *Schools) GetFileName() string {
  return "Schools.csv"
}

// GetFilePath returns the path of the source file
func (m *Schools) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Schools.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Schools) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*Schools, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("Schools ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("Schools ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("Schools ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}