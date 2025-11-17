package api

type TransmutationRequestDto struct {
	AlquimistaID uint   `json:"alquimista_id"`
	MaterialID   uint   `json:"material_id"`
	Costo        int    `json:"costo"`
	Resultado    string `json:"resultado"`
	Estado       string `json:"estado"`
}

type TransmutationResponseDto struct {
	ID           int    `json:"id"`
	AlquimistaID uint   `json:"alquimista_id"`
	MaterialID   uint   `json:"material_id"`
	Costo        int    `json:"costo"`
	Resultado    string `json:"resultado"`
	Estado       string `json:"estado"`
	FechaCreacion string `json:"fecha_creacion"`
}
