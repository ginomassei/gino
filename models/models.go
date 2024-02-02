package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	RolId    primitive.ObjectID `json:"rolId" bson:"rolId"`
	Changed  bool               `json:"changed" bson:"changed"`
}

type Employee struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LastName    string             `json:"lastName" bson:"lastName"`
	FirstName   string             `json:"firstName" bson:"firstName"`
	Email       string             `json:"email" bson:"email"`
	BirthDate   time.Time          `json:"birthDate" bson:"birthDate"`
	Cellphone   string             `json:"cellphone" bson:"cellPhone"`
	JoinDate    time.Time          `json:"joinDate" bson:"joinDate"`
	GitlabId    int                `json:"gitlabId" bson:"gitlabId"`
	DiscordUser string             `json:"discordUser" bson:"discordUser"`
	AreaId      primitive.ObjectID `json:"areaId" bson:"areaId"`
	SeniorityId primitive.ObjectID `json:"seniorityId" bson:"seniorityId"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
}

type Rol struct {
	Id   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}
