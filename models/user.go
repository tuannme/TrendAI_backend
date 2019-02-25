package models

import (
	"cloud.google.com/go/firestore"
	"errors"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"time"
)

var userCollection *firestore.CollectionRef

func init() {
	userCollection = databases.GetFirestoreCollection("user")
}

func GetUserCollection() *firestore.CollectionRef {
	return userCollection
}

const (
	UserGenderMale   = 0
	UserGenderFemale = 1
)

var UserGenders = map[int]string{
	UserGenderMale:   "male",
	UserGenderFemale: "female",
}

type User struct {
	Id            string         `json:"id,omitempty" firestore:"-"`
	Name          string         `json:"name" firestore:"name"`
	Email         string         `json:"email" firestore:"email"`
	Gender        int            `json:"gender" firestore:"gender"`
	Dob           time.Time      `json:"dob" firestore:"dob"`
	Education     string         `json:"education" firestore:"education"`
	ExternalUsers []ExternalUser `json:"external_users" firestore:"external_users"`
	CreatedAt     time.Time      `json:"created_at" firestore:"created_at"`
}

type ExternalUser struct {
	AppId           string    `json:"app_id" firestore:"app_id"`
	UserId          string    `json:"user_id" firestore:"user_id"`
	LastConnectedAt time.Time `json:"last_connected_at" firestore:"last_connected_at"`
	CreatedAt       time.Time `json:"created_at" firestore:"created_at"`
}

type UserResponse struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Dob       time.Time `json:"dob"`
	Education string    `json:"education"`
	CreatedAt time.Time `json:"created_at"`
}

// Get response data for current user
func (u *User) ToResponse() UserResponse {
	gender, _ := UserGenderToStr(u.Gender)
	return UserResponse{
		Id:        u.Id,
		Name:      u.Name,
		Email:     u.Email,
		Gender:    gender,
		Dob:       u.Dob,
		Education: u.Education,
		CreatedAt: u.CreatedAt,
	}
}

// Convert gender in int to string
func UserGenderToStr(gender int) (string, error) {
	val, exists := UserGenders[gender]
	if !exists {
		return "", errors.New("gender doesn't exists")
	}
	return val, nil
}

// Convert gender in string to int
func UserGenderToInt(gender string) (int, error) {
	for k, v := range UserGenders {
		if v == gender {
			return k, nil
		}
	}
	return 0, errors.New("gender doesn't exists")
}
