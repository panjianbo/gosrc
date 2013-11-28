package main

import (
  "os"
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func main(){
  db, err := sql.Open("mysql", "ruser:ruser.readonly@webauth@tcp(172.28.28.33:3306)/webauth?charset=utf8")
  if err != nil{
     panic(err.Error())
  }
  defer db.Close()

  err = db.Ping()
  if err != nil{
     panic(err.Error())
  }

  stmt, err := db.Prepare("select id, url from boundary_download_url where security_status in (20, 21) and id > ? order by id limit 5000")
  if err != nil{
    panic(err.Error())
  }
  defer stmt.Close()

  var isFinish bool
  var db_id, id int
  var url string
  db_id = 0
  counter := 0
  fd, _ := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0644) 
  defer fd.Close()
  for{
    rows, err := stmt.Query(db_id)
    if err != nil{
      panic(err.Error())
    }
    
    isFinish = true

    for rows.Next(){
       isFinish = false
       if err:= rows.Scan(&id, &url);err==nil{
          counter ++
          if counter % 10000 == 0{
              fmt.Println(id)
          }
          db_id = id
          fd.WriteString(url+"\n")
       }
    }
    rows.Close()
    if isFinish{
      break
    }
  }
}
