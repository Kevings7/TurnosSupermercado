package turnos

import "time"

type RepartoDeTareas struct {
	fecha                 time.Time
	huecosPorTrabajadorID map[TrabajadorID][]Hueco
	huecosSinCubrir       []Hueco
}
