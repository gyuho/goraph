package example_with_testing

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/algorithm/spfw"
	"github.com/gyuho/goraph/graph/gt"
	// go test -v github.com/gyuho/goraph/example_with_testing
	// go test -v /Users/gyuho/go/src/github.com/gyuho/goraph/example_with_testing/spfw_json_test.go
)

func Test_JSON_SPFW(test *testing.T) {
	fmt.Println("SPFW on testgraph4:")
	g4 := gt.JSONGraph("../example_files/testgraph.json", "testgraph.004")
	g4s, g4m := spfw.SPFW(g4, "S", "T")
	fmt.Println(g4s) // 44
	fmt.Println(g4m)
	/*
	     0  19  14  56  32  38  44  34
	    15   0   5  47  23  29  35  20
	    14   5   0  42  18  24  30  20
	     9  28  23   0  24  30  36  26
	    32  23  18  24   0   6  12   2
	    38  29  24  30   6   0   6   8
	    44  35  30  36  12   6   0  14
	   9999999999.9999 9999999999.9999 9999999999.9999 9999999999.9999 9999999999.9999 9999999999.9999 9999999999.9999   0
	*/

	println()
	fmt.Println("SPFW on testgraph5:")
	g5 := gt.JSONGraph("../example_files/testgraph.json", "testgraph.005")
	g5s, g5m := spfw.SPFW(g5, "A", "E")
	fmt.Println(g5s) // 20
	fmt.Println(g5m)
	/*
	    0   7   9  20  20  11
	    7   0  10  15  17  12
	    9  10   0  11  11   2
	   20  15  11   0   2  11
	   20  17  11   2   0   9
	   11  12   2  11   9   0
	*/

	println()
	fmt.Println("SPFW on testgraph10:")
	g10 := gt.JSONGraph("../example_files/testgraph.json", "testgraph.010")
	g10s, g10m := spfw.SPFW(g10, "A", "E")
	fmt.Println(g10s) // 36
	fmt.Println(g10m)
	/*
	    0  19   9  34  36  20
	    1   0  10  15  17  21
	   11  10   0  25  27  31
	   22  21  11   0   2  11
	   22  21  11  36   0   9
	   13  12   2  27  29   0
	*/

	println()
	fmt.Println("SPFW on testgraph10:")
	g10o := gt.JSONGraph("../example_files/testgraph.json", "testgraph.010")
	g10so, g10m := spfw.SPFW(g10o, "E", "A")
	fmt.Println(g10so) // 22
	fmt.Println(g10m)
	/*
	    0  19   9  34  36  20
	    1   0  10  15  17  21
	   11  10   0  25  27  31
	   22  21  11   0   2  11
	   22  21  11  36   0   9
	   13  12   2  27  29   0
	*/

	println()
	fmt.Println("SPFW on testgraph10:")
	g11 := gt.JSONGraph("../example_files/testgraph.json", "testgraph.011")
	g11s, g11m := spfw.SPFW(g11, "S", "T")
	fmt.Println(g11s) // 68
	fmt.Println(g11m)
	/*
	    0  11  16   9  46  49  55  68
	   11   0   5  20  35  38  44  57
	   17  28   0  26  30  33  39  52
	    9  20  25   0  55  58  64  77
	   38  49  21  47   0   3   9  22
	   35  46  18  44   2   0   6  19
	   41  52  24  50   8   6   0  25
	   48  59  31  57  10  13  19   0
	*/

	println()
	fmt.Println("SPFW on testgraph11:")
	g11o := gt.JSONGraph("../example_files/testgraph.json", "testgraph.011")
	g11so, g11om := spfw.SPFW(g11o, "T", "S")
	fmt.Println(g11so) // 48
	fmt.Println(g11om)
	/*
	    0  11  16   9  46  49  55  68
	   11   0   5  20  35  38  44  57
	   17  28   0  26  30  33  39  52
	    9  20  25   0  55  58  64  77
	   38  49  21  47   0   3   9  22
	   35  46  18  44   2   0   6  19
	   41  52  24  50   8   6   0  25
	   48  59  31  57  10  13  19   0
	*/
}
