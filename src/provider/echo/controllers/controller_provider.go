package controllers

type ControllerProvider struct {
	ConverterControllerProvider *ConverterControllerProvider
}

func NewControllerProvider(
	ConverterControllerProvider *ConverterControllerProvider,
) *ControllerProvider {
	return &ControllerProvider{
		ConverterControllerProvider,
	}
}
