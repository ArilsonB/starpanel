package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeInstallNginx = "nginx:install"
)

type InstallNginxPayload struct {
	PackageName    string `json:"package_name"`
	PackagePath    string `json:"package_path"`
	PackageVersion string `json:"package_version"`
	PackageURL     string `json:"package_url"`
}

func InstallNginxTask(packageName string, packagePath string, packageVersion string, packageUrl string) (*asynq.Task, error) {
	payload, err := json.Marshal(InstallNginxPayload{
		PackageName:    packageName,
		PackagePath:    packagePath,
		PackageVersion: packageVersion,
		PackageURL:     packageUrl,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeInstallNginx, payload), nil
}
