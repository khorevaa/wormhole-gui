package transport

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"fyne.io/fyne"
	"github.com/psanford/wormhole-william/wormhole"
)

// NewFileSend takes the chosen file and sends it using wormhole-william.
func (c *Client) NewFileSend(file fyne.URIReadCloser, progress wormhole.SendOption) (string, chan wormhole.SendResult, fyne.URIReadCloser, error) {
	code, result, err := c.SendFile(context.Background(), file.URI().Name(), file.(io.ReadSeeker), progress)
	return code, result, file, err
}

// NewDirSend takes a listable URI and sends it using wormhole-william.
func (c *Client) NewDirSend(dir fyne.ListableURI, progress wormhole.SendOption) (string, chan wormhole.SendResult, error) {
	dirpath := dir.String()[7:]
	prefix, _ := filepath.Split(dirpath)

	var files []wormhole.DirectoryEntry
	if err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}

		if runtime.GOOS == "windows" {
			path = filepath.ToSlash(path)
		}

		files = append(files, wormhole.DirectoryEntry{
			Path: strings.TrimPrefix(path, prefix),
			Mode: info.Mode(),
			Reader: func() (io.ReadCloser, error) {
				return os.Open(path) // #nosec - path is already cleaned by filepath.Walk
			},
		})

		return nil
	}); err != nil {
		fyne.LogError("Error on walking directory", err)
		return "", nil, err
	}

	return c.SendDirectory(context.Background(), dir.Name(), files, progress)
}

// NewTextSend takes a text input and sends the text using wormhole-william.
func (c *Client) NewTextSend(text string, progress wormhole.SendOption) (string, chan wormhole.SendResult, error) {
	return c.SendText(context.Background(), text, progress)
}
