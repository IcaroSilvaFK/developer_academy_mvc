package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadTemplatesFiles(path string) ([]string, error) {

	paths := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if strings.Contains(path, ".tmpl") {
			paths = append(paths, path)

			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}
