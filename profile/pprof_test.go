package profile

import (
	"log"
	"os"
	"runtime/pprof"
	"testing"
)

func TestPrfile(t *testing.T) {
	f1, err := os.Create("./cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	err = pprof.StartCPUProfile(f1)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	f2, err := os.Create("./mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	// runtime.GC()
	err = pprof.WriteHeapProfile(f2)
	if err != nil {
		log.Fatal(err)
	}
	f3, err := os.Create("./goroutine.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f3.Close()
	gProf := pprof.Lookup("goroutine")
	if gProf == nil {
		log.Fatal(gProf)
	} else {
		gProf.WriteTo(f3, 0)
	}
	AULF()
}
