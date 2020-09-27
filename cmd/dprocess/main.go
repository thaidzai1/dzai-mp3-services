package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/thaidzai285/dzai-mp3-protobuf/pkg/pb"
	"github.com/thaidzai285/dzai-mp3-services/configs"
	"github.com/thaidzai285/dzai-mp3-services/internal/pkg/dprocess"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

var (
	ctxCancel context.CancelFunc
	ctx       context.Context
	flConfigs = flag.String("config-file", "", "Path to list songs want to download")
)

func main() {
	flag.Parse()
	ctx, ctxCancel = context.WithCancel(context.Background())

	go func() {
		err := serverGRPC()
		if err != nil {
			log.Fatal("GRPC server error: ", err)
			return
		}
	}()

	<-ctx.Done()
}

func serverGRPC() error {
	defer ctxCancel()

	log.Println("ahoy listen")
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		return err
	}

	log.Println("ahoy load config")
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	log.Println("ahoy start Dial", cfg)
	connCrawler, err := grpc.Dial(cfg.Dcrawl.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	dzaiCrawler := pb.NewCrawlerClient(connCrawler)
	if err != nil {
		return err
	}

	songService := dprocess.NewSongService(dzaiCrawler)

	s := grpc.NewServer()
	pb.RegisterSongsServiceServer(s, songService)

	log.Println("GRPC listen on 3001")
	return s.Serve(listen)
}

func loadConfig() (configs.Config, error) {
	var cfg configs.Config
	log.Println("flConfigs", *flConfigs)
	if *flConfigs == "" {
		cfg = configs.DefaultConfig()
		return cfg, nil
	}

	byteCfgSchema, err := ioutil.ReadFile(*flConfigs)
	log.Println("schema", string(byteCfgSchema))
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(byteCfgSchema, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
