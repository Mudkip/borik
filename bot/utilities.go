package bot

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// ImageURIFromMessage attempts to retrieve an image URI for a given message.
func ImageURIFromMessage(m *discordgo.Message) (string, bool) {
	if len(m.Embeds) == 1 {
		embed := m.Embeds[0]

		if embed.Type == "Image" {
			return embed.URL, true
		}
	}

	if len(m.Attachments) == 1 {
		attachment := m.Attachments[0]
		return attachment.URL, true
	}

	return "", false
}

// ImageURIFromCommand attempts to retrieve an image URI from a given command.
func ImageURIFromCommand(s *discordgo.Session, m *discordgo.MessageCreate, commandPrefix string) (string, error) {
	argument := strings.TrimSpace(strings.TrimPrefix(m.Content, commandPrefix))

	if argument != "" {
		return argument, nil
	}

	imageURI, found := ImageURIFromMessage(m.Message)
	if found {
		return imageURI, nil
	}

	messages, err := s.ChannelMessages(m.ChannelID, 20, m.ID, "", "")
	if err != nil {
		return "", err
	}

	for _, message := range messages {
		imageURI, found := ImageURIFromMessage(message)
		if found {
			return imageURI, nil
		}
	}

	return "", errors.New("no image found from message")
}

// DownloadFile downloads a file from a URL to a given path on disk.
func DownloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}