package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func contains(list []*flag.Flag, f *flag.Flag) bool {
	for _, i := range list {
		if i == f {
			return true
		}
	}
	return false
}

func readConfig() map[string]string {
	bytes, err := ioutil.ReadFile(*FLAG_CONFIG_FILE)
	if err != nil {
		log.Fatalf("Failed to read config file %s: %s", *FLAG_CONFIG_FILE, err)
	}
	lines := strings.Split(string(bytes), "\n")
	result := make(map[string]string, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || trimmed[0] == '#' {
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			log.Fatalf("Invalid config line: %s, len: %d", line, len(parts))
		}
		result[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	return result
}

func Parse() {
	flag.Parse()

	if *FLAG_CONFIG_FILE == "" {
		return
	}
	config := readConfig()
	explicit := make([]*flag.Flag, 0)
	all := make([]*flag.Flag, 0)
	flag.Visit(func(f *flag.Flag) { //启动项设置
		explicit = append(explicit, f)
	})
	flag.VisitAll(func(f *flag.Flag) {
		all = append(all, f)
		if !contains(explicit, f) {
			val := config[f.Name]
			if val != "" {
				err := f.Value.Set(val)
				if err != nil {
					log.Fatalf("Failed to set flag %s with value %s", f.Name, val)
				}
			}
		}
	})
Outer:
	for name, val := range config {
		for _, f := range all {
			if f.Name == name {
				continue Outer
			}
		}
		log.Fatalf("Unknown flag %s=%s in config file.", name, val) //with os.Exit(1)
	}
}

func PrintFlag() {
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(os.Stderr, "  -%s=%s: %s\n", f.Name, f.Value.String(), f.Usage)
	})
}

/*
func init() {
	flag.Parse()
	Parse()
}
*/
