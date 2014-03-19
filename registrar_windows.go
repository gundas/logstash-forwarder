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

  encoder := json.NewEncoder(file)
  encoder.Encode(state)
  file.Close()

  old := path + ".old"
  os.Rename(path, old)
  os.Rename(tmp, path)
}
