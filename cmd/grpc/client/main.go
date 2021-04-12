package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	pb "github.com/tengla/nsr-parser/trainroute/protos"
	"google.golang.org/grpc"
)

var (
	id = flag.String("id", "", "Id of train")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	client := pb.NewOperatingTrainRouteServiceClient(conn)
	if len(*id) > 0 {
		ids := strings.Split(*id, ",")
		fmt.Printf("Fetching ids %s\n", strings.Join(ids, ","))
		list := &pb.OtrIdList{}
		for _, _id := range ids {
			list.TrainIds = append(list.TrainIds, &pb.OtrId{
				TrainId: _id,
			})
		}
		item, err := client.GetItems(context.Background(), list)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(item)
	} else {
		list, err := client.FullState(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("%v ", list.TrainIds)
	}
}
