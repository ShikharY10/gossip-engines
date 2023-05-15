package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PartnerRequest struct {
	ID                string `bson:"id" json:"id"`
	RequesterId       string `bson:"requestedId" json:"requestedId"`
	RequesterUsername string `bson:"requestedUsername" json:"requestedUsername"`
	RequesterName     string `bson:"requestedName" json:"requestedName"`
	TargetId          string `bson:"targetId" json:"targetId"`
	TargetUsername    string `bson:"targetUsername" json:"targetUsername"`
	TargetName        string `bson:"targetName" json:"targetName"`
	PublicKey         string `bson:"publicKey" json:"publicKey"`
	CreatedAt         string `bson:"createdAt" json:"createdAt"`
}

type PartnerResponse struct {
	ID          string `json:"id"`
	IsAccepted  bool   `json:"isAccepted"`
	ResponserId string `json:"responderId"`
	TargetId    string `json:"targetId"`
	SharedKey   string `json:"key"`
}

type RemovePartner struct {
	ID           string `json:"id"`
	ExtraditorId string `json:"extraditorId"`
	ExtraditeeId string `json:"extraditeeId"`
}

type User struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name             string               `bson:"name" json:"name,omitempty"`
	Username         string               `bson:"username" json:"username,omitempty"`
	Email            string               `bson:"email" json:"email,omitempty"`
	Avatar           Avatar               `bson:"avatar" json:"avatar,omitempty"`
	DeliveryId       primitive.ObjectID   `bson:"deliveryId" json:"deliveryId"`
	Posts            []primitive.ObjectID `bson:"posts" json:"posts"`
	Partners         []primitive.ObjectID `bson:"partners" json:"partners,omitempty"`
	PartnerRequests  []PartnerRequest     `bson:"partnerrequests" json:"partnerrequests,omitempty"`
	PartnerRequested []PartnerRequest     `bson:"partnerrequested" json:"partnerrequested,omitempty"`
	Role             string               `bson:"role" json:"role,omitempty"`
	Token            string               `bson:"token" json:"token,omitempty"`
	Logout           bool                 `bson:"logout" json:"logout,omitempty"`
	CreatedAt        string               `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt        string               `bson:"updatedAt" json:"updatedAt,omitempty"`
	DeletedAt        string               `bson:"deletedAt" json:"deletedAt,omitempty"`
}

type DeliveryID struct {
	DeliveryId primitive.ObjectID `bson:"deliveryId" json:"deliveryId"`
}

type Request struct {
	ID             string `bson:"id" json:"id"`
	SenderUsername string `bson:"senderUsername" json:"senderUsername"`
	SenderName     string `bson:"senderName" json:"senderName"`
	TargetUsername string `bson:"targetUsername" json:"targetUsername"`
	TargetName     string `bson:"targetName" json:"targetName"`
	PublicKey      string `bson:"publicKey" json:"publicKey"`
	CreatedAt      string `bson:"createdAt" json:"createdAt"`
}

type Follow struct {
	UserId    string `bson:"userId" json:"userId"`
	Username  string `bson:"username" json:"username"`
	Name      string `bson:"name" json:"name"`
	Avatar    string `bson:"avatar" json:"avatar"`
	CreatedAt string `bson:"createdAt" json:"creaetdAt"`
}

type Avatar struct {
	PublicId  string `json:"publicId" bson:"publicId"`
	FileName  string `json:"fileName" bson:"fileName"`
	SecureUrl string `json:"secureUrl" bson:"secureUrl"`
}

type Transfer struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Payloads bson.M             `bson:"chats" json:"chats"`
}

type Chat struct {
	Key  string `bson:"key" json:"key"`
	Data string `bson:"data" json:"data"`
}

type Packet struct {
	NodeName string `json:"name"`
	Type     string `json:"type"`
	Message  string `json:"message"`
}

type Log struct {
	TimeStamp   string `bson:"timeStamp" json:"timeStamp"`
	ServiceType string `bson:"serviceType" json:"serviceType"`
	Type        string `bson:"type" json:"type"`
	FileName    string `bson:"fileName" json:"fileName"`
	LineNumber  int    `bson:"lineNumber" json:"lineNumber"`
	Message     string `bson:"errorMessage" json:"errorMessage"`
}
