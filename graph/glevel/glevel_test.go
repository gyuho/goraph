package glevel

import (
	"os"
	"reflect"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestOpenGraph(t *testing.T) {
	db1, err := leveldb.OpenFile("test_leveldb", nil)
	if err != nil {
		t.Fatal(err)
	}
	c1 := reflect.TypeOf(db1)
	db1.Close()
	os.RemoveAll("test_leveldb")

	db2 := OpenGraph("test_glevel")
	c2 := reflect.TypeOf(db2)
	db2.Close()
	os.RemoveAll("test_glevel")
	if c1 != c2 {
		t.Fatalf("expected to be same but\n%v\n%v", c1, c2)
	}
}

func TestPutGetVertex(t *testing.T) {
	defer os.RemoveAll("test_putgetvertex")

	db := OpenGraph("test_putgetvertex")
	db.Put([]byte("A"), []byte("dd"), nil)
	PutVertex(db, "credit", 123.23)

	c1, _ := db.Get([]byte("A"), nil)
	c2 := GetVertex(db, "credit")

	if (string(c1) != "dd") || (c2 != 123.23) {
		t.Fatalf("expected\n%v\n%v", c1, c2)
	}
}

func TestDeleteVertex(t *testing.T) {
	defer os.RemoveAll("test_deletevertex")

	db := OpenGraph("test_deletevertex")
	PutVertex(db, "credit", 123.23)
	db.Close()

	db = OpenGraph("test_deletevertex")
	DeleteVertex(db, "credit")
	db.Close()

	data, err := db.Get([]byte("credit"), nil)
	if err == nil {
		t.Fatal(err, data)
	}
}

func TestPutGetEdge(t *testing.T) {
	defer os.RemoveAll("test_putgetedge")
	db := OpenGraph("test_putgetedge")

	PutEdge(db, "credit", 123.23, "card")
	c1 := GetEdge(db, "credit", "card")
	if c1 != 123.23 {
		t.Fatalf("expected\n%v\n%v", c1)
	}
}

func TestDeleteEdge(t *testing.T) {
	defer os.RemoveAll("test_deleteedge")
	db := OpenGraph("test_deleteedge")

	PutEdge(db, "credit", 123.23, "card")
	data, err := db.Get([]byte("credit----->card"), nil)
	if (StringToFloat64(string(data)) != 123.23) && (err != nil) {
		t.Fatal(data, err)
	}
	db.Close()

	db = OpenGraph("test_deleteedge")
	DeleteEdge(db, "credit", "card")
	data, err = db.Get([]byte("credit----->card"), nil)
	if err == nil {
		t.Fatal(err, data)
	}
	db.Close()
}

func TestGetOutInEdges(t *testing.T) {
	defer os.RemoveAll("test_getedges")
	db := OpenGraph("test_getedges")

	PutEdge(db, "credit", 123.23, "card1")
	PutEdge(db, "credit", 123.23, "card2")
	PutEdge(db, "credit", 123.23, "card3")
	PutEdge(db, "credit", 123.23, "card4")
	PutEdge(db, "credit", 123.23, "card5")

	PutEdge(db, "credit1", 123.23, "card")
	PutEdge(db, "credit2", 123.23, "card")
	PutEdge(db, "credit3", 123.23, "card")
	PutEdge(db, "credit4", 123.23, "card")
	PutEdge(db, "credit5", 123.23, "card")

	db.Close()

	db = OpenGraph("test_getedges")
	rs1 := GetOutEdges(db, "credit")
	if len(rs1) != 5 {
		t.Fatalf("GetOutEdges expected 5 but %#v\n", rs1)
	}

	rs2 := GetInEdges(db, "card")
	if len(rs2) != 5 {
		t.Fatalf("GetInEdges expected 5 but %#v\n", rs2)
	}

	db.Close()
}

func TestGetOutInVertices(t *testing.T) {
	defer os.RemoveAll("test_getvertices")
	db := OpenGraph("test_getvertices")

	PutEdge(db, "credit", 123.23, "card1")
	PutEdge(db, "credit", 123.23, "card2")
	PutEdge(db, "credit", 123.23, "card3")
	PutEdge(db, "credit", 123.23, "card4")
	PutEdge(db, "credit", 123.23, "card5")

	PutEdge(db, "credit1", 123.23, "card")
	PutEdge(db, "credit2", 123.23, "card")
	PutEdge(db, "credit3", 123.23, "card")
	PutEdge(db, "credit4", 123.23, "card")
	PutEdge(db, "credit5", 123.23, "card")

	db.Close()

	db = OpenGraph("test_getvertices")
	rs1 := GetOutVertices(db, "credit")
	if len(rs1) != 5 {
		t.Fatalf("GetOutVertices expected 5 but %#v\n", rs1)
	}

	rs2 := GetInVertices(db, "card")
	if len(rs2) != 5 {
		t.Fatalf("GetInVertices expected 5 but %#v\n", rs2)
	}

	db.Close()
}

func TestStringVertex(t *testing.T) {
	defer os.RemoveAll("test_stringvertex")
	db := OpenGraph("test_stringvertex")

	PutVertex(db, "credit", 123.23)

	PutEdge(db, "credit", 123.23, "card1")
	PutEdge(db, "credit", 123.23, "card2")
	PutEdge(db, "credit", 123.23, "card3")
	PutEdge(db, "credit", 123.23, "card4")
	PutEdge(db, "credit", 123.23, "card5")

	PutEdge(db, "credit1", 123.23, "card")
	PutEdge(db, "credit2", 123.23, "card")
	PutEdge(db, "credit3", 123.23, "card")
	PutEdge(db, "credit4", 123.23, "card")
	PutEdge(db, "credit5", 123.23, "card")

	db.Close()

	db = OpenGraph("test_stringvertex")

	rs1 := StringVertex(db, "credit")
	rs2 := `credit has a value of 123.23.

Outgoing Vertices:
credit -- 123.23 -> card3
credit -- 123.23 -> card4
credit -- 123.23 -> card5
credit -- 123.23 -> card1
credit -- 123.23 -> card2

Incoming Vertices:
No Incoming Vertices...
`
	if len(rs1) != len(rs2) {
		t.Fatalf("Expected %d but %d", len(rs1), len(rs2))
	}

	db.Close()
}

func TestSearchVertexBy(t *testing.T) {
	defer os.RemoveAll("test_searchvtxby")
	db := OpenGraph("test_searchvtxby")

	PutVertex(db, "credit", 123.23)

	PutVertex(db, "credit-1231232", 123.23)
	PutVertex(db, "credit-sadsf", 123.23)
	PutVertex(db, "creditasdfdas231232", 123.23)
	PutVertex(db, "creditasdf31232", 123.23)
	PutVertex(db, "creditasdfasdf1232", 123.23)

	PutVertex(db, "sdfasdf1232credit", 123.23)
	PutVertex(db, "s123213232credit", 123.23)
	PutVertex(db, "asdfdasfasdf1232credit", 123.23)
	PutVertex(db, "zxvzxcvdf1232credit", 123.23)
	PutVertex(db, "zxcvczxvxvsdfsdf1232credit", 123.23)

	PutVertex(db, "sdfasdf1credit453454235432", 123.23)
	PutVertex(db, "s1232132credit4235432", 123.23)
	PutVertex(db, "asdfdasfacredit23453454235432", 123.23)
	PutVertex(db, "zxvzxcvdf12credit235432", 123.23)
	PutVertex(db, "zxcvczxcreditfsdf12323453454235432", 123.23)

	PutEdge(db, "credit", 123.23, "card1")
	PutEdge(db, "credit", 123.23, "card2")
	PutEdge(db, "credit", 123.23, "card3")
	PutEdge(db, "credit", 123.23, "card4")
	PutEdge(db, "credit", 123.23, "card5")

	PutEdge(db, "credit1", 123.23, "card")
	PutEdge(db, "credit2", 123.23, "card")
	PutEdge(db, "credit3", 123.23, "card")
	PutEdge(db, "credit4", 123.23, "card")
	PutEdge(db, "credit5", 123.23, "card")

	db.Close()

	db = OpenGraph("test_searchvtxby")

	rs1 := SearchVertexByPrefix(db, "credit")
	if len(rs1) != 6 {
		t.Fatalf("Expected 6 but\n%v", rs1)
	}

	rs2 := SearchVertexBySuffix(db, "credit")
	if len(rs2) != 6 {
		t.Fatalf("Expected 6 but\n%v", rs2)
	}

	rs3 := SearchVertexBySubstring(db, "credit")
	if len(rs3) != 16 {
		t.Fatalf("Expected 16 but\n%v", rs3)
	}

	db.Close()
}

func TestStringToFloat64(t *testing.T) {
	num1 := 123.123
	num2 := StringToFloat64("123.123")
	if num1 != num2 {
		t.Fatalf("expected\n%v\n%v", num1, num2)
	}
}
