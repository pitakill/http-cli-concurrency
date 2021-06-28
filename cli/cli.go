package cli

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func Start(ch chan<- string) {
	ch <- "Starting the CLI"
	ch <- "You can start typing the method and the path to make the petition"
	ch <- "Example: GET /people"

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")

		cmd, err := exec.Command("curl", "-L", "-X", text[0], "http://localhost:8080"+text[1]).Output()
		if err != nil {
			ch <- err.Error()
			continue
		}
		ch <- string(cmd)
	}
}
