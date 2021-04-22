////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2021-04-22 10:13
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
)
// 适配器用于接口不一致时的接口装换

// Step 1
// 为媒体播放器和更高级的媒体播放器创建接口。
type MediaPlayer interface {
	Play(audioType, fileName string)
}

type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

// Step 2
// 创建实现了 AdvancedMediaPlayer 接口的实体类。
type VlcPlayer struct{}

func (p *VlcPlayer) PlayVlc(fileName string) {
	fmt.Println("Playing vlc file. Name: " + fileName)
}

func (p *VlcPlayer) PlayMp4(fileName string) {

}

type Mp4Player struct{}

func (p *Mp4Player) PlayVlc(fileName string) {
}

func (p *Mp4Player) PlayMp4(fileName string) {
	fmt.Println("Playing mp4 file. Name: " + fileName)
}

// Step 3
// 创建实现了 MediaPlayer 接口的适配器类。
type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	var advancedMusicPlayer AdvancedMediaPlayer
	if strings.ToLower(audioType) == "vlc" {
		advancedMusicPlayer = &VlcPlayer{}
	} else if strings.ToLower(audioType) == "mp4" {
		advancedMusicPlayer = &Mp4Player{}
	}
	return &MediaAdapter{advancedMusicPlayer: advancedMusicPlayer}
}

func (p *MediaAdapter) Play(audioType, fileName string) {
	if strings.ToLower(audioType) == "vlc" {
		p.advancedMusicPlayer.PlayVlc(fileName)
	} else if strings.ToLower(audioType) == "mp4" {
		p.advancedMusicPlayer.PlayMp4(fileName)
	}
}

// Step 4
// 创建实现了 MediaPlayer 接口的实体类。
type AudioPlayer struct{}

func (p *AudioPlayer) Play(audioType, fileName string) {
	if strings.ToLower(audioType) == "mp3" {
		fmt.Println("Playing mp3 file. Name: " + fileName)
	} else if strings.ToLower(audioType) == "vlc" || strings.ToLower(audioType) == "mp4" {
		mediaAdapter := NewMediaAdapter(audioType)
		mediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}
}

func main() {
	audioPlayer := &AudioPlayer{}
	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}