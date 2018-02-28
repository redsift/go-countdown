package countdown

import (
	"context"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	//ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	ctdwn := NewCountdown(ctx, 2*time.Second)

	go func() {
		time.Sleep(2 * time.Second)
		ctdwn.Sub(2 * time.Second)
	}()

	select {
	case <-ctdwn.Done():
	case <-time.Tick(3 * time.Second):
		t.Fatalf("Should have timeout out before 3 seconds")
	}
}
