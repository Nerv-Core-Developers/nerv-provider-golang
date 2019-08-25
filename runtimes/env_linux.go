package runtimes

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/naokichau/nerv-provider-golang/shared"
)

// Create isolate env to run service
func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "nervrt"), 0755)
	shared.Must(ioutil.WriteFile(filepath.Join(pids, "nervrt/pids.max"), []byte("20"), 0700))
	// Removes the new cgroup in place after the container exits
	shared.Must(ioutil.WriteFile(filepath.Join(pids, "nervrt/notify_on_release"), []byte("1"), 0700))
	shared.Must(ioutil.WriteFile(filepath.Join(pids, "nervrt/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func child() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	shared.Must(syscall.Sethostname([]byte("container")))
	shared.Must(syscall.Chroot("/nervrt"))
	shared.Must(os.Chdir("/"))
	shared.Must(syscall.Mount("proc", "proc", "proc", 0, ""))
	shared.Must(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))

	shared.Must(cmd.Run())

	shared.Must(syscall.Unmount("proc", 0))
	shared.Must(syscall.Unmount("thing", 0))
}
func run() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	shared.Must(cmd.Run())
}
