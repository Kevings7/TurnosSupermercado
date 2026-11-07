package turnos

type Trabajador struct {
	nombre           string
	disponibilidad   []Intervalo
	tareasPermitidas map[string]struct{}
}
