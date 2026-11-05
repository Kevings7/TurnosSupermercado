package turnos

import "time"

type RepartoDeTareas struct {
	fecha               time.Time
	huecosPorTrabajador map[string][]Hueco
	huecosSinCubrir     []Hueco
}
