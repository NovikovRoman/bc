package bc

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"archive/zip"
)

const bcURL = "http://api.bestchange.ru/info.zip"

func Download(ctx context.Context, dstDir string, client *http.Client) (err error) {
	if _, err = os.Stat(dstDir); os.IsNotExist(err) {
		if err = os.MkdirAll(dstDir, os.ModePerm); err != nil {
			return
		}
	}

	if client == nil {
		client = &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				ResponseHeaderTimeout: 60 * time.Second,
			},
		}
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, bcURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var b []byte
	if b, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	if err = os.WriteFile(filepath.Join(dstDir, "info.zip"), b, 0644); err != nil {
		return
	}

	err = unzipFromBytes(b, dstDir)
	return
}

func unzipFromBytes(data []byte, dest string) error {
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return fmt.Errorf("failed to create zip reader: %w", err)
	}

	for _, file := range reader.File {
		filePath := filepath.Join(dest, file.Name)

		// Проверяем на попытку выхода за пределы dest
		if !strings.HasPrefix(filepath.Clean(filePath), filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", filePath)
		}

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory for file: %w", err)
		}

		// Открывает файл внутри zip-архива
		srcFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file in zip: %w", err)
		}
		defer srcFile.Close()

		destFile, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, srcFile); err != nil {
			return fmt.Errorf("failed to copy file contents: %w", err)
		}
	}

	return nil
}
