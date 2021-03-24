package easyjson

//内置的json解析
//利用反射实现，通过fieldtag来标示对应的json值

type BasicInfo struct {
	Name string `json: "name"`
	Age string  `json: "name"`
}
type JobInfo struct {
	Skills []string `json: "skills"`
}
type Employee struct {
	BasicInfo BasicInfo `json: "basic_info"`
	JobInfo JobInfo     `json: "job_info"`
}

