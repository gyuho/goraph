package kosaraju

func main() {}

/*
import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestSCC(t *testing.T) {
	g15 := gs.FromJSON("../../../files/testgraph.json", "testgraph.015")
	gr15 := gs.FromJSONT("../../../files/testgraph.json", "testgraph.015")
	rs := SCC(g15, gr15)
	fmt.Println(rs)
	// [[B E A] [D C] [G F] [H]]

	if len(rs) != 4 {
		t.Errorf("expected 4 but %v", rs)
	}
	//
	//
	// TODO
	// g16 := gs.FromJSON("../../../files/testgraph.json", "testgraph.016")
	// gr16 := gs.FromJSONT("../../../files/testgraph.json", "testgraph.016")
	// fmt.Println(SCC(g16, gr16))
	// [[B F G A] [D H C] [I] [E J]]
}
*/

/*
=== RUN TestSCC
SIGQUIT: quit
PC=0x42233e

goroutine 25 [running]:
runtime.mallocgc(0x20, 0x513f40, 0xc200000000)
	/usr/local/go/src/pkg/runtime/malloc.goc:156 +0x32e fp=0xc20808dc98 sp=0xc20808dc30
runtime.new(0x513f40, 0x0)
	/usr/local/go/src/pkg/runtime/malloc.goc:826 +0x3b fp=0xc20808dcb8 sp=0xc20808dc98
github.com/gyuho/goraph/algorithm/scc/kosaraju.DFSandSCC(0xc208001970, 0xc20801a370, 0x0, 0x0, 0x0)
	/home/travis/gopath/src/github.com/gyuho/goraph/algorithm/scc/kosaraju/dfs.go:97 +0x41 fp=0xc20808dd00 sp=0xc20808dcb8
github.com/gyuho/goraph/algorithm/scc/kosaraju.SCC(0xc208001970, 0xc208089d70, 0x0, 0x0, 0x0)
	/home/travis/gopath/src/github.com/gyuho/goraph/algorithm/scc/kosaraju/kosaraju.go:20 +0x24f fp=0xc20808dec8 sp=0xc20808dd00
github.com/gyuho/goraph/algorithm/scc/kosaraju.TestSCC(0xc208048240)
	/home/travis/gopath/src/github.com/gyuho/goraph/algorithm/scc/kosaraju/kosaraju_test.go:13 +0xa4 fp=0xc20808df68 sp=0xc20808dec8
testing.tRunner(0xc208048240, 0x65d7b8)
	/usr/local/go/src/pkg/testing/testing.go:422 +0x8b fp=0xc20808df98 sp=0xc20808df68
runtime.goexit()
	/usr/local/go/src/pkg/runtime/proc.c:1445 fp=0xc20808dfa0 sp=0xc20808df98
created by testing.RunTests
	/usr/local/go/src/pkg/testing/testing.go:504 +0x8db

goroutine 16 [chan receive, 9 minutes]:
testing.RunTests(0x5cc860, 0x65d740, 0x6, 0x6, 0x1)
	/usr/local/go/src/pkg/testing/testing.go:505 +0x923
testing.Main(0x5cc860, 0x65d740, 0x6, 0x6, 0x665fc0, 0x0, 0x0, 0x665fc0, 0x0, 0x0)
	/usr/local/go/src/pkg/testing/testing.go:435 +0x84
main.main()
	github.com/gyuho/goraph/algorithm/scc/kosaraju/_test/_testmain.go:57 +0x9c

goroutine 19 [finalizer wait, 10 minutes]:
runtime.park(0x4130d0, 0x662a58, 0x661589)
	/usr/local/go/src/pkg/runtime/proc.c:1369 +0x89
runtime.parkunlock(0x662a58, 0x661589)
	/usr/local/go/src/pkg/runtime/proc.c:1385 +0x3b
runfinq()
	/usr/local/go/src/pkg/runtime/mgc0.c:2644 +0xcf
runtime.goexit()
	/usr/local/go/src/pkg/runtime/proc.c:1445

rax     0xc2080724c0
rbx     0x3
rcx     0xc2080724e0
rdx     0xc2080724e0
rdi     0x7f6a990bd000
rsi     0xc2080724c0
rbp     0x7f6a990c3100
rsp     0xc20808dc30
r8      0xc20808db98
r9      0x25
r10     0x0
r11     0x246
r12     0x411890
r13     0xc2080400f0
r14     0x0
r15     0xc2080003f0
rip     0x42233e
rflags  0x206
cs      0x33
fs      0x0
gs      0x0
*/
