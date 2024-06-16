package main

import (
	"fmt"
	"github.com/zergon321/reisen"
)

func main() {
	media, err := reisen.NewMedia("videos/XFiles.mp4")
	handleError(err)
	defer media.Close()

	for _, stream := range media.Streams() {
		dur, err := stream.Duration()
		handleError(err)
		tbNum, tbDen := stream.TimeBase()
		fpsNum, fpsDen := stream.FrameRate()

		//Print the properites common
		//for both stresm types

		fmt.Println("Index:", stream.Index())
		fmt.Println("Stream Type:", stream.Type())
		fmt.Println("Codec Name", stream.CodecName())
		fmt.Println("Codec long name", stream.CodecLongName())
		fmt.Println("Stream Duration:", dur)
		fmt.Println("Stream Bitrate", stream.BitRate())
		fmt.Printf("Time base: %d%d\n", tbNum, tbDen)
		fmt.Printf("Frame rate: %d%d\n", fpsNum, fpsDen)
		fmt.Println("Frame count", stream.FrameCount())
		fmt.Print("\n")

	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
