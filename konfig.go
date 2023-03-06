package konfig

import (
	"io"
	"os"
	"strings"
)

func parse(config string) map[string]string {
	var cache = make(map[string]string)
	lines := strings.Split(config, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		index := strings.Index(line, "=")
		if index <= 0 {
			continue
		}
		cache[line[:index]] = line[index+1:]
	}
	return cache
}

func ReadConfig(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	st, _ := file.Stat()

	buffer := make([]byte, st.Size())
	_, err = io.ReadFull(file, buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return parse(string(buffer)), nil
}
