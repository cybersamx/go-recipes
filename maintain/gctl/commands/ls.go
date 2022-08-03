package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var lsCmd = cobra.Command{
	Use:   "ls",                         // Name of the command
	Short: "Run a task on the cluster.", // Short descriptor in the usage table of cobra-cmd --help
	Long:  "Run a task on the cluster.", // Long help description in the output of cobra-cmd run --help
	RunE:  run,
}

type Project struct {
	Name           string
	Path           string
	HasStdMakefile string
	Next           *Project
}

func init() {
	rootCmd.AddCommand(&lsCmd)
}

func isHidden(file string) bool {
	if len(file) == 0 {
		return false // Treat it as .
	}

	// Except for .. and ., any name that starts with . is hidden.
	return file[0:1] == "." && (file != ".." && file != ".")
}

func matchIgnoreFiles(file string) bool {
	ignores := []string{"node_modules"}
	for _, ignore := range ignores {
		if file == ignore {
			return true
		}
	}

	return false
}

func genDirTree(rootdir, filepath string, ll *LinkedList[Project]) {
	_, file := path.Split(filepath)

	// Ignore hidden directory/file.
	if isHidden(file) {
		return
	}
	// Ignore if name matches.
	if matchIgnoreFiles(file) {
		return
	}

	finfos, err := ioutil.ReadDir(filepath)
	if err != nil {
		panic(err)
	}

	// Does current directory has a Makefile.
	hasMakefile := false
	for _, finfo := range finfos {
		if finfo.Name() == "Makefile" {
			hasMakefile = true
			break
		}
	}

	// Go deeper.
	for _, finfo := range finfos {
		if finfo.IsDir() {
			genDirTree(rootdir, path.Join(filepath, finfo.Name()), ll)
		}
	}

	// Ignore if directory not contain Makefile.
	if !hasMakefile {
		return
	}

	proj := Project{
		Name: filepath[len(rootdir):],
		Path: filepath,
	}

	ll.Add(proj)
}

func run(cmd *cobra.Command, args []string) error {
	rootdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if len(args) > 0 {
		rootdir = args[0]
	}

	var projs = &LinkedList[Project]{}
	genDirTree(rootdir, rootdir, projs)

	for elem := projs; elem != nil; elem = elem.Next {
		fmt.Println(elem.Value.Name)
	}

	return nil
}
