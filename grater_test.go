
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

	var max uint64 = 894514441

	for n = 0; n < max; n++ {
		_ = 100 * n / 8
	}
	grater.Stop()
	
}

