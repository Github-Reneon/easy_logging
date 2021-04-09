package easy_logging

import (
  "log"
  io "io/ioutil"
  "os"
  "time"
)

type FileWriter struct {
  Path string
}

func (f *FileWriter) Write(Message string, Print bool) {
  //If file doesn't exist
  if _, err := os.Stat(f.Path); err == nil {
    if err != nil {
      log.Fatal(err)
    }
  } else {
    //Write "" to file to creat a new one
    err2 := io.WriteFile(f.Path, []byte(""), 0644)
    if err2 != nil {
      log.Fatal(err)
      return
    }
  }
  // Open file for append
  file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
      log.Fatal(err)
      return
  }
  //Defer closing after program finish
  defer file.Close()

  //Write to file
  if _, err := file.WriteString(GetDateTime() + Message); err != nil {
    log.Fatal(err)
    return
  }
  if Print == true {
    print(GetDateTime() + Message)
  }
}

func (f *FileWriter) Delete() {
  os.Remove(f.Path)
}

func GetDateTime() string {
  ret := time.Now().String()
  ret = ret[0:25]
  return "[" + ret + "] "
}
