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

type Comet struct {
        Name    string
        Size    uint
}
type Galaxy struct {
        Name    string
        Size    uint
}

func buildObj(name, objtype string, val map[string]interface{}) (interface{}, error) {
        fmt.Println("Creating a ", objtype)
        var obj interface{}
        switch objtype {
        case "Galaxy":
                g := Galaxy{Name: name}
                obj = g
        case "Comet":
                c := Comet{Name: name}
                obj = c
        default:
                return nil, fmt.Errorf("Can't create a ", objtype)
        }
        return obj, nil
}

func build(name string, val interface{}) (interface{}, error) {
        switch val.(type) {
        case string:
                fmt.Println("type is string", name, val)
                return val, nil

        case map[string]interface{}:
                fmt.Println("type is map", name, val)
                vmap := val.(map[string]interface{})
                if t, ok := vmap["Type"]; ok {
                        fmt.Println("Creating a ", t)
                        if objtype, ok := t.(string); ok {
                                return buildObj(name, objtype, vmap)
                        }
                }

                var res map[string]interface{}
                for k, v := range vmap {
                        r, err := build(k, v)
                        if err != nil {
                                return nil, err
                        }
                        res[k] = r
                }
                return res, nil

        case []string:
                fmt.Println("type is []string", name, val)
                return val, nil

        default:
                fmt.Println("type is ??", name, val)
                return nil, fmt.Errorf("Don't know how to create a %v", val)
        }
}

func main() {
        for _, a := range os.Args[1:] {
                y, err := getData(a)
                if err != nil {
                        log.Println(err)
                        continue
                }
                obj, err := build("", y)
                if err != nil {
                        fmt.Println(err)
                        break
                }

                fmt.Println(obj)
        }
}
