package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ModbusPort int
	ModbusHost string
	DBPath     string
	DBName     string
}

var Cfg Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	modbusPortStr := os.Getenv("MODBUS_PORT")
	Cfg.ModbusPort, err = strconv.Atoi(modbusPortStr)
	if err != nil {
		Cfg.ModbusPort = 1502
	}
	Cfg.DBPath = os.Getenv("DB_PATH")
	Cfg.DBName = os.Getenv("DB_NAME")
	Cfg.ModbusHost = os.Getenv("MODBUS_HOST")
	switch {
	case Cfg.DBPath == "":
		Cfg.DBPath = "../db/"
	case Cfg.DBName == "":
		Cfg.DBName = "modbus.db"
	case Cfg.ModbusHost == "":
		Cfg.ModbusHost = "127.0.0.1"
	default:
	}
}
