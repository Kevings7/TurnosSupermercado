package turnos

import "time"

type RepartoDeTareas struct {
	fecha        time.Time
	trabajadores map[string]*Trabajador
	huecos       []Hueco
}
