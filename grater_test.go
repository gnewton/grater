
package grater

import(
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	var n uint64 = 0

	grater, err := NewGrater(time.Second, &n)
	if err != nil{
	    t.Error(err)
	}
	runGrater(grater, &n)
}


 func TestNegativeDuration(t *testing.T) {
 	var n uint64 = 0

 	_, err := NewGrater(0*time.Second, &n)
 	if err == nil{
 	    t.Error(err)
 	}


 }

 func TestNilCounter(t *testing.T) {
 	_, err := NewGrater(time.Second, nil)
 	if err == nil{
 	    t.Error(err)
 	}


 }

func TestGraterInitNilOutFunc(t *testing.T) {
	var n uint64;
	_, err := graterInit(time.Second, &n, nil, nil)
	if err == nil{
		t.Error(err)
	}
 }

func TestGraterWriterOutNilGrater(t *testing.T) {
	err := writerOut(nil, "")
	if err == nil{
		t.Error()
	}
}


func TestGraterWriterOutNilWriter(t *testing.T) {
	var n uint64 = 0

	grater, err := NewGrater(time.Second, &n)
	if err != nil{
	    t.Error(err)
	}
	err = writerOut(grater, "")
	if err == nil{
		t.Error()
	}
}

func TestGraterLogOutNilGrater(t *testing.T) {
	err := logOut(nil, "")
	if err == nil{
		t.Error()
	}
}


func runGrater(g *Grater, n *uint64){
	var max uint64 = 894514441

	for *n = 0; *n < max; (*n)++ {
		_ = 100 * (*n) / 8
	}
	g.Stop()
}
