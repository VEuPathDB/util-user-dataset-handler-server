package process

import "time"

type Details struct {
	Started    time.Time `json:"started"`
	UserID     uint      `json:"userId"`
	TarName    string    `json:"tarName"`
	Token      string    `json:"token"`
	InputFiles []string  `json:"files"`
}
