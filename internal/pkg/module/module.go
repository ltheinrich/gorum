package module

import (
	"io/ioutil"
	"os"
	"plugin"
	"strings"
)

var (
	// module list ID:Module
	modules = map[string]Module{}
)

// Module structure
type Module struct {
	Name    string
	Author  string
	Version string
}

// Exists check if module exists
func Exists(name string) bool {
	// loop through modules
	for _, module := range modules {
		// check if module name is equal to provided
		if module.Name == name {
			// exists
			return true
		}
	}

	// not found
	return false
}

// LoadModules load all plugins in modules
func LoadModules(dir string) error {
	var err error

	// create directory if not exists
	_, exists := os.Stat(dir)
	if os.IsNotExist(exists) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// list files
	var files []os.FileInfo
	files, err = ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// loop through files and load modules
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			err = LoadModule(dir + "/" + file.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// LoadModule load plugin from file
func LoadModule(path string) error {
	var err error

	// load plugin
	var p *plugin.Plugin
	p, err = plugin.Open(path)
	if err != nil {
		return err
	}

	// load start function
	var start plugin.Symbol
	start, err = p.Lookup("Start")
	if err != nil {
		return err
	}

	// start module and add to map
	modules[path] = start.(func() Module)()

	return nil
}

// RemoveModule disable plugin and delete
func RemoveModule(path string) error {
	var err error

	// load plugin
	var p *plugin.Plugin
	p, err = plugin.Open(path)
	if err != nil {
		return err
	}

	// load stop function
	var stop plugin.Symbol
	stop, err = p.Lookup("Stop")
	if err != nil {
		return err
	}

	// stop and remove module
	stop.(func())()
	delete(modules, path)

	// delete
	os.Remove(path)

	return nil
}

// GetPath of module by name
func GetPath(name string) string {
	// loop through modules
	for path, module := range modules {
		// check if module name is equal to provided
		if module.Name == name {
			// return path
			return path
		}
	}

	// not found
	return ""
}
