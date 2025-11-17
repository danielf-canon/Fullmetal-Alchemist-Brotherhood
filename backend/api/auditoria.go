package api

type AuditoriaResponseDto struct {
	ID            int    `json:"id"`
	User          string `json:"user"`
	Accion        string `json:"accion"`
	Entidad       string `json:"entidad"`
	Descripcion   string `json:"descripcion"`
	FechaCreacion string `json:"fecha_creacion"`
}
