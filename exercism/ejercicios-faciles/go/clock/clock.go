package clock

import "fmt"

type Clock struct {
	hora    int
	minutos int
}

func New(h, m int) Clock {
	// Convertir todo a minutos totales
	totalMinutos := h*60 + m

	// Calcular horas y minutos
	horas := totalMinutos / 60
	minutos := totalMinutos % 60

	// Manejar minutos negativos
	if minutos < 0 {
		minutos += 60
		horas -= 1
	}

	// Normalizar horas a 0-23
	horas = ((horas % 24) + 24) % 24

	return Clock{hora: horas, minutos: minutos}
}

func (c Clock) Add(m int) Clock {

	return New(c.hora, c.minutos+m)
}

func (c Clock) Subtract(m int) Clock {

	return New(c.hora, c.minutos-m)
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", c.hora, c.minutos)

}
