package utils

type EngineName struct {
	Names []string `json:"names"`
}

type TransportMsg struct {
	Sid string `json:"sid"`
	Tid string `json:"tid"`
	Msg []byte `json:"msg"`
}

type ChatData struct {
	Raw []byte `json:"raw"`
}

type UserData struct {
	MsgId       string         `bson:"msgid,omitempty"`
	Name        string         `bson:"name,omitempty"`
	Age         string         `bson:"age,omitempty"`
	PhoneNo     string         `bson:"phone_no,omitempty"`
	Email       string         `bson:"email,omitempty"`
	ProfilePic  string         `bson:"profile_pic,omitempty"`
	MainKey     string         `bson:"main_key,omitempty"`
	Gender      string         `bson:"gender,omitempty"`
	Password    string         `bson:"password,omitempty"`
	Connections map[string]int `bson:"connections,omitempty"`
}

type MsgFormat struct {
	Sid  string `bson:"snum,omitempty"`
	Msg  string `bson:"msg,omitempty"`
	Mloc string `bson:"mloc,omitempty"`
}

type Df struct {
	Id  string            `bson:"_id,omitempty"`
	Msg map[string][]byte `bson:"msg,omitempty"`
}
