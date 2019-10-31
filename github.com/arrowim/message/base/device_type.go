package base

type DeviceType int16

func (self DeviceType) GetValue() int16 {
	return int16(self)
}

const (
	DEVICE_MOBILE DeviceType = 1
	DEVICE_PAD    DeviceType = 2
	DEVICE_PC     DeviceType = 4
)
