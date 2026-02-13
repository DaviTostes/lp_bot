package vercel

import "os/exec"

func Deploy(dir string) (string, error) {
	cmd := exec.Command("npm", "install")
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return "", err
	}

	cmd = exec.Command("vercel", "link", "--yes")
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return "", err
	}

	cmd = exec.Command("vercel", "--prod")
	cmd.Dir = dir
	url, err := cmd.Output()
	return string(url), err
}
