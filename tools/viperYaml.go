
package tools

import (
    "os"
    "fmt"
    "bytes"
    "strings"

    "gopkg.in/yaml.v2"
    "github.com/spf13/viper"
)


func saveConfigToFile(filePath string, numSpaces int) error {
    // Get all settings from Viper
    config := viper.AllSettings()

    // Marshal the config to YAML
    configData, err := yaml.Marshal(config)
    if err != nil {
        return fmt.Errorf("Err marshelling config to YAML: %s", err)
    }

    // Control the indentation space number
    spaces := strings.Repeat(" ", numSpaces)
    configData = bytes.Replace(configData, []byte("  "), []byte(spaces), -1)

    // Make sure each first-level key has new line above the previous keys
    prevIndentLevel := 0
    var yamlFmt bytes.Buffer
    lines := strings.Split(string(configData), "\n")

    for _, line := range lines {
        // Determine the current indent level
        indentLevel := len(line) - len(strings.TrimLeft(line, " "))

        // If the previous line had less indent, add a new line before this line
        if prevIndentLevel > 0 && indentLevel == 0 {
            yamlFmt.WriteString("\n")
        }

        yamlFmt.WriteString(line + "\n")
        prevIndentLevel = indentLevel
    }

    err = os.WriteFile(filePath, yamlFmt.Bytes(), 0644)
    if err != nil {
        return fmt.Errorf("err writing config file: %s", err)
    }

    return nil
}


func SaveViper() {
    viper.SetConfigFile("tools/testViper.yaml")
    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("err: %s", err)
    }

    fmt.Println(viper.GetString("cors.methods"))

    err := saveConfigToFile("tools/testViper2.yaml", 4)
    if err != nil {
        fmt.Println("Err saving config file: %s", err)
    }

    fmt.Println("Config file updated")
}

