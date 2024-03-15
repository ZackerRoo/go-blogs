package service

import (
	"go-blogs/common/global"
	"go-blogs/models"
)

type ChannelService struct {
}

// 添加
func (u *ChannelService) AddChannel(channel models.Channel) int64 {
	return global.Db.Table("channel").Create(&channel).RowsAffected
}

// 删除
func (u *ChannelService) DelChannel(id int) int64 {
	return global.Db.Delete(&models.Channel{}, id).RowsAffected
}

// 修改
func (u *ChannelService) UpdateChannel(channel models.Channel) int64 {
	return global.Db.Updates(&channel).RowsAffected
}

// get
func (u *ChannelService) GetChannel(id int) models.Channel {
	var channel models.Channel
	global.Db.First(&channel, id)
	return channel
}

// get channel list
func (u *ChannelService) GetChannelList() []models.Channel {
	channelList := make([]models.Channel, 0)
	global.Db.Find(&channelList)
	return channelList
}
