package turnos

type Plantilla struct {
	trabajadores map[*Trabajador]struct{}
}
