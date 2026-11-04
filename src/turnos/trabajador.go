package turnos

type Trabajador struct {
	id               string
	nombre           string
	disponibilidad   []Intervalo
	tareasPermitidas map[string]struct{}
}
