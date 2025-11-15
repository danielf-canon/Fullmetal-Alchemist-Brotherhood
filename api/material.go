package api

type MaterialRequestDto struct {
	NombreMaterial string `json:"nombre_material"`
}

type MaterialResponseDto struct {
	ID             int    `json:"id"`
	NombreMaterial string `json:"nombre_material"`
	FechaCreacion  string `json:"fecha_creacion"`
}
