package main

import (
	"log"

	"math/rand"
	"sync"

	"os"
	"strconv"

	"fmt"
	"time"

	"github.com/morganwu277/kvdb/server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func randomChar() rune {
	return rune(int('a') + (rand.Int() % 26))
}
func randomStr(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += string(randomChar())
	}
	return str
}

const (
	address = "localhost:50051"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		log.Println(args[0])
		log.Println(args[1])
	}
	keyNum, _ := strconv.ParseInt(args[1], 0, 0)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKVDBServiceClient(conn)

	// Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	//defer cancel()

	var wg sync.WaitGroup
	keyBits := 10
	valueBits := 100

	ch := make(chan string, 1000)

	go func() {
		for {
			s := <-ch
			fmt.Println(s)
			if s == "exit" {
				os.Exit(1)
			}
		}
	}()

	for i := 1; i <= int(keyNum); i++ {
		wg.Add(1)
		go func(i int) {
			k := randomStr(keyBits)
			v := randomStr(valueBits)
			start := time.Now()
			_, err := c.Write(context.Background(), &pb.KVRequest{Key: k, Value: v})
			end := time.Now()
			if err != nil {
				log.Printf(
					"could not write, i: %v, key: %v, value: %v, error: %v \n",
					i, k, v, err)
			}
			//log.Printf("success write <%s, %s>", r.Key, r.Value)
			ch <- fmt.Sprintf("put %v-th kv-pair, elapsed seconds: %v ", i, end.Sub(start).Seconds())
			wg.Done()
		}(i)
	}
	wg.Wait()
	ch <- "exit"
}
