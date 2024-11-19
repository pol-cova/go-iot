package models

import "time"

type SensorData struct {
	ID        uint      `json:"id"`
	DeviceMAC string    `json:"device_mac"`
	Timestamp time.Time `json:"timestamp"`

	// Water level data
	WaterLevelRaw     float64 `json:"waterLevelRaw"`
	WaterLevelPercent float64 `json:"waterLevelPercent"`
	// Soil moisture data
	SoilMoistureRaw     float64 `json:"soilMoistureRaw"`
	SoilMoisturePercent float64 `json:"soilMoisturePercent"`
	SoilDry             bool    `json:"soilDry"`
	// DHT sensor data
	AmbientTemperature float64 `json:"ambientTemperature"`
	RelativeHumidity   float64 `json:"relativeHumidity"`
	HeatIndex          float64 `json:"heatIndex"`
	// BMP sensor data
	AtmosphericPressure float64 `json:"atmosphericPressure"`
	PressureMmHg        float64 `json:"pressureMmHg"`
	Altitude            float64 `json:"altitude"`
	BmpTemperature      float64 `json:"bmpTemperature"`
	SeaLevelPressure    float64 `json:"seaLevelPressure"`
	RealAltitude        float64 `json:"realAltitude"`
}

type SensorDataResponse struct {
	Data    []SensorData `json:"data"`
	Total   int64        `json:"total"`
	Page    int          `json:"page"`
	PerPage int          `json:"per_page"`
}

type SensorStats struct {
	DeviceMac   string    `json:"device_mac"`
	LastReading time.Time `json:"last_reading"`

	// Avg values
	AvgWaterLevel   float64 `json:"avg_water_level"`
	AvgSoilMoisture float64 `json:"avg_soil_moisture"`
	AvgTemperature  float64 `json:"avg_temperature"`
	AvgHumidity     float64 `json:"avg_humidity"`
	AvgPressure     float64 `json:"avg_pressure"`

	// Max and min values
	MaxTemperature float64 `json:"max_temperature"`
	MinTemperature float64 `json:"min_temperature"`
	MaxHumidity    float64 `json:"max_humidity"`
	MinHumidity    float64 `json:"min_humidity"`

	// Alerts
	SoilDryCount  int `json:"soil_dry_count"`
	LowWaterCount int `json:"low_water_count"`
}

type QueryParams struct {
	DeviceID  string    `form:"device_id"`
	StartDate time.Time `form:"start_date"`
	EndDate   time.Time `form:"end_date"`
	Page      int       `form:"page,default=1"`
	PerPage   int       `form:"per_page,default=50"`
}
