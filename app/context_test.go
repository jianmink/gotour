package app

import (
	"context"
	//"net/http"
	//"reflect"
	"testing"
	"fmt"
	"time"
)

func Test_newContextWithRequestID(t *testing.T) {
	c := context.Background()

	var r Req
	r.kv = map [string]string{}
	r.Set("X-Request-ID","my id")

	c1 := newContextWithRequestID(c,&r)
	fmt.Println(requestIDFromContext(c1))
	fmt.Println(c)
	fmt.Println(c1)

	c2 := context.WithValue(c1, "2", "second id")
	c2 = context.WithValue(c2, "2", "second id")
	fmt.Println(c)
	fmt.Println(c1)
	fmt.Println(c2)
}


// otherContext is a Context that's not one of the types defined in context.go.
// This lets us test code paths that differ based on the underlying type of the
// Context.
type otherContext struct {
	context.Context
}


func Test_newContextWithValue(t *testing.T) {
	c1 := context.WithValue(context.Background(), "1", "a")

	c1.Err()

	v1 := c1.Value("1").(string)  // .(int) will raise panic

	fmt.Printf("%s\n",v1)
}

func Test_newContextWithCancel(t *testing.T) {
	c1, cancel := context.WithCancel(context.Background())

	if got, want := fmt.Sprint(c1), "context.Background.WithCancel"; got != want {
		t.Errorf("c1.String() = %q want %q", got, want)
	}


	//o := otherContext{c1}
	//o,_:= context.WithCancel(c1)
	//c2, _ := context.WithCancel(o)
	c2,_ := context.WithCancel(c1)
	contexts := []context.Context{c1, c2}

	for i, c := range contexts {
		fmt.Println(i,c.Done())
		if d := c.Done(); d == nil {
			t.Errorf("c[%d].Done() == %v want non-nil", i, d)
		}
		if e := c.Err(); e != nil {
			t.Errorf("c[%d].Err() == %v want nil", i, e)
		}

		select {
		case x := <-c.Done():
			t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
		default:
		}
	}

	cancel()
	time.Sleep(100 * time.Millisecond) // let cancelation propagate

	for i, c := range contexts {
		select {
		case x:=<-c.Done():
			fmt.Println(x)
		default:
			t.Errorf("<-c[%d].Done() blocked, but shouldn't have", i)
		}
		if e := c.Err(); e != context.Canceled {
			t.Errorf("c[%d].Err() == %v want %v", i, e, context.Canceled)
		}
	}
}

