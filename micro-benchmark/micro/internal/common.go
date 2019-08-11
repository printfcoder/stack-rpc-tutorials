package internal

import (
	"context"
	"flag"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/micro-in-cn/tutorials/micro-benchmark/pb"
	"github.com/micro/go-micro/util/log"
	"github.com/montanaflynn/stats"
)

var delay = flag.Duration("delay", 0, "delay to mock business processing")

type HelloS struct{}

func (t *HelloS) Say(ctx context.Context, args *pb.BenchmarkMessage, reply *pb.BenchmarkMessage) error {
	s := "OK"
	var i int32 = 100
	args.Field1 = s
	args.Field2 = i
	*reply = *args
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return nil
}

func PrepareArgs() *pb.BenchmarkMessage {
	b := true
	var i int32 = 100000
	var i64 int64 = 100000
	var s = "谢谢前人的脚本与程序，用起来是那么飘逸与爽朗"

	var args pb.BenchmarkMessage

	v := reflect.ValueOf(&args).Elem()
	num := v.NumField()
	for k := 0; k < num; k++ {
		field := v.Field(k)
		if field.Type().Kind() == reflect.Ptr {
			switch v.Field(k).Type().Elem().Kind() {
			case reflect.Int, reflect.Int32:
				field.Set(reflect.ValueOf(&i))
			case reflect.Int64:
				field.Set(reflect.ValueOf(&i64))
			case reflect.Bool:
				field.Set(reflect.ValueOf(&b))
			case reflect.String:
				field.Set(reflect.ValueOf(&s))
			}
		} else {
			switch field.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				field.SetInt(100000)
			case reflect.Bool:
				field.SetBool(true)
			case reflect.String:
				field.SetString(s)
			}
		}

	}
	return &args
}

func ClientRun(m, n int, c pb.HelloService) {
	selected := -1

	log.Infof("concurrency: %d\nrequests per client: %d\n\n", n, m)

	args := PrepareArgs()

	b, _ := proto.Marshal(args)
	log.Infof("message size: %d bytes\n\n", len(b))

	var wg sync.WaitGroup
	wg.Add(n * m)

	var trans uint64
	var transOK uint64

	d := make([][]int64, n, n)

	// it contains warmup time but we can ignore it
	totalT := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		dt := make([]int64, 0, m)
		d = append(d, dt)
		selected = selected + 1

		go func(i int, selected int) {

			//warmup
			for j := 0; j < 5; j++ {
				c.Say(context.Background(), args)
			}

			for j := 0; j < m; j++ {
				t := time.Now().UnixNano()
				reply, err := c.Say(context.Background(), args)
				t = time.Now().UnixNano() - t

				d[i] = append(d[i], t)

				if err == nil && reply.Field1 == "OK" {
					atomic.AddUint64(&transOK, 1)
				}

				atomic.AddUint64(&trans, 1)
				wg.Done()
			}

		}(i, selected)

	}

	wg.Wait()
	totalT = time.Now().UnixNano() - totalT
	totalT = totalT / 1000000
	log.Infof("took %d ms for %d requests\n", totalT, n*m)

	totalD := make([]int64, 0, n*m)
	for _, k := range d {
		totalD = append(totalD, k...)
	}
	totalD2 := make([]float64, 0, n*m)
	for _, k := range totalD {
		totalD2 = append(totalD2, float64(k))
	}

	mean, _ := stats.Mean(totalD2)
	median, _ := stats.Median(totalD2)
	max, _ := stats.Max(totalD2)
	min, _ := stats.Min(totalD2)
	p99, _ := stats.Percentile(totalD2, 99.9)
	p90, _ := stats.Percentile(totalD2, 90)
	tps := int64(n*m) * 1000 / totalT

	log.Infof("sent     requests    : %d\n", n*m)
	log.Infof("received requests    : %d\n", atomic.LoadUint64(&trans))
	log.Infof("received requests_OK : %d\n", atomic.LoadUint64(&transOK))
	log.Infof("throughput  (TPS)    : %d\n", int64(n*m)*1000/totalT)
	log.Infof("concurrency\tmean\tmedian\tmax\tmin\tp90\tp99\tTPS\n")
	log.Infof("%d \t%.fns\t%.fns\t%.fns\t%.fns\t%.fns\t%.fns\t%d\n", n, mean, median, max, min, p99, p90, tps)
	log.Infof("%d \t%.3fms\t%.3fms\t%.3fms\t%.3fms\t%.3fms\t%.fms\t%d\n", n, mean/1000000, median/1000000, max/1000000, min/1000000, p90/1000000, p99/1000000, tps)
}
