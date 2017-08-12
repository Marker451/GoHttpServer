package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"server/models"
	"server/resource"
	"strconv"
	"strings"
)

func main() {
	//InitLogger()
	//InitConfig()
	err := resource.InitDatabase()
	defer resource.Close()
	if err != nil {
	}

	http.HandleFunc("/index/", index)
	http.HandleFunc("/register-student", registerStudent)
	http.HandleFunc("/register-class", registerClass)
	http.HandleFunc("/get-class-total-score/", getClassTotalScore)
	http.HandleFunc("/get-top-teacher", getTopTeacher)
	err = http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintln(w, r.URL.RequestURI())
}

func registerStudent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	paraList := []string{"id", "classNumber", "score"}
	ret := inputCheck(paraList, r.Form)
	fmt.Fprintf(w, "%t", ret)

	id, err := strconv.Atoi(r.Form["id"][0])
	if err != nil || len(r.Form["id"][0]) != 5 {

	}
	classNumber, err := strconv.Atoi(r.Form["classNumber"][0])
	if err != nil || classNumber < 0 || classNumber > 99 {

	}
	score, err := strconv.Atoi(r.Form["score"][0])
	if err != nil || score < 0 || score > 100 {

	}
	stu := models.StudentInfo{
		Id:       id,
		ClassNum: classNumber,
		Score:    score,
	}
	_, empty, err := models.StudentInfo{Id: id}.Get()
	if err != nil {
	}
	if empty {
		err = stu.Add()
	} else {
		err = models.StudentInfo{Id: id}.Update(stu)
	}

}

func registerClass(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	paraList := []string{"classNumber", "teacher"}
	ret := inputCheck(paraList, r.Form)
	fmt.Fprintf(w, "%t", ret)

	teacher := r.Form["teacher"][0]
	if len(teacher) > 20 {
	}
	classNumber, err := strconv.Atoi(r.Form["classNumber"][0])
	if err != nil || classNumber < 0 || classNumber > 99 {

	}
	class := models.ClassInfo{
		ClassNum: classNumber,
		Teacher:  teacher,
	}
	_, empty, err := models.ClassInfo{ClassNum: classNumber}.Get()
	if err != nil {
	}
	if empty {
		err = class.Add()
	} else {
		err = models.ClassInfo{ClassNum: classNumber}.Update(class)
	}

}

func getClassTotalScore(w http.ResponseWriter, r *http.Request) {
	requestUrl := r.URL.RequestURI()
	paras := strings.Split(requestUrl, "/")
	if len(paras) != 2 {

	}

	id, err := strconv.Atoi(paras[1])
	if err != nil || len(paras[1]) != 5 {

	}
	stu, empty, err := models.StudentInfo{Id: id}.Get()
	if err != nil {
	}
	if empty {
	}
	sum, err := stu.GetScoreSum()
	if err != nil {
	}
	type ret struct {
		Total int    `json:"total"`
		Err   string `json:"error"`
	}
	retVal := ret{}
	if empty {
		retVal.Err = "student-not-found"
	} else {
		retVal.Total = sum
	}

	returnJson(w, retVal)

}

func getTopTeacher(w http.ResponseWriter, r *http.Request) {
	max, err := models.StudentInfo{}.GetMaxScore()
	if err != nil {
	}
	stu, err := models.StudentInfo{}.GetStudentInfoByScore(max)
	class, empty, err := models.ClassInfo{ClassNum: stu.ClassNum}.Get()
	if err != nil || empty {
	}
	type ret struct {
		Teacher string `json:"teacher"`
	}
	returnJson(w, ret{Teacher: class.Teacher})

}

func inputCheck(paraList []string, form url.Values) bool {
	for _, para := range paraList {
		value := form.Get(para)
		if value == "" {
			return false
		}
	}
	return true
}

func returnJson(w http.ResponseWriter, retVal interface{}) {
	data, err := json.Marshal(retVal)
	if err != nil {
	}
	fmt.Fprint(w, data)
}
