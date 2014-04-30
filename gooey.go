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

func createObj(t string, in interface{}) (interface{}, error) {
        val, ok := in.(map[interface{}]interface{})
        if !ok {
                return nil, fmt.Errorf("Can't convert %v to a map\n", val)
        }
        var obj interface{}
        switch t {
        case "Galaxy":
                obj = Galaxy{}
        default:
                return nil, fmt.Errorf("Don't know how to create a %v\n", t)
        }

        return obj, nil
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
                for t, val := range y {
                        fmt.Println("t is ", t)
                        obj, err := createObj(t, val)
                        if err != nil {
                                fmt.Println(err)
                                continue
                        }
                        objs = append(objs, obj)
                }
                fmt.Println(objs)
        }
}
