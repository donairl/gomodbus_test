package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/goburrow/modbus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*
func mainr() {
	// Modbus TCP

	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x01
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)

	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)
	results, err := client.ReadDiscreteInputs(1, 4)
	fmt.Println(results)
	fmt.Println(err)
	results, err = client.WriteSingleRegister(2, 88)
	//	results, err = client.WriteMultipleRegisters(1, 1, []byte{2})
	results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})

}
*/
//konversi fake float16 as float32
func Float16frombytes(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)

	floatx := math.Float32frombits(bits)

	return floatx
}

//baca
func ReadInput(c echo.Context) error {
	id := c.Param("id")

	handler := modbus.NewTCPClientHandler("10.0.0.1:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x01
	handler.Logger = log.New(os.Stdout, "data comm: ", log.LstdFlags)

	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()

	defer handler.Close()

	client := modbus.NewClient(handler)
	value, err := strconv.ParseUint(id, 10, 16)
	results, err := client.ReadInputRegisters(uint16(value), 2)

	if err != nil {
		log.Println("Error returned : ", err)
	}
	fmt.Println(value)
	fmt.Println(results)
	//angka := int(results[1])
	//angka2 := int(results[0])
	//number := (angka2 << 8) + angka

	//return c.String(http.StatusOK, s)
	//prefix := make([]byte, 2)

	//gval := append(prefix[:], results[:]...)
	//fmt.Println(gval)
	number := Float16frombytes(results)
	fmt.Println(number)
	return c.String(http.StatusOK, "Y")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/about", func(c echo.Context) error {

		return c.String(http.StatusOK, "Modbus API Reader, (c)2019, Donny \n")
	})

	e.GET("/bacainput/:id", ReadInput)

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")

		return c.String(http.StatusOK, "Anda memanggil lewat /users/"+id)
	})
	e.Static("/", "web")
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
