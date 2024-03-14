package main
// программа объединения файлов xslx в один файл
////

import (
    "fmt"
    "github.com/tealeg/xlsx/v3"
    "io/ioutil"
    "os"
    "path/filepath"
)

func main() {
    var dirPath string
    fmt.Println("Введите путь к папке с xlsx файлами:")
    fmt.Scanln(&dirPath)

    files, err := ioutil.ReadDir(dirPath)
    if err != nil {
        fmt.Println(err)
        return
    }

    mergedFile := xlsx.NewFile()

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".xlsx" {
            f, err := xlsx.OpenFile(filepath.Join(dirPath, file.Name()))
            if err != nil {
                fmt.Println(err)
                return
            }

            for _, sheet := range f.Sheets {
                newSheet := mergedFile.AddSheet(sheet.Name)
                for _, row := range sheet.Rows {
                    newRow := newSheet.AddRow()
                    for _, cell := range row.Cells {
                        newCell := newRow.AddCell()
                        newCell.SetValue(cell.String())
                    }
                }
            }
        }
    }

    mergedFilePath := filepath.Join(dirPath, "merged.xlsx")
    if err := mergedFile.Save(mergedFilePath); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Файлы успешно объединены в", mergedFilePath)
}