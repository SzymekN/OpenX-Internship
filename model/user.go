package model

import (
	"fmt"
	"math"
	"strconv"
)

type User struct {
	Address  `json:"address"`
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     `json:"name"`
	Phone    string `json:"phone"`
	V        int    `json:"__v"`
}

type Users []User

type Address struct {
	Geolocation `json:"geolocation"`
	City        string `json:"city"`
	Street      string `json:"street"`
	Number      int    `json:"number"`
	Zipcode     string `json:"zipcode"`
}

type Geolocation struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

func (users Users) GetByID(id int) User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return User{}
}

func convertCoords(coordsStr []string) ([]float64, error) {
	result := []float64{}
	for _, coord := range coordsStr {
		v, err := strconv.ParseFloat(coord, 64)

		if err != nil {
			return nil, err
		}

		vRad := v * math.Pi / 180
		result = append(result, vRad)
	}
	return result, nil
}

func (users Users) FindFurthestAwayLiving() Users {

	if len(users) < 2 {
		return Users{}
	}

	maxDist := -1.0
	u1Index, u2Index := 0, 0
	earthRadius := 6371.0

	for i := range users {
		for j := i + 1; j < len(users); j++ {
			coordsString := []string{users[i].Address.Geolocation.Lat, users[j].Address.Geolocation.Lat, users[i].Address.Geolocation.Long, users[j].Address.Geolocation.Long}
			coordsRads, err := convertCoords(coordsString)

			if err != nil {
				fmt.Println("ERROR while converting coord of user ", users[i])
				fmt.Println("and ", users[j])
				continue
			}

			// Haversine formula - distance on a sphere
			distance := 2 * earthRadius * math.Asin(math.Sqrt(math.Pow(math.Sin((coordsRads[1]-coordsRads[0])/2), 2)+
				math.Cos(coordsRads[0])*math.Cos(coordsRads[1])*math.Pow(math.Sin((coordsRads[3]-coordsRads[2])/2), 2)))

			if distance > maxDist {
				maxDist = distance
				u1Index = i
				u2Index = j

			}
		}

	}

	return Users{users[u1Index], users[u2Index]}

}
