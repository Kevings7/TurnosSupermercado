package turnos

import "time"

type RepartoDeTareas struct {
	Fecha                 time.Time
	HuecosPorTrabajadorID map[TrabajadorID][]Hueco
	HuecosSinCubrir       []Hueco
}
