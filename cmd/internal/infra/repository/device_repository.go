package repository

import (
	"errors"
	"sync"

	"github.com/lhupalo/mcp-network-sentinel-go/cmd/internal/domain"
)

type MemoryDeviceRepository struct {
	devicesList  map[string]*domain.Device
	sync.RWMutex // herda as capacidades de um mutex
}

func NewMemoryDeviceRepository() *MemoryDeviceRepository {
	return &MemoryDeviceRepository{devicesList: make(map[string]*domain.Device, 100)}
}

func (m *MemoryDeviceRepository) Upsert(newDevice *domain.Device) error {

	if newDevice.MAC == "" {
		return errors.New("MAC address is required")
	}

	m.Lock() // bloqueia para escrita
	defer m.Unlock()

	m.devicesList[newDevice.MAC] = newDevice

	return nil
}

func (m *MemoryDeviceRepository) FindByMAC(mac string) (*domain.Device, error) {

	m.RLock()         // bloqueia para leitura
	defer m.RUnlock() // vai executar esse defer quando chegar no return

	device, found := m.devicesList[mac]

	if found == false {
		return nil, errors.New("device not found")
	}

	return device, nil
}

func (m *MemoryDeviceRepository) FindByIP(ip string) (*domain.Device, error) {

	m.RLock()         // bloqueia para leitura
	defer m.RUnlock() // vai executar esse defer quando chegar no return

	devices := m.devicesList

	for _, device := range devices {
		if ip == device.IP {
			return device, nil
		}
	}

	return nil, errors.New("device not found")
}

func (m *MemoryDeviceRepository) GetAll() ([]*domain.Device, error) {

	m.RLock()
	defer m.RUnlock()

	devices := make([]*domain.Device, 0, len(m.devicesList))

	for _, device := range m.devicesList {
		devices = append(devices, device)
	}

	return devices, nil
}
