package turnos

type TrabajadorID string

type Trabajador struct {
	Nombre           string
	Disponibilidad   []Intervalo
	TareasPermitidas map[Tarea]struct{}
}
