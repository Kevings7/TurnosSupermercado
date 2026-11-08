package turnos

type Tarea string

type Hueco struct {
	tarea        Tarea
	intervalo    Intervalo
	aplazable    bool
	trabajadorID TrabajadorID
}
