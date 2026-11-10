package turnos

type Tarea string

type Hueco struct {
	tarea        Tarea
	periodo      Intervalo
	aplazable    bool
	trabajadorID *TrabajadorID
}
