package main

import (
 "fmt"
 "net/http"
 "os"
 "path/filepath"
 "time"
)

const (
 numFiles  = 200
 directory = "/tmp/server"
)

func main() {
 // Создаем каталог, если он не существует
 if err := os.MkdirAll(directory, 0755); err != nil {
  fmt.Printf("Error creating directory: %v\n", err)
  return
 }

 http.HandleFunc("/bugs", bugsHandler)

 fmt.Println("Starting server at :8080")
 if err := http.ListenAndServe(":8080", nil); err != nil {
  fmt.Printf("Server failed: %s\n", err)
 }
}

func bugsHandler(w http.ResponseWriter, r *http.Request) {
 for i := 0; i < numFiles; i++ {
  // Генерируем уникальное имя файла
  fileName := fmt.Sprintf("file_%d.txt", time.Now().UnixNano())
  filePath := filepath.Join(directory, fileName)

  // Открываем файл для записи
  file, err := os.Create(filePath)
  if err != nil {
   http.Error(w, fmt.Sprintf("Could not create file: %v", err), http.StatusInternalServerError)
   return
  }

  // Записываем текущее время в файл
  currentTime := time.Now().Format(time.RFC3339)
  if _, err := file.WriteString(currentTime); err != nil {
   http.Error(w, fmt.Sprintf("Could not write to file: %v", err), http.StatusInternalServerError)
   file.Close()
   return
  }

  // Закрываем файл (но по вашему запросу файлы не закрываем)
  // file.Close()
 }

 fmt.Fprintf(w, "Created and wrote current time to %d files.\n", numFiles)
}