package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//func main() {
//	//rate, err := GetAtCoderRate("tourist")
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//fmt.Println(rate)
//	//data, _ := GetUserJSONData("yuji9511")
//	//fmt.Println(data)
//	//
//	//color, _ := GetUserColor("tourist")
//	//fmt.Println(color)
//
//	data, _ := GetUser("yuji9511")
//	fmt.Println(data)
//}

func GetUser(userID string) (AtCoderUser, error) {
	var result AtCoderUser
	rating, err := GetAtCoderRate(userID)
	if err != nil {
		return result, err
	}
	color, err := GetUserColor(userID)
	if err != nil {
		return result, err
	}
	jsonData, err := GetUserJSONData(userID)
	if err != nil {
		return result, err
	}
	result.Color = color
	result.Rating = rating
	result.Details = jsonData
	return result, nil
}

// GetAtCoderRate returns current rating
func GetAtCoderRate(userID string) (int, error) {
	if userID == "" {
		return -1, errors.New("userID is blank")
	}
	url := "https://atcoder.jp/users/" + userID
	out, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	if out.StatusCode != 200 {
		return -1, errors.New("User: " + userID + " doesn't exist")
	}

	jsonURL := "https://atcoder.jp/users/" + userID + "/history/json"
	res, err := http.Get(jsonURL)
	if err != nil {
		return -1, err
	}
	defer res.Body.Close()
	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	bytes := []byte(byteArray)
	var result []AtCoderRating
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return -1, err
	}
	if len(result) == 0 {
		return 0, nil
	}
	rating := result[len(result)-1].NewRating

	return rating, nil

}

// GetUserJSONData returns all user data at https://atcoder.jp/users/{user_id}/history/json
func GetUserJSONData(userID string) ([]AtCoderRating, error) {
	var data []AtCoderRating
	if userID == "" {
		return data, errors.New("userID is blank")
	}
	url := "https://atcoder.jp/users/" + userID
	out, err := http.Get(url)
	if err != nil {
		return data, err
	}
	if out.StatusCode != 200 {
		return data, errors.New("User: " + userID + " doesn't exist")
	}

	jsonURL := "https://atcoder.jp/users/" + userID + "/history/json"
	res, err := http.Get(jsonURL)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()
	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	bytes := []byte(byteArray)
	var result []AtCoderRating
	json.Unmarshal(bytes, &result)
	return result, nil
}

func GetUserColor(userID string) (string , error) {
	rate, err := GetAtCoderRate(userID)
	if err != nil {
		return "", err
	}
	var color string
	switch {
	case rate <= 0:
		color = "#000000"
	case rate <= 399:
		color = "#808080"
	case rate <= 799:
		color = "#804000"
	case rate <= 1199:
		color = "#008000"
	case rate <= 1599:
		color = "#00C0C0"
	case rate <= 1999:
		color = "#0000FF"
	case rate <= 2399:
		color = "#C0C000"
	case rate <= 2799:
		color = "#FF7F01"
	default:
		color = "#FF0000"
	}
	return color, nil
}
