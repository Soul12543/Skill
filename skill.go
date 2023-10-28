package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SkillList struct {
	List []struct {
		Name        string `json:"Name"`
		Age         int    `json:"Age"`
		Description string `json:"Description"`
		Skill       []struct {
			Name        string `json:"Name"`
			Description string `json:"Description"`
		} `json:"Skill"`
		Ultra struct {
			Name        string `json:"Name"`
			Description string `json:"Description"`
		} `json:"Ultra"`
	} `json:"List"`
}

type requestJson struct {
	Method      string `json:"Method"`
	Name        string `json:"Name"`
	Age         int    `json:"Age"`
	Description string `json:"Description"`
	Skill       []struct {
		Name        string `json:"Name"`
		Description string `json:"Description"`
	} `json:"Skill"`
	Ultra struct {
		Name        string `json:"Name"`
		Description string `json:"Description"`
	} `json:"Ultra"`
}

func ReadJson() SkillList {
	empty := SkillList{}

	file, err := os.Open("./novel/txt/人物技能表.json")
	if err != nil {
		fmt.Printf("[Error]:打开 人物技能表.json 失败，原因：%s\n", err.Error())
		return empty
	}
	defer file.Close()

	jsondata, _ := io.ReadAll(file)
	var skillList SkillList
	err = json.Unmarshal(jsondata, &skillList)
	if err != nil {
		fmt.Printf("解析json失败，原因%s\n", err.Error())
	}

	return skillList
}

func readAll(w http.ResponseWriter) []byte {
	fmt.Println("正在读取所有人物信息。")

	jsondata, err := json.Marshal(ReadJson())
	if err != nil {
		fmt.Printf("格式化json失败，原因%s\n", err.Error())
		return []byte("")
	}

	w.Write(jsondata)
	w.WriteHeader(http.StatusOK)
	fmt.Println("成功")
	return jsondata
}

func ReadPerson(data requestJson, w http.ResponseWriter) []byte {
	fmt.Println("正在读取所有人物信息。")

	jsondata, err := json.Marshal(ReadJson())
	if err != nil {
		fmt.Printf("格式化json失败，原因%s\n", err.Error())
		return []byte("")
	}

	w.Write(jsondata)
	w.WriteHeader(http.StatusOK)
	fmt.Println("成功")
	return jsondata
}

func ModifySkill(data requestJson, w http.ResponseWriter) {
	fmt.Println("正在修改人物技能。")

	oriJson := ReadJson()
	flag := true
	for _, List := range oriJson.List {
		if List.Name == data.Name {
			flag = false
			break
		}
	}

	if flag {
		w.Write([]byte("未找到人物，请使用Add方式添加新人物"))
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Write([]byte("修改完成"))
	w.WriteHeader(http.StatusOK)
	fmt.Println("成功")
	return
}

func ModifyPerson(data requestJson, w http.ResponseWriter) {
	fmt.Println("正在修改人物信息。")

	oriJson := ReadJson()
	flag := true
	for _, List := range oriJson.List {
		if List.Name == data.Name {
			flag = false
			break
		}
	}

	if flag {
		w.Write([]byte("未找到人物，请使用Add方式添加新人物"))
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Write([]byte("修改完成"))
	w.WriteHeader(http.StatusOK)
	fmt.Println("成功")
	return
}

func DeletePerson(data requestJson, w http.ResponseWriter) {
	fmt.Println("正在修改人物信息。")
	w.Write([]byte("修改完成"))
	w.WriteHeader(http.StatusOK)
	fmt.Println("成功")
	return
}

func DeleteSkill(data requestJson, w http.ResponseWriter) {
	fmt.Println("正在修改人物信息。")
	w.Write([]byte("修改完成"))
	w.WriteHeader(http.StatusOK)
	fmt.Println("成功")
	return
}

func AddPerson(data requestJson, w http.ResponseWriter) {
	fmt.Println("正在修改人物信息。")
	w.Write([]byte("修改完成"))
	w.WriteHeader(http.StatusOK)
	fmt.Println("成功")
	return
}

func readRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("正在分析数据")

	var request requestJson
	req, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(req, &request)
	if err != nil {
		fmt.Println("解析Json失败")
		w.Write([]byte("403 Forbidden\n解析Json失败"))
		w.WriteHeader(http.StatusForbidden)
		return
	}

	switch request.Method {
	case "ModifyPerson":
		ModifyPerson(request, w)
		fmt.Println("成功")
	case "ModifySkill":
		ModifySkill(request, w)
		fmt.Println("成功")
	case "DeletePerson":
		DeletePerson(request, w)
		fmt.Println("成功")
	case "DeleteSkill":
		DeleteSkill(request, w)
		fmt.Println("成功")
	case "AddPerson":
		AddPerson(request, w)
		fmt.Println("成功")
	case "ReadPerson":
		ReadPerson(request, w)
		fmt.Println("成功")
	case "ReadAll":
		readAll(w)
		fmt.Println("成功")
	default:
		w.Write([]byte("403 Forbidden\n未知的方法"))
		w.WriteHeader(http.StatusForbidden)
	}

}

func main() {
	cors := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == "OPTIONS" {
				return
			}
			h.ServeHTTP(w, r)
		})
	}
	http.HandleFunc("/", readRequest)
	fmt.Println("开始监听8010端口")

	http.ListenAndServe("localhost:8010", cors(http.DefaultServeMux))

}
