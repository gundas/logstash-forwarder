// +build !windows

package main

import (
  "encoding/json"
  "log"
  "os"
)

func WriteRegistry(state map[string]*FileState, path string) {
  tmp := path + ".new"
  file, err := os.Create(tmp)
  if err != nil {
    log.Printf("Failed to open %s for writing: %s\n", tmp, err)
    return
  }
  defer file.Close()

  encoder := json.NewEncoder(file)
  encoder.Encode(state)

  os.Rename(tmp, path)
}
