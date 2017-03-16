package expvar

import (
	"fmt"
	"net/http"
	"sort"
	"sync"

	"nkwangwenfang.com/kit/monitor"
	"nkwangwenfang.com/log"
)

type Variable interface {
	String() string
}

var (
	mutex   sync.RWMutex
	vars    = make(map[string]Variable)
	varKeys []string
)

func init() {
	monitor.HandleFunc("/debug/vars", expvarHandler)
}

func expvarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{\n")
	first := true
	mutex.RLock()
	for _, key := range varKeys {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", key, vars[key])
	}
	mutex.RUnlock()
	fmt.Fprintf(w, "\n}\n")
}

func publish(name string, v Variable) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, existing := vars[name]; existing {
		log.Error("variable export twice", "name", name)
		return
	}

	vars[name] = v
	varKeys = append(varKeys, name)
	sort.Strings(varKeys)
}
