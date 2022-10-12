package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
	"github.com/vearutop/statigz"
)

//go:embed godot-docs-html/*
var static embed.FS

//go:embed icon.png
var iconPNG []byte

//go:embed version.txt
var versionFile []byte

func init() {
	gtk.Init()
}

func main() {
	allArgs := os.Args[1:]
	host := "localhost"
	port := "2300"

	for _, element := range allArgs {
		if _, err := strconv.Atoi(element); err != nil {
			log.Println(errors.New("invalid PORT"))
		} else {
			port = element
		}
	}

	port = getOpenPort(host, port)

	version := string(versionFile)
	// remove the newline
	version = version[:len(version)-1]

	icon, err := loadPNG(iconPNG)
	if err != nil {
		log.Fatalln("failed to load icon.png:", err)
	}

	label := gtk.NewLabel("Serving on " + host + ":" + port)
	label.Show()

	hbox := gtk.NewBox(gtk.OrientationVertical, 6)
	hbox.Show()

	hbox.PackStart(label, false, false, 6)

	button := gtk.NewButtonWithLabel("Open Browser")
	button.ConnectClicked(func() {
		openBrowser("http://" + host + ":" + port)
	})

	hbox.PackStart(button, false, false, 6)
	button.Show()

	window := gtk.NewWindow(gtk.WindowToplevel)
	window.ConnectDestroy(gtk.MainQuit)
	window.SetTitle("Godot " + version + " Docs")
	window.SetIcon(icon)
	window.Add(hbox)
	window.SetDefaultSize(320, 240)
	window.Show()

	button2 := gtk.NewButtonWithLabel("Hide Window")
	button2.ConnectClicked(func() {
		window.Hide()
	})

	hbox.PackStart(button2, false, false, 6)
	button2.Show()

	go httpServer(host, port)

	gtk.Main()
}

func loadPNG(data []byte) (*gdkpixbuf.Pixbuf, error) {
	l, err := gdkpixbuf.NewPixbufLoaderWithType("png")
	if err != nil {
		return nil, fmt.Errorf("NewLoaderWithType png: %w", err)
	}
	defer l.Close()

	if err := l.Write(data); err != nil {
		return nil, fmt.Errorf("PixbufLoader.Write: %w", err)
	}

	if err := l.Close(); err != nil {
		return nil, fmt.Errorf("PixbufLoader.Close: %w", err)
	}

	return l.Pixbuf(), nil
}

func getOpenPort(host string, port string) string {
	listener, err := net.Listen("tcp", host + ":" + port)
	if err == nil {
		listener.Close()
	}
	for err != nil {
		newPort, _ := strconv.Atoi(port)
		newPort++
		port = strconv.Itoa(newPort)
		listener, err := net.Listen("tcp", host + ":" + port)
		if err == nil {
			listener.Close()
			break
		}
	}
	return port
}

func httpServer(host string, port string) {
	contentStatic, _ := fs.Sub(static, "godot-docs-html")
    http.Handle("/", statigz.FileServer(contentStatic.(fs.ReadDirFS)))
	println("Serving on port: " + port)
    log.Fatal(http.ListenAndServe(host + ":" + port, nil))
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
