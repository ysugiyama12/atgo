package atgo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// func main() {
// 	rate, err := GetAtCoderRate("Um_nik")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(rate)
// }

// GetAtCoderRate return current rating
func GetAtCoderRate(userID string) (int, error) {
	if userID == "" {
		return -1, errors.New("user_id is blank")
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
	json.Unmarshal(bytes, &result)
	if len(result) == 0 {
		return 0, nil
	}
	rating := result[len(result)-1].NewRating

	return rating, nil

}
