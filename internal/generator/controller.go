package generator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateControllerStructure(input string, subFolder string) error {
	// 1. Validate controller name
	if !strings.HasSuffix(input, "Controller") {
		return errors.New("name must end with Controller")
	}

	// 2. Extract root name
	root := strings.TrimSuffix(input, "Controller")
	root = strings.ToLower(root)

	// 3. Current directory = project root
	projectRoot, err := os.Getwd()
	if err != nil {
		return err
	}

	// 4. Ensure cmd/routes exists
	routesPath := filepath.Join(projectRoot, "cmd", "routes")
	if err := os.MkdirAll(routesPath, 0755); err != nil {
		return err
	}

	// 5. If subFolder provided → cmd/routes/subFolder
	baseRoutesPath := routesPath
	if subFolder != "" {
		baseRoutesPath = filepath.Join(routesPath, subFolder)
		if err := os.MkdirAll(baseRoutesPath, 0755); err != nil {
			return err
		}
	}

	// 6. Final base path
	basePath := filepath.Join(baseRoutesPath, root)

	// 7. Create directories
	dirs := []string{
		"database/entity",
		"database/service",
		"domain/http",
		"domain/model",
		"model",
	}

	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(basePath, d), 0755); err != nil {
			return err
		}
	}

	// 8. Create files (no overwrite)
	files := map[string]string{
		fmt.Sprintf("%sController.go", root):               fmt.Sprintf("package %sController", root),
		fmt.Sprintf("database/entity/%sEntity.go", root):   "package entity",
		fmt.Sprintf("database/service/%sService.go", root): "package service",
		"domain/http/httpInterface.go":                     "package httpInterface",
		fmt.Sprintf("domain/model/%sHttp.go", root):        "package httpModel",
		fmt.Sprintf("model/%sModel.go", root):              fmt.Sprintf("package %sModels", root),
	}

	for path, content := range files {
		fullPath := filepath.Join(basePath, path)

		if _, err := os.Stat(fullPath); err == nil {
			continue
		}

		if err := os.WriteFile(fullPath, []byte(content+"\n"), 0644); err != nil {
			return err
		}
	}

	fmt.Println("✅ Controller created at:", basePath)
	return nil
}
