package api

type AlquimistaRequestDto struct {
	Nombre       string `json:"nombre"`
	Edad         int32  `json:"edad"`
	Especialidad string `json:"especialidad"`
	Rango        string `json:"rango"`
}

type AlquimistaResponseDto struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Edad         int    `json:"edad"`
	Especialidad string `json:"especialidad"`
	Rango        string `json:"rango"`
	FechaCreacion string `json:"fecha_creacion"`
}

type ErrorResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Message     string `json:"message"`
}