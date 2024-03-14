package main
// 2 часть программы для разархивации обновлений и копирования их в папку
// используется для обновления антивирусной программы
// на сервере настраиваем запуск в нерабочее время (если первая часть программы копирует в диод в субботу вечером, 
// то вторая часть - к примеру в воскресенье вечером
// запускаем программу, ищем в диоде archive.zip, если находим - разархивируем в папку на сервере d:\\updates
// удаляем архив из диода
////

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func main() {
    archiveFile := "d:\\archive.zip"
    destinationFolder := "d:\\x"// если нужно в сетевую папку, то '\\192.168.0.1\updates'

    // Проверяем наличие файла архива
    if _, err := os.Stat(archiveFile); os.IsNotExist(err) {
        fmt.Println("Файл архива не найден")
        return
    }

    err := unzipFile(archiveFile, destinationFolder)
    if err != nil {
        fmt.Println("Ошибка при разархивации архива:", err)
        return
    }

    fmt.Println("Архив успешно разархивирован в", destinationFolder)
}

func unzipFile(archiveFile, destinationFolder string) error {
    reader, err := zip.OpenReader(archiveFile)
    if err != nil {
        return err
    }
    defer reader.Close()

    for _, file := range reader.File {
        filePath := filepath.Join(destinationFolder, file.Name)

        if file.FileInfo().IsDir() {
            os.MkdirAll(filePath, os.ModePerm)
            continue
        }

        if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
            return err
        }

        writer, err := os.Create(filePath)
        if err != nil {
            return err
        }
        defer writer.Close()

        reader, err := file.Open()
        if err != nil {
            return err
        }
        defer reader.Close()

        if _, err := io.Copy(writer, reader); err != nil {
            return err
        }
    }

    return nil
}