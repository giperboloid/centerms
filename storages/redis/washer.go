package storages

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/giperboloid/centerms/entities"
	"github.com/pkg/errors"
)

func (rs *RedisDevStorage) getWasherData(m *entities.DevMeta) (*entities.DevData, error) {
	devKey := "device:" + m.Type + ":" + m.Name + ":" + m.MAC
	devParamsKey := devKey + ":" + "params"

	dd := entities.DevData{
		Meta: *m,
		Data: make(map[string][]string),
	}

	params, err := rs.Client.SMembers(devParamsKey)
	if err != nil {
		errors.Wrap(err, "RedisDevStorage: getWasherData(): can't read members from devParamsKeys")
	}

	data := make([][]string, len(params))
	for i, p := range params {
		data[i], err = rs.Client.ZRangeByScore(devParamsKey+":"+p, "-inf", "inf")
		if err != nil {
			errors.Wrap(err, "RedisDevStorage: getWasherData(): can't read members from sorted set")
		}
		dd.Data[p] = data[i]
	}

	return &dd, err
}

func (rs *RedisDevStorage) setWasherData(r *entities.Request) error {
	var wd entities.WasherData
	err := json.NewDecoder(bytes.NewBuffer(r.Data)).Decode(&wd)
	if err != nil {
		errors.Wrap(err, "RedisDevStorage: setWasherData(): WasherData decoding has failed")
		return err
	}

	devKey := "device" + ":" + r.Meta.Type + ":" + r.Meta.Name + ":" + r.Meta.MAC
	devParamsKey := devKey + ":" + "params"

	if _, err = rs.Client.Multi(); err != nil {
		errors.Wrap(err, "RedisDevStorage: setWasherData(): Multi() has failed")
		rs.Client.Discard()
		return err
	}
	if err = rs.setTurnoversData(wd.Turnovers, devParamsKey+":"+"Turnovers"); err != nil {
		errors.Wrap(err, "RedisDevStorage: setWasherData(): setTurnoversData() has failed")
		rs.Client.Discard()
		return err
	}
	if err = rs.setWaterTempData(wd.WaterTemp, devParamsKey+":"+"WaterTemp"); err != nil {
		errors.Wrap(err, "RedisDevStorage: setWasherData(): setWaterTempData() has failed")
		rs.Client.Discard()
		return err
	}
	if _, err = rs.Client.Exec(); err != nil {
		errors.Wrap(err, "RedisDevStorage: setWasherData(): Exec() has failed")
		rs.Client.Discard()
		return err
	}

	return nil
}

func (rs *RedisDevStorage) setTurnoversData(TempCam map[int64]int64, key string) error {
	for t, v := range TempCam {
		_, err := rs.Client.ZAdd(key, strconv.FormatInt(int64(t), 10),
			strconv.FormatInt(int64(t), 10)+":"+strconv.FormatInt(int64(v), 10))
		if err != nil {
			errors.Wrap(err, "RedisDevStorage: setTurnoversData(): adding to sorted set has failed")
			return err
		}
	}

	return nil
}

func (rs *RedisDevStorage) setWaterTempData(TempCam map[int64]float32, key string) error {
	for t, v := range TempCam {
		_, err := rs.Client.ZAdd(key, strconv.FormatInt(int64(t), 10),
			strconv.FormatInt(int64(t), 10)+":"+
				strconv.FormatFloat(float64(v), 'f', -1, 32))
		if err != nil {
			errors.Wrap(err, "RedisDevStorage: setWaterTempData(): adding to sorted set has failed")
			return err
		}
	}

	return nil
}

func (rs *RedisDevStorage) getWasherConfig(m *entities.DevMeta) (*entities.DevConfig, error) {
	config, err := rs.getWasherDefaultConfig(m)
	config.MAC = m.MAC
	configKey := m.MAC + ":config"
	unixTime:= int64(100) // fake
	mode, err := rs.Client.ZRangeByScore(configKey, unixTime-100, unixTime+100)
	if err != nil {
		errors.Wrap(err, "RedisDevStorage: getWasherConfig(): ZRangeByScore() has failed")
	}

	if len(mode) == 0 {
		return config, err
	}

	configWasher := entities.LightMode
	config.Data, err = json.Marshal(configWasher)
	if err != nil {
		errors.Wrap(err, "RedisDevStorage: getWasherConfig(): WasherConfig marshalling has failed")
	}
	return config, err
}

func (rs *RedisDevStorage) setWasherConfig(c *entities.DevConfig) error {
	var timerMode *entities.TimerMode
	err := json.NewDecoder(bytes.NewBuffer(c.Data)).Decode(&timerMode)
	if err != nil {
		errors.Wrap(err, "RedisDevStorage: setWasherConfig(): Decode() has failed")
	}

	configKey := c.MAC + ":config"
	rs.Client.ZAdd(configKey, timerMode.StartTime, timerMode.Name)
	if err != nil {
		errors.Wrap(err, "RedisDevStorage: setWasherConfig(): ZAdd() has failed")
	}

	return err
}

func (rs *RedisDevStorage) getWasherDefaultConfig(m *entities.DevMeta) (*entities.DevConfig, error) {
	b, err := json.Marshal(entities.StandardMode)
	if err != nil {
		errors.Wrap(err, "RedisDevStorage: getWasherDefaultConfig(): WasherConfig marshalling has failed")
		return nil, err
	}

	return &entities.DevConfig{
		MAC:  m.MAC,
		Data: b,
	}, nil
}
