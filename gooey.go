package main

import (
        //"encoding/json"
        "fmt"
        "gopkg.in/yaml.v1"
        "io/ioutil"
        "log"
        "os"
)

func getData(dataFile string) (interface{}, error) {
        f, err := os.Open(dataFile)
        if err != nil {
                log.Println(err)
                return nil, err
        }
        data, err := ioutil.ReadAll(f)
        if err != nil {
                log.Println(err)
                return nil, err
        }
        var y interface{}
        if err := yaml.Unmarshal([]byte(data), &y); err != nil {
                log.Println(err)
                return nil, err
        }
        d, err := yaml.Marshal(y)
        if err != nil {
                log.Println(err)
                return nil, err
        }
        return y, nil
}

func main() {
        for _, a := range os.Args[1:] {
                y, err := getData(a)
                if err != nil {
                        log.Println(err)
                        continue
                }
                fmt.Print(y)
        }
}
