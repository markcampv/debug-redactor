package redactor

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

// GenerateRedactedPath generates a new path for the redacted tar.gz file based on the source file path.
func GenerateRedactedPath(sourceTarGz string) string {
	dir, file := filepath.Split(sourceTarGz)
	ext := filepath.Ext(file)
	baseName := file[:len(file)-len(ext)]
	redactedFileName := fmt.Sprintf("%s-redacted%s", baseName, ext)
	return filepath.Join(dir, redactedFileName)
}

// ProcessAndRepackageTarGz redacts sensitive information from the source tar.gz and repackages it.
// It returns the path to the redacted tar.gz file and any error encountered.
func ProcessAndRepackageTarGz(sourceTarGz string) (string, error) {
	redactedTarGzPath := GenerateRedactedPath(sourceTarGz)

	gzipFile, err := os.Open(sourceTarGz)
	if err != nil {
		return "", fmt.Errorf("failed to open gzip archive: %w", err) // Ensure an empty string is returned for the path
	}
	defer gzipFile.Close()

	uncompressedStream, err := gzip.NewReader(gzipFile)
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader: %w", err) // Ensure an empty string is returned for the path
	}
	defer uncompressedStream.Close()

	tarReader := tar.NewReader(uncompressedStream)
	redactedFile, err := os.Create(redactedTarGzPath)
	if err != nil {
		return "", fmt.Errorf("failed to create redacted gzip archive: %w", err) // Ensure an empty string is returned for the path
	}
	defer redactedFile.Close()

	gzipWriter := gzip.NewWriter(redactedFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	ipRegex := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read tar entry: %w", err) // Ensure an empty string is returned for the path
		}

		var buf bytes.Buffer
		if header.Typeflag == tar.TypeReg {
			_, err = io.Copy(&buf, tarReader)
			if err != nil {
				return "", fmt.Errorf("failed to read file content: %w", err) // Ensure an empty string is returned for the path
			}

			redactedContent := ipRegex.ReplaceAllString(buf.String(), "[REDACTED]")
			header.Size = int64(len(redactedContent)) // Update the header size to match the redacted content length

			if err := tarWriter.WriteHeader(header); err != nil {
				return "", fmt.Errorf("failed to write header to redacted archive: %w", err) // Ensure an empty string is returned for the path
			}

			if _, err := tarWriter.Write([]byte(redactedContent)); err != nil {
				return "", fmt.Errorf("failed to write redacted content: %w", err) // Ensure an empty string is returned for the path
			}
		} else {
			// For non-file entries, just copy the header
			if err := tarWriter.WriteHeader(header); err != nil {
				return "", fmt.Errorf("failed to write header for non-file entry: %w", err) // Ensure an empty string is returned for the path
			}
			if _, err := io.Copy(tarWriter, tarReader); err != nil {
				return "", fmt.Errorf("failed to copy non-file entry: %w", err) // Ensure an empty string is returned for the path
			}
		}
	}

	// Return the path to the successfully processed (redacted) file
	return redactedTarGzPath, nil
}
