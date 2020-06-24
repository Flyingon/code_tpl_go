package util

import (
    "fmt"
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

// CreateDir create dir if it doesn't exist,
// return error if `dir` existed while it is not a directory,
// return error if any other error occurs.
func CreateDir(fileDir string) error {
    var err error
    fin, err := os.Lstat(fileDir)
    if err != nil {
        if !os.IsNotExist(err) {
            return err
        }
        return os.MkdirAll(fileDir, os.ModePerm)
    }
    if !fin.IsDir() {
        return fmt.Errorf("target %s already existed", fileDir)
    }
    return nil
}