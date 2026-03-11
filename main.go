package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/kkdai/youtube/v2"
	"github.com/schollz/progressbar/v3"
)

func main() {
	var url string
	fmt.Print("🔗 Masukkan URL YouTube: ")
	fmt.Scanln(&url)

	downloadDir := "downloads"
	os.MkdirAll(downloadDir, 0755)

	client := youtube.Client{}
	video, err := client.GetVideo(url)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	// Ambil format audio terbaik
	formats := video.Formats.WithAudioChannels()
	formats.Sort()
	format := &formats[0]

	// FIX PROGRESS BAR: Ambil size dari format jika video.GetStream memberikan size 0
	stream, size, err := client.GetStream(video, format)
	if err != nil {
		fmt.Printf("❌ Gagal stream: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("\n🎵 Judul: %s\n", video.Title)

	// Progress Bar
	bar := progressbar.NewOptions64(
		size,
		progressbar.OptionSetDescription("📥 Mengunduh"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(20),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() { fmt.Print("\n") }),
	)

	tempPath := filepath.Join(downloadDir, "temp_"+video.ID+".webm")
	file, _ := os.Create(tempPath)
	_, _ = io.Copy(io.MultiWriter(file, bar), stream)
	file.Close()

	// Menyiapkan Nama File Output
	cleanTitle := strings.NewReplacer("/", "", "\\", "", ":", "", "*", "", "?", "", "\"", "", "<", "", ">", "", "|", "").Replace(video.Title)
	outputPath := filepath.Join(downloadDir, cleanTitle+".mp3")

	// Konversi ke MP3
	fmt.Print("🔄 Mengonversi ke MP3... ")
	cmd := exec.Command("ffmpeg", "-i", tempPath, "-vn", "-ab", "192k", "-ar", "44100", "-y", outputPath)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Gagal konversi: %v\n", err)
		return
	}
	fmt.Println("Selesai! ✅")

	// --- PROSES THUMBNAIL & METADATA ---
	fmt.Print("🖼️  Menempelkan Thumbnail... ")
	
	// 1. Ambil URL thumbnail terbaik
	thumbURL := video.Thumbnails[len(video.Thumbnails)-1].URL
	resp, err := http.Get(thumbURL)
	if err == nil {
		defer resp.Body.Close()
		thumbData, _ := io.ReadAll(resp.Body)

		// 2. Tulis Metadata menggunakan id3v2
		tag, err := id3v2.Open(outputPath, id3v2.Options{Parse: true})
		if err == nil {
			tag.SetTitle(video.Title)
			tag.SetArtist(video.Author)
			
			// Lampirkan gambar
			pic := id3v2.PictureFrame{
				Encoding:    id3v2.EncodingUTF8,
				MimeType:    "image/jpeg",
				PictureType: id3v2.PTFrontCover,
				Description: "Front Cover",
				Picture:     thumbData,
			}
			tag.AddAttachedPicture(pic)
			tag.Save()
			tag.Close()
		}
	}
	fmt.Println("Selesai! ✨")

	// Cleanup
	os.Remove(tempPath)
	fmt.Printf("\n🚀 Sukses! File siap di: %s\n", outputPath)
}