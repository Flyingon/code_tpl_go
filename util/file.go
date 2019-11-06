package util

import (
    "os"
    "bufio"
)

func ReadLines(path string) ([]string, error) {
    inFile, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer inFile.Close()
    lines := make([]string, 0)
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, nil
}
