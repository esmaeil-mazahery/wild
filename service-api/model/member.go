package model

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/EsmaeilMazahery/wild/infrastructure/constant"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// NilMember is the nil value for Member
var NilMember Member

// Member contains member's information
type Member struct {
	ID           primitive.ObjectID `bson:"_id"`
	Username     string             `bson:"username"`
	Name         string             `bson:"name"`
	Family       string             `bson:"family"`
	Image        string             `bson:"image"`
	Password     string             `bson:"password"`
	Mobile       string             `bson:"mobile"`
	Email        string             `bson:"email"`
	Enable       bool               `bson:"enable"`
	RegisterDate time.Time          `bson:"register_date"`
	ImageHeader  string             `bson:"image_header"`
	Following    int64              `bson:"following"`
	Follower     int64              `bson:"follower"`
	Biography    string             `bson:"biography"`

	MemberFollow bool `bson:"member_follow,omitempty" json:"member_follow,omitempty"`
}

//MemberMongoType ...
func MemberMongoType() struct {
	ID           string
	Username     string
	Name         string
	Family       string
	Image        string
	Password     string
	Mobile       string
	Email        string
	Enable       string
	RegisterDate string
	ImageHeader  string
	Following    string
	Follower     string
	Biography    string
} {
	return struct {
		ID           string
		Username     string
		Name         string
		Family       string
		Image        string
		Password     string
		Mobile       string
		Email        string
		Enable       string
		RegisterDate string
		ImageHeader  string
		Following    string
		Follower     string
		Biography    string
	}{
		"_id",
		"username",
		"name",
		"family",
		"image",
		"password",
		"mobile",
		"email",
		"enable",
		"register_date",
		"image_header",
		"following",
		"follower",
		"biography",
	}
}

// NewMember returns a new member
func NewMember(
	username string,
	password string,
	name string,
	family string,
	image string,
	mobile string,
	email string,
	enable bool,
	registerDate time.Time,
) (*Member, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	member := &Member{
		Password: string(hashedPassword),
	}

	return member, nil
}

// GetToken returns the Member's JWT
func (member Member) GetToken() string {
	byteSlc, _ := json.Marshal(member)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": string(byteSlc),
	})
	tokenString, _ := token.SignedString([]byte(constant.JwtSecret))
	return tokenString
}

// MemberFromToken returns the Member which is authenticated with this Token
func MemberFromToken(token string) Member {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.JwtSecret), nil
	})
	var result Member
	json.Unmarshal([]byte(claims["data"].(string)), &result)
	return result
}

// IsCorrectPassword checks if the provided password is correct or not
func (member *Member) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(password))
	return err == nil
}
