package lookup

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Flag interface {
	String() string
}

type Lookup struct {
	store map[string]string
}

func (l *Lookup) Set(key, value Flag) {
	l.store[key.String()] = value.String()
}

func (l *Lookup) Get(key Flag) string {
	return l.store[key.String()]
}

func New(cmd *cobra.Command) *Lookup {
	var lookupMap = make(map[string]string, 0)

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		fmt.Printf("Flags().VisitAll: %s=[%s]\n", f.Name, f.Value.String())
		lookupMap[f.Name] = f.Value.String()
	})

	return &Lookup{store: lookupMap}
}

func NewPersistent(cmd *cobra.Command) *Lookup {
	var lookupMap = make(map[string]string, 0)

	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		fmt.Printf("PersistentFlags().VisitAll: %s=[%s]\n", f.Name, f.Value.String())
		lookupMap[f.Name] = f.Value.String()
	})

	return &Lookup{store: lookupMap}
}
