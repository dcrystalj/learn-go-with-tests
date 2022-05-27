package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s1 := &S1{}
	s2 := &S1{}
	go s1.Run(ctx)
	go s2.Run(ctx)

	<-ctx.Done()
}

// func main() {
// 	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
// 	defer cancel()
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	s1 := &S1{}
// 	s2 := &S1{}
// 	go func() {
// 		defer wg.Done()
// 		s1.Run(ctx)
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		s2.Run(ctx)
// 	}()

// 	wg.Wait()
// }

type S1 struct{}

func (s *S1) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := ctx.Done()
	for {
		select {
		case <-done:
			fmt.Println("Started long service shutdown")
			time.Sleep(time.Second)
			fmt.Println("Finished long service shutdown")
			return
		case <-ticker.C:
			fmt.Println("ticking s1")
			time.Sleep(time.Second * 10)
			fmt.Println("Finished long service ticking")
		}
	}
}
