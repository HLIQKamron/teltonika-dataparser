package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Projects/teltonika-dataparser/db/redis"
	"github.com/Projects/teltonika-dataparser/models"
)

func Update(d models.Device) error {

	fmt.Println("Updating data in redis :", d)
	var data = make(map[string]models.Device)

	val, err := redis.RDB.Get(context.TODO(), "Devices").Result()
	if err != nil {
		log.Println("Error in getting data from redis", err)
		return err
	}
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		log.Println("Error in unmarshalling data")
		return err
	}
	data[d.IMEI] = d
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error in marshalling data")
		return err
	}
	return redis.RDB.Set(context.TODO(), "Devices", jsonData, 0).Err()
}

func GetDeviceInfoByImei(imei string) (models.Device, error) {
	var data = make(map[string]models.Device)
	val, err := redis.RDB.Get(context.TODO(), "Devices").Result()
	if err != nil {
		log.Println("Error in getting data from redis", err)
		return models.Device{}, err
	}
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		log.Println("Error in unmarshalling data")
		return models.Device{}, err
	}
	return data[imei], nil
}
