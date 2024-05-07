package dto

type ServiceDto struct {
	Data     any          `json:"data"`
	Metadata []Hypermedia `json:"_metadata"`
}

func NewServiceDto(data any) *ServiceDto {
	var dto = &ServiceDto{}
	dto.Data = data
	return dto
}
