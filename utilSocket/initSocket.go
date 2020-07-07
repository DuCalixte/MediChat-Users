package utilSocket

import ()

var BUS *Bus

func InitBus() {
  BUS = NewBus()
}

func AddChannel(channelName string, channelId uint, channelType int) {
  channel := NewChannel(channelName, channelId, channelType)
  BUS.Register(channelId, *channel)
}

func AddChatBotChannel( channelId uint) {
  channel := NewChannel("ChatBot", channelId, 2)
  BUS.Register(channelId, *channel)
}
