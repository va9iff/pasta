package main

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "bufio"
    "strings"
)

// $before [pasta: $title] ($path) $after
var re = regexp.MustCompile(`(.*)\[\s*(pasta\s*:\s*.*)\s*\]\s*\((.*)\)(.*)`)

func pasta(builder *strings.Builder, relPath string) {
    absolutePath, err := filepath.Abs(relPath)
    if err != nil { fmt.Println("Error getting absolute path:", err); return }

    if len(os.Args) == 3 {
        fmt.Printf("pasta: %s\n", absolutePath)
    }

    file, err := os.Open(absolutePath)
    if err != nil { fmt.Println("Error opening file:", err); return }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    if err := scanner.Err(); err != nil { fmt.Println("Error reading file:", err); return }

    first := true // append \n before every line except first line.
    for scanner.Scan() {
        if !first { builder.WriteString("\n") }
        first = false // so you can use it for embedding assets like "[pasta:](sound.mp3)"
        line := scanner.Text()
        if re.MatchString(line) {
            matches := re.FindStringSubmatch(line)
            embeddingPath, err := filepath.Abs(matches[3])
            if err != nil { fmt.Println("Error getting absolute path of embedded document:", err); return }
            builder.WriteString(matches[1])
            pasta(builder, embeddingPath)
            builder.WriteString(matches[4])
        } else {
            builder.WriteString(line)
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

