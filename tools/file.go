////////////////////////////////////////////////////////////////////////

package tools

import (
    "os"
    "fmt"
    "bufio"
    "strings"

    // for func "GetFilePath"
    "regexp"
    "path/filepath"
)

////////////////////////////////////////////////////////////////////////


func GetFilePath(list *[]string, full_pth bool, extension string, file_dir string) error {
    // Regular expression to filter out the unwanted file path.
    valid := regexp.MustCompile(`^[^\.].*(` + extension + `)$`)
    // Loop through the folders and files that exist in the targeted folder.
    err := filepath.Walk(file_dir, 
        func(path string, info os.FileInfo, err error) error {
            if err != nil {
                fmt.Println(err)
                return err
            }
            if valid.MatchString(info.Name()) {
                if !info.IsDir() && full_pth {
                    // fmt.Printf("%v | %s\n", info.IsDir(), path)
                    *list = append(*list, path)
                } else if !info.IsDir() && !full_pth {
                    //fmt.Printf("%v | %s\n", info.IsDir(), info.Name())
                    *list = append(*list, info.Name())
                }
            }
            return nil
    })
    return err
}


func ReadFile() {
    fmt.Println("\n7-1. read file")
    read, err := os.Open("tools/test.yaml")
    if err != nil {
        fmt.Println(err)
    }

    var lines []string
    var m map[string]string = map[string]string{}

    scanner := bufio.NewScanner(read)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        var s string = scanner.Text()

        if s != "" && strings.Index(s, "#") == -1 {
            idxSplit := strings.Index(s, ":")
            if idxSplit != -1 {  // If the split sign ":" is found
                var key string = strings.Replace(s[:idxSplit], " ", "", -1)
                var val string = strings.Replace(s[idxSplit + 1:], " ", "", -1)
                m[key] = val
            }

            lines = append(lines, s)
        }
    }
    fmt.Println(lines)
    read.Close()
}


func ParamTest() {
    fmt.Println("\n7-2. config file api")
    p := InitParam("tools/test.yaml", false)

    io, err := p.keyExist("BATCH")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(io)
    }

    val, err2 := p.getValue("MODEL_NAME")
    if err2 != nil {
        fmt.Println(err2)
    } else {
        fmt.Println(val)
    }

    _, err3 := p.updateValue("MODEL_NAME", "hello world")
    if err3 != nil {
        fmt.Println(err3)
    }

    _, err4 := p.setValue("NEW_NAME", "fresh")
    if err4 != nil {
        fmt.Println(err4)
    }

    p.saveParams()
}


type Param struct {
    totalParam int
    saveBack bool
    fileName string
    params map[string]string
}


func InitParam(file string, save bool) Param {
    // Init a new struct and the corresponding parameters
    p := Param{}
    p.totalParam = 0
    p.saveBack = save
    p.fileName = file

    // Read the targeted file and put data into map
    read, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
    }

    var m map[string]string = map[string]string{}

    scanner := bufio.NewScanner(read)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        var s string = scanner.Text()

        if s != "" && strings.Index(s, "#") == -1 {
            idxSplit := strings.Index(s, ":")
            if idxSplit != -1 {  // If the split sign ":" is found
                var key string = strings.Replace(s[:idxSplit], " ", "", -1)
                var val string = strings.Replace(s[idxSplit + 1:], " ", "", -1)
                m[key] = val

                p.totalParam += 1
            }
        }
    }
    read.Close()

    p.params = m
    return p
}


func (p *Param) keyExist(key string) (bool, error) {
    if p.params[key] == "" {
        return false, fmt.Errorf("key not found: %s", key)
    }

    return true, nil
}


func (p *Param) getValue(key string) (string, error) {
    _, err := p.keyExist(key)
    if err != nil {
        return "", fmt.Errorf("key not found: %s", key)
    }

    return p.params[key], nil
}


func (p *Param) updateValue(key string, value string) (bool, error) {
    _, err := p.keyExist(key)
    if err != nil {
        return false, fmt.Errorf("key not found: %s", key)
    }

    p.params[key] = value
    return true, nil
}


func (p *Param) setValue(key string, value string) (bool, error) {
    _, err := p.keyExist(key)

    if err != nil {
        val, _ := p.getValue("CONFIG_FIX_PARAM_NUM")

        if val == "true" {
            return false, fmt.Errorf("refuse to add new param: %s", key)
        }

        p.params[key] = value
        p.totalParam += 1
        return true, nil
    }

    // key already exists, failed to setup value
    return false, fmt.Errorf("key already exists: %s", key)
}


func (p *Param) saveParams() (bool, error) {
    upd, _ := p.getValue("CONFIG_UPDATE")
    var lines []string

    if p.saveBack || upd == "true" {
        // Read the targeted file and put data into map
        read, err := os.Open(p.fileName)
        if err != nil {
            return false, fmt.Errorf("invalid file: %s", p.fileName)
        }

        scanner := bufio.NewScanner(read)
        scanner.Split(bufio.ScanLines)
        for scanner.Scan() {
            var s string = scanner.Text()

            if s != "" && strings.Index(s, "#") == -1 {
                idxSplit := strings.Index(s, ":")
                if idxSplit != -1 {  // If the split sign ":" is found
                    var key string = strings.Replace(s[:idxSplit], " ", "", -1)
                    var cfg string = key + " : " + p.params[key]
                    lines = append(lines, cfg)

                } else {
                    lines = append(lines, s)
                }

            } else {
                lines = append(lines, s)
            }
        }
        read.Close()

        // Save lines to the file
        destination, err := os.Create(p.fileName)
        if err != nil {
            return false, fmt.Errorf("failed to save file: %s", p.fileName)
        }
        defer destination.Close()

        for i := 0; i < len(lines); i++ {
            fmt.Fprintf(destination, lines[i] + "\n")
        }

        return true, nil
    }

    return false, fmt.Errorf("no permission to save file: %s", p.fileName)
}


