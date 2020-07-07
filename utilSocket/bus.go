package utilSocket

import (
)

type Bus struct {
  channels map[uint]Channel
}

func NewBus() *Bus {
	return &Bus{
		channels: make(map[uint]Channel),
	}
}

func (bus *Bus) Register(channelId uint, channel Channel){
  bus.channels[channelId] = channel
  go channel.Run()
}

func (bus *Bus) Delete(channel Channel){
  // delete(bus.channels, channel)
}

func (bus *Bus) At(channelId uint) Channel {
  return bus.channels[channelId]
}
