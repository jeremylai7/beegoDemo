package vo

type UserVo struct {
	Id         int64  `json:"id""`
	Name       string `json:"name"`
	Age        int    `json:"age""`
	SubmitTime string `json:"submitTime"`
}
