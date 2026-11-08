package turnos

type TrabajadorID string

type Trabajador struct {
	nombre           string
	disponibilidad   []Intervalo
	tareasPermitidas map[Tarea]struct{}
}
