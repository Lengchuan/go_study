package main

import (
	"encoding/json"
	"fmt"
)

//通过key查找到json path
func FindPathByKey(j interface{}, key string) (paths []string, err error) {

	var dumpJson func(j interface{}, key string, rootpath string)
	//paths = make([]string, 1, 10)

	//
	dumpJson = func(j interface{}, key string, rootpath string) {
		//遍历map
		iterMap := func(m map[string]interface{}, key string) {

			for k, v := range m {
				tmpPath := rootpath
				tmpPath = fmt.Sprintf("%v.%v", tmpPath, k)
				if k == key {
					paths = append(paths, tmpPath)
				}

				//继续遍历
				dumpJson(v, key, tmpPath)
			}
		}

		//遍历slice
		iterSlice := func(s []interface{}, key string) {
			for i, v := range s {
				tmpPath := rootpath
				tmpPath = fmt.Sprintf("%v.%v", tmpPath, i)
				//继续遍历
				dumpJson(v, key, tmpPath)
			}
		}

		//处理
		switch v := j.(type) {
		//map
		case map[string]interface{}:
			iterMap(v, key)

		//slice
		case []interface{}:
			iterSlice(v, key)

		default:

		}
	}

	dumpJson(j, key, "root")

	return
}

//通过value 查找到 json path
func FindPathByValue(j interface{}, value interface{}) (paths []string, err error) {
	var dumpJson func(j interface{}, value interface{}, rootpath string)
	//paths = make([]string, 1, 10)

	//
	dumpJson = func(j interface{}, value interface{}, rootpath string) {
		//遍历map
		iterMap := func(m map[string]interface{}, value interface{}) {

			for k, v := range m {
				tmpPath := rootpath
				tmpPath = fmt.Sprintf("%v.%v", tmpPath, k)
				if v == value {
					paths = append(paths, tmpPath)
				}

				//继续遍历
				dumpJson(v, value, tmpPath)
			}
		}

		//遍历slice
		iterSlice := func(s []interface{}, value interface{}) {
			for i, v := range s {
				tmpPath := rootpath
				tmpPath = fmt.Sprintf("%v.%v", tmpPath, i)

				switch v.(type) {

				case []interface{}:
					//继续遍历
					dumpJson(v, value, tmpPath)
				case map[string]interface{}:
					//继续遍历
					dumpJson(v, value, tmpPath)

				default:
					if v == value {
						paths = append(paths, tmpPath)
					}
				}
			}
		}

		//处理
		switch v := j.(type) {
		//map
		case map[string]interface{}:
			iterMap(v, value)

		//slice
		case []interface{}:
			iterSlice(v, value)

		default:

		}
	}

	dumpJson(j, value, "root")

	return
}

func main() {

	b := []byte(
		`
		{
			"iw": {
				"Ie": {
					"Itye": {
						"e": "eIe"
					}
				}
			},
			"InnerJSON2": "NoneValue",
			"outterJSON": {
				"innerJSON1": {
					"value1": 10,
					"value2": 22,
					"InnerInnerArray": [
						"test1",
						"test2"
					],
					"InnerInnerJSONArray": [
						{
							"fld1": "val1"
						},
						{
							"fld2": "val2"
						}
					]
				},
				"InnerJSON2": "NoneValue"
			},
			"e": "eeee"
		}
	`)

	j := make(map[string]interface{})
	if err := json.Unmarshal(b, &j); err != nil {
		panic(err)
	}

	var paths []string
	var err error
	if paths, err = FindPathByKey(j, "aaa"); err != nil {
		panic(err)
	}

	for _, v := range paths {
		fmt.Println(v)
	}

	if paths, err = FindPathByValue(j, "test2"); err != nil {
		panic(err)
	}

	for _, v := range paths {
		fmt.Println(v)
	}

}
