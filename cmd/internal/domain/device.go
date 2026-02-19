package domain

import "time"

type Device struct {
	MAC           string
	IP            string
	Hostname      string
	ThroughputBPS uint64
	LastSeen      time.Time
	IsOnline      bool
}

type DeviceRepository interface {
	Upsert(device *Device) error // upsert = atualiza se existe, se nao cria
	GetAll() ([]*Device, error)  // retorna lista de ponteiros de Devices
	FindByIP(ip string) (*Device, error)
	FindByMAC(mac string) (*Device, error)
}
