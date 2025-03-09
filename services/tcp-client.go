package services

import (
	"fmt"
	"log"
	"modbus-server/config"
	"net"
	"time"
)

type TcpClient struct {
	Conn net.Conn
}

var tcpClient *TcpClient

func GetOrCreateTcpClient() *TcpClient {
	if tcpClient == nil {
		conn := RetryConnect(5)
		tcpClient = &TcpClient{
			Conn: conn,
		}
	}
	return tcpClient
}

func (client *TcpClient) Close() {
	client.Conn.Close()
}

func (client *TcpClient) Write(data []byte) {
	client.Conn.Write(data)
}

func (client *TcpClient) Read(data []byte) {
	client.Conn.Read(data)
}

func RetryConnect(count int) net.Conn {
	var conn net.Conn
	var err error

	address := fmt.Sprintf("%s:%d", config.Cfg.ModbusHost, config.Cfg.ModbusPort)

	for i := 0; i < count; i++ {
		conn, err = net.Dial("tcp", address)
		if err == nil {
			log.Println("Успешное подключение к", address)
			return conn
		}

		log.Printf("Ошибка подключения (%d/%d): %v", i+1, count, err)
		time.Sleep(1 * time.Second)
	}

	panic(fmt.Sprintf("Не удалось подключиться к %s после %d попыток", address, count))
}
