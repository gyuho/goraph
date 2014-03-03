package spfw

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gt"
)

func Example_SPFW_4s() {
	testgraph4 := `
S|A,15|B,14|C,9
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
D|A,20|B,30|E,2|F,11|T,16
E|B,18|C,24|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`
	g4 := gt.ParseToGraph(testgraph4)
	g4s, g4m := SPFW(g4, "S", "T")
	fmt.Println(g4s)
	fmt.Println(g4m)

	// Output:
	// 44
	//   0  15  14   9  34  44  32  38
	//  15   0   5  24  20  34  22  28
	//  14   5   0  23  20  30  18  24
	//   9  24  23   0  26  36  24  30
	//  34  20  20  26   0  14   2   8
	//  44  34  30  36  14   0  12   6
	//  32  22  18  24   2  12   0   6
	//  38  28  24  30   8   6   6   0
}

func Example_SPFW_5s() {
	testgraph5 := `
A|B,7|C,9|F,20
B|A,7|C,10|D,15
C|A,9|B,10|D,11|E,30|F,2
D|B,15|C,11|E,2
E|C,30|D,2|F,9
F|A,20|C,2|E,9
`
	g5 := gt.ParseToGraph(testgraph5)
	g5s, g5m := SPFW(g5, "A", "E")
	fmt.Println(g5s)
	fmt.Println(g5m)

	// Output:
	// 20
	//   0   7   9  11  20  20
	//   7   0  10  12  15  17
	//   9  10   0   2  11  11
	//  11  12   2   0  11   9
	//  20  15  11  11   0   2
	//  20  17  11   9   2   0
}

func Example_SPFW_10s() {
	testgraph10 := `
A|C,9|F,20
B|A,1|D,15
C|B,10|E,30
D|C,11|E,2
E|C,30|F,9
F|A,20|C,2
`
	g10 := gt.ParseToGraph(testgraph10)
	g10s, g10m := SPFW(g10, "A", "E")
	fmt.Println(g10s)
	fmt.Println(g10m)

	// Output:
	// 36
	//   0   9  20  19  34  36
	//  11   0  31  10  25  27
	//  13   2   0  12  27  29
	//   1  10  21   0  15  17
	//  22  11  11  21   0   2
	//  22  11   9  21  36   0
}

func Example_SPFW_10o() {
	testgraph10 := `
A|C,9|F,20
B|A,1|D,15
C|B,10|E,30
D|C,11|E,2
E|C,30|F,9
F|A,20|C,2
`
	g10o := gt.ParseToGraph(testgraph10)
	g10so, g10m := SPFW(g10o, "E", "A")
	fmt.Println(g10so)
	fmt.Println(g10m)

	// Output:
	// 22
	//   0   9  20  19  34  36
	//  11   0  31  10  25  27
	//  13   2   0  12  27  29
	//   1  10  21   0  15  17
	//  22  11  11  21   0   2
	//  22  11   9  21  36   0

}

func Example_SPFW_11s() {
	testgraph11 := `
S|A,11|B,17|C,9
A|S,11|B,5|D,50|T,500
B|S,17|D,30
C|S,9
D|A,50|B,30|E,3|F,11
E|B,18|C,27|D,3|T,19
F|D,11|E,6|T,77
T|A,500|D,10|F,77|E,19
`
	g11 := gt.ParseToGraph(testgraph11)
	g11s, g11m := SPFW(g11, "S", "T")
	fmt.Println(g11s)
	fmt.Println(g11m)

	// Output:
	// 68
	//   0  11  16   9  46  68  49  57
	//  11   0   5  20  35  57  38  46
	//  17  28   0  26  30  52  33  41
	//   9  20  25   0  55  77  58  66
	//  38  49  21  30   0  22   3  11
	//  48  59  31  40  10   0  13  21
	//  35  46  18  27   3  19   0  14
	//  41  52  24  33   9  25   6   0
}

func Example_SPFW_11o() {
	testgraph11 := `
S|A,11|B,17|C,9
A|S,11|B,5|D,50|T,500
B|S,17|D,30
C|S,9
D|A,50|B,30|E,3|F,11
E|B,18|C,27|D,3|T,19
F|D,11|E,6|T,77
T|A,500|D,10|F,77|E,19
`
	g11o := gt.ParseToGraph(testgraph11)
	g11so, g11om := SPFW(g11o, "T", "S")
	fmt.Println(g11so)
	fmt.Println(g11om)

	// Output:
	// 48
	//   0  11  16   9  46  68  49  57
	//  11   0   5  20  35  57  38  46
	//  17  28   0  26  30  52  33  41
	//   9  20  25   0  55  77  58  66
	//  38  49  21  30   0  22   3  11
	//  48  59  31  40  10   0  13  21
	//  35  46  18  27   3  19   0  14
	//  41  52  24  33   9  25   6   0
}
