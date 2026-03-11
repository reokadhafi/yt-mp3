# 🎵 YouTube to MP3 Downloader (Golang Version)

Aplikasi CLI sederhana namun kuat yang dibuat dengan **Go** untuk mengunduh audio YouTube dan mengonversinya langsung ke format **MP3** dengan kualitas tinggi (192kbps), lengkap dengan **Progress Bar** dan **Thumbnail Embedding**.

## ✨ Fitur
- ✅ **High Quality Audio:** Ekstraksi audio terbaik dengan bitrate 192kbps.
- ✅ **Progress Bar:** Indikator unduhan yang informatif dan real-time.
- ✅ **Auto Metadata:** Otomatis mengisi judul lagu dan nama artis.
- ✅ **Thumbnail Album Art:** Menyisipkan gambar thumbnail video ke dalam file MP3.
- ✅ **Clean Filename:** Menghapus karakter ilegal agar kompatibel di semua OS.

## 🛠️ Prasyarat
Sebelum menjalankan aplikasi ini, pastikan Anda telah menginstal:
1. [Go](https://go.dev/doc/install) (versi 1.18 atau terbaru)
2. [FFmpeg](https://ffmpeg.org/download.html) (Wajib ada di PATH sistem Anda)

### Instalasi FFmpeg (Termux)
```bash
pkg install ffmpeg
```
### Instalasi ytmp3 (Termux)
```bash
https://github.com/reokadhafi/yt-mp3.git
go mod init yt-mp3
go mod tidy
```
### run ytmp3 (Termux)
```bash
go run main.go
