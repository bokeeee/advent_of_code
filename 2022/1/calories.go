package main

import (
        "log"
        "os"
        "fmt"
        "bufio"
        "strconv"
)

func main() {
        filePath := os.Args[1]
        file, err := os.Open(filePath)
        if err != nil {
                log.Fatal(err)
        }
                
        fileScan := bufio.NewScanner(file)
        fileScan.Split(bufio.ScanLines)
        var fileLines []string

        for fileScan.Scan() {
                fileLines = append(fileLines, fileScan.Text())
        }

        file.Close()

        max := 0
        current_total := 0

        for _, line := range fileLines {
                if line == "" {
                        if current_total > max {
                                max = current_total
                        }
                        current_total = 0
                        continue
                }
                calorie, err := strconv.Atoi(line)
                if err != nil{
                        log.Fatal(err)
                }
                current_total = current_total + calorie
        }
        
        fmt.Printf("The most total calories carried by an elf is: %d\n", max)
}
