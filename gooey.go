package main

import (
        //"encoding/json"
        "fmt"
        "gopkg.in/yaml.v1"
        "io/ioutil"
        "log"
        "os"
)

func getData(dataFile string) (map[string]interface{}, error) {
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
        var y map[string]interface{}
        if err := yaml.Unmarshal([]byte(data), &y); err != nil {
                log.Println(err)
                return nil, err
        }
        return y, nil
}

type Galaxy struct {
        Name    string
        Size    uint
}

func main() {
        for _, a := range os.Args[1:] {
                y, err := getData(a)
                if err != nil {
                        log.Println(err)
                        continue
                }
                var objs []interface{}
                var obj interface{}
                for _, intf := range y {
                        switch intf {
                        case "Galaxy":
                                obj = Galaxy{}
                        default:
                                log.Printf("Don't know how to create a %v\n", intf)
                        }
                        objs = append(objs, obj)
                }
                fmt.Println(objs)
        }
}
