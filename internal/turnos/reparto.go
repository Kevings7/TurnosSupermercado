package turnos

import "time"

type RepartoDeTareas struct {
	fecha                 time.Time
	huecosPorTrabajadorID map[string][]Hueco
	huecosSinCubrir       []Hueco
}
