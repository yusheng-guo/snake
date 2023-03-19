package snake

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const sampleRate = 48000

type Music struct {
	player *audio.Player
}

func NewMusic(name string) *Music {
	audioContext := audio.NewContext(sampleRate)
	f, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	stream, err := mp3.DecodeWithoutResampling(bytes.NewReader(f)) // 解码
	if err != nil {
		log.Fatal(err)
	}
	// 创建循环
	s := audio.NewInfiniteLoop(stream, stream.Length())
	player, err := audioContext.NewPlayer(s)
	if err != nil {
		log.Fatal(err)
	}
	return &Music{
		player: player,
	}
}
