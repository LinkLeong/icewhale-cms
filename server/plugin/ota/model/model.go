package model

type Request struct {
}

type Response struct {
}
type Version struct {
	ChangeLog string `json:"change_log"`
	Version   string `json:"version"`
}
