package devices

import (
	"fmt"
	"log"
	"strings"

	"github.com/tarm/serial"
)

//Gps : Estructura que contiene datos del gps
type Gps struct {
	Name     string
	Device   string
	Baudrate int
	File     string
	port     *serial.Port

	Hour   int
	Minute int
	Sec    int
}

//GetName function
func (gps *Gps) GetName() string {
	return gps.Name
}

//GetFilePath function
func (gps *Gps) GetFilePath() string {
	return gps.File
}

//Init : Inicializa
func (gps *Gps) Init() error {
	var err error
	c := &serial.Config{Name: gps.Device, Baud: gps.Baudrate}
	gps.port, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
		return err
	}
	go gps.reading()
	return err
}

//Reading : Rutina para leer constantemente los datos del Gps
func (gps *Gps) reading() {
	var nmeaString string
	for {
		_, err := fmt.Fscanln(gps.port, &nmeaString)
		if err != nil {
			log.Println("Error Escaneo:", err)
			nmeaString = ""
			continue
		}

		if strings.Contains(nmeaString, "$GPRMC") {
			_, err = fmt.Sscanf(nmeaString, "$GPRMC,%2d%2d%2d.", &gps.Hour, &gps.Minute, &gps.Sec)
			if err != nil {
				log.Println("Error GPRMC:", err)
				nmeaString = ""
				continue
			}
		}

	}
}
