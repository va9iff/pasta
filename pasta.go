package main

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "bufio"
    "strings"
)

// regex to match the [pasta: ](path)
// $1 = title (including pasta); $2 = path
var re = regexp.MustCompile(`\[\s*(pasta\s*:\s*.*)\s*\]\s*\((.*)\)`)

func pasta(builder *strings.Builder, relPath string) {
    absolutePath, err := filepath.Abs(relPath)
    if err != nil { fmt.Println("Error getting absolute path:", err); return }

    if len(os.Args) == 3 {
        fmt.Printf("pasta: %s\n", absolutePath)
    }

    // Open the file
    file, err := os.Open(absolutePath)
    if err != nil { fmt.Println("Error opening file:", err); return }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    if err := scanner.Err(); err != nil { fmt.Println("Error reading file:", err); return }

    for scanner.Scan() {
        line := scanner.Text()
        if re.MatchString(line) {
            embeddingPath, err := filepath.Abs(re.ReplaceAllString(line, "$2"))
            if err != nil { fmt.Println("Error getting absolute path of embedded document:", err); return }
            pasta(builder, embeddingPath)
        } else {
            builder.WriteString(line + "\n")
        }
    }

}

func main() {
    if len(os.Args) < 2 || len(os.Args) > 3 {
        fmt.Println("usage:")
        fmt.Println("pasta input.md")
        fmt.Println("pasta input.md output.md")
        return
    }

    filename := os.Args[1]

    var builder strings.Builder
    pasta(&builder, filename)

    content := builder.String()

    if len(os.Args) == 2 {
        fmt.Printf("%s", content)
    }

    if len(os.Args) == 3 {
        outputPath, err := filepath.Abs(os.Args[2])
        if err != nil { fmt.Println("Error getting absolute path of output:", err); return }

        outputFile, err := os.Create(outputPath)
        if err != nil {
            fmt.Println("Error creating output file:", err)
            return
        }
        defer outputFile.Close()

        _, err = outputFile.WriteString(content)
        if err != nil {
            fmt.Println("Error writing to output file:", err)
            return
        }
    }
}

