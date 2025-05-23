package ansible

import (
	"fmt"
	"os"
	goexec "os/exec"
	"path/filepath"

	"k8s.io/klog/v2"

	"sigs.k8s.io/provider-ibmcloud-test-infra/kubetest2-tf/data"
)

const (
	ansibleDataDir = "k8s-ansible"
)

func Playbook(dir, inventory, extraVars, playbook string) (int, error) {
	err := unpackAnsible(dir)
	if err != nil {
		return 1, fmt.Errorf("failed to unpack the ansible code: %v", err)
	}
	args := []string{
		fmt.Sprintf("--inventory=%s", inventory),
		fmt.Sprintf("--extra-vars=%s", extraVars),
		fmt.Sprintf("%s", filepath.Join(dir, playbook)),
	}
	klog.Infof("ansible-playbook with args: %v", args)
	c := goexec.Command("ansible-playbook", args...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		// Try to get the exit code
		if exitError, ok := err.(*goexec.ExitError); ok {
			return exitError.ExitCode(), err
		} else {
			// This will happen if ansible is not available in $PATH
			return 1, err
		}
	} else {
		// successful execution of ansible playbook
		return 0, nil
	}
}

func unpackAnsible(dir string) error {
	return data.Unpack(dir, ansibleDataDir)
}
