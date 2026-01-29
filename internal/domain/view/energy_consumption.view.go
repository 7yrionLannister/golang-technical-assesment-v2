package view

type EnergyConsumption struct {
	MeterId          uint8
	Address          string // TODO consume external API
	TotalConsumption float32
}
