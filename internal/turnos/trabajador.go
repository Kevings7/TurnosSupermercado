package turnos

type Trabajador struct {
	nombre           string
	disponibilidad   []Intervalo
	tareasPermitidas map[Tarea]struct{}
}
