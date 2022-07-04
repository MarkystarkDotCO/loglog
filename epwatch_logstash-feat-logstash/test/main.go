package main

import (
	"fmt"
	"time"
)

func testPassByValue(a []string) {
	fmt.Println(a)
	a = append(a, "b")
	fmt.Println(a)

}
func main() {
	fmt.Println(time.Now().Format(time.RFC3339Nano))
	var a []string
	a = append(a, "a")
	fmt.Println(a)
	testPassByValue(a)
	fmt.Println(a)

	// t := time.Now().Format("2006-01-02")
	// a, _ := time.Parse("2006-01-02", t)
	// fmt.Println(a)
	// timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	// fmt.Println(timestamp)
	// max("TrendMicro-2022-06-23.log")
	// L:
	// 	for {
	// 	R:
	// 		for {
	// 			for {
	// 				break R
	// 			}
	// 		}
	// 		fmt.Println("break R")
	// 		break L
	// 	}
	// 	fmt.Println("break L")
	// f, err := os.Open("/mnt/shared-log/logstash/TrendMicro-2022-06-21.log")
	// if err != nil {
	// 	log.Println(err)
	// }

	// defer f.Close()

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.)
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://pokpong:pokpongmongo@10.10.21.58:27017/dev"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancle()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)

	// /*
	//    List databases
	// */
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)
	// collection := client.Database("dev").Collection("qradarlogs")

	// t, err := tail.TailFile(
	// 	"/mnt/shared-log/logstash/trendmicro/TrendMicro-2022-06-24.log", tail.Config{Follow: true, ReOpen: true})
	// if err != nil {
	// 	panic(err)
	// }
	// var jsonMap interface{}

	// // Print the text of each received line
	// for line := range t.Lines {
	// 	json.Unmarshal([]byte(line.Text), &jsonMap)
	// 	// fmt.Println(jsonMap)
	// 	_, insertErr := collection.InsertOne(ctx, jsonMap)
	// 	if insertErr != nil {
	// 		log.Fatal(insertErr)
	// 	}
	// 	break
	// }

}
