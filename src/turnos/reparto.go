package turnos

import "time"

type RepartoDeTareas struct {
	fecha           time.Time
	idsTrabajadores map[string]struct{}
	huecos          []Hueco
}
