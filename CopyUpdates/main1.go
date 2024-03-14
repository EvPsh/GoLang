
package main
// 1 часть программы 
// используется для архивации баз данных антивируса 
// алгоритм:
// архивируем папку Updates в archive.zip
// настраиваем средствами windows копирование в диод в нерабочее время (к примеру суббота вечер)
// после копирования в диод - удаляем архив
// на этом работа первой части закончена
////

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func main() {
    sourceFolder := '\\wsms\KLSHARE\Updates' // какую папку архивируем
    destinationFile := "d:\\archive.zip"    // куда архивируем

    err := zipFolder(sourceFolder, destinationFile)
    if err != nil {
        fmt.Println("Ошибка при архивации папки:", err)
        return
    }

    fmt.Println("Папка успешно архивирована в", destinationFile)
}

func zipFolder(sourceFolder, destinationFile string) error {
    zipFile, err := os.Create(destinationFile)
    if err != nil {
        return err
    }
    defer zipFile.Close()

    archiveWriter := zip.NewWriter(zipFile)
    defer archiveWriter.Close()

    err = filepath.Walk(sourceFolder, func(filePath string, fileInfo os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !fileInfo.Mode().IsRegular() {
            return nil
        }

        relativePath, err := filepath.Rel(sourceFolder, filePath)
        if err != nil {
            return err
        }

        archivePath := filepath.ToSlash(filepath.Join(filepath.Base(sourceFolder), relativePath))

        file, err := os.Open(filePath)
        if err != nil {
            return err
        }
        defer file.Close()

        header, err := zip.FileInfoHeader(fileInfo)
        if err != nil {
            return err
        }
        header.Name = archivePath
        header.Method = zip.Deflate
        header.Flags |= 1 << 11 // Set the UTF-8 flag

        writer, err := archiveWriter.CreateHeader(header)
        if err != nil {
            return err
        }

        _, err = io.Copy(writer, file)
        if err != nil {
            return err
        }

        return nil
    })

    return err
}



// func archiveFile(filePath string, baseFolder string, archiveWriter *zip.Writer) error {
//     // Открытие файла для чтения
//     file, err := os.Open(filePath)
//     if err != nil {
//         return err
//     }
//     defer file.Close()

//     // Получение информации о файле
//     fileInfo, err := file.Stat()
//     if err != nil {
//         return err
//     }

//     // Создание заголовка для файла в архиве
//     header, err := zip.FileInfoHeader(fileInfo)
//     if err != nil {
//         return err
//     }

//     // Установка пути к файлу относительно базовой папки
//     if baseFolder != "" {
//         header.Name = filepath.Join(baseFolder, header.Name)
//     }

//     // Установка метода сжатия на максимальное сжатие
//     header.Method = zip.Deflate

//     // Создание новой записи файла в архиве
//     writer, err := archiveWriter.CreateHeader(header)
//     if err != nil {
//         return err
//     }

//     // Копирование содержимого файла в архив
//     _, err = io.Copy(writer, file)
//     if err != nil {
//         return err
//     }

//     return nil
// }



// package main

// import (
//     "archive/zip"
//     "fmt"
//     "io"
//     "log"
//     "os"
//     "path/filepath"
// )

// func main() {
//     sourceDir := "D:\\x"
//     archiveDir := "D:\\archive"

//     zipFile := filepath.Join(archiveDir, "archive.zip")

//     err := createArchive(zipFile, sourceDir)
//     if err != nil {
//         log.Fatal(err)
//     }

//     err = copyFile(zipFile, filepath.Join(archiveDir, "archive.zip"))
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Println("Архивация и копирование завершены успешно!")
// }

// func createArchive(zipFile, sourceDir string) error {
//     // Создание нового архива
//     archive, err := os.Create(zipFile)
//     if err != nil {
//         return err
//     }
//     defer archive.Close()

//     zipWriter := zip.NewWriter(archive)
//     defer zipWriter.Close()

//     err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
//         if err != nil {
//             return err
//         }

//         // Создание нового файла в архиве
//         fileInArchive, err := zipWriter.Create(filepath.Join(sourceDir, info.Name()))
//         if err != nil {
//             return err
//         }

//         // Копирование содержимого файла в архив
//         if !info.IsDir() {
//             file, err := os.Open(path)
//             if err != nil {
//                 return err
//             }
//             defer file.Close()

//             _, err = io.Copy(fileInArchive, file)
//             if err != nil {
//                 return err
//             }
//         }

//         return nil
//     })

//     return err
// }

// func copyFile(sourceFile, destinationFile string) error {
//     source, err := os.Open(sourceFile)
//     if err != nil {
//         return err
//     }
//     defer source.Close()

//     destination, err := os.Create(destinationFile)
//     if err != nil {
//         return err
//     }
//     defer destination.Close()

//     _, err = io.Copy(destination, source)
//     return err
// }