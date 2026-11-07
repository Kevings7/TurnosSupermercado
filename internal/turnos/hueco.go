package turnos

type Tarea string

type Hueco struct {
	Tarea     Tarea
	Intervalo Intervalo
	Aplazable bool
}
