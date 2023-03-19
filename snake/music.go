package snake

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type Music struct {
	context *audio.Context
	player  *audio.Player
	on      bool
}

func NewMusic(sampleRate int) *Music {
	m := &Music{
		context: audio.NewContext(sampleRate),
		player:  nil,
		on:      true,
	}
	return m
}

func (m *Music) Play(name string) {
	f, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	stream, err := mp3.DecodeWithoutResampling(bytes.NewReader(f)) // 解码
	if err != nil {
		log.Fatal(err)
	}
	// 创建循环
	// stream = audio.NewInfiniteLoop(stream, stream.Length())
	m.player, err = m.context.NewPlayer(stream)
	if err != nil {
		log.Fatal(err)
	}
}
