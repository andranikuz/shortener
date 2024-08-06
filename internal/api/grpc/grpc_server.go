package grpcserver

import (
	"context"
	"fmt"
	"github.com/andranikuz/shortener/internal/container"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"

	pb "github.com/andranikuz/shortener/internal/api/grpc/proto"
	"github.com/andranikuz/shortener/internal/services"
	"github.com/andranikuz/shortener/internal/storage"
)

// GRPCServer implements the gRPC server interface.
type GRPCServer struct {
	pb.UnimplementedShortenerServiceServer
	shortener services.Shortener
	storage   storage.Storage
}

// NewGRPCServer creates a new GRPCServer.
func NewGRPCServer(shortener services.Shortener, storage storage.Storage) *GRPCServer {
	return &GRPCServer{
		shortener: shortener,
		storage:   storage,
	}
}

// GenerateShortURL generate short url.
func (s *GRPCServer) GenerateShortURL(ctx context.Context, req *pb.GenerateShortURLRequest) (*pb.GenerateShortURLResponse, error) {
	shortURL, err := s.shortener.GenerateShortURL(ctx, req.Url, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GenerateShortURLResponse{Result: shortURL}, nil
}

// GenerateShortURLBatch generate bach url.
func (s *GRPCServer) GenerateShortURLBatch(ctx context.Context, req *pb.GenerateShortURLBatchRequest) (*pb.GenerateShortURLBatchResponse, error) {
	var shortenItems []*pb.ShortenItem
	for _, item := range req.Items {
		shortURL, err := s.shortener.GenerateShortURL(ctx, item.OriginalUrl, req.UserId)
		if err != nil {
			return nil, err
		}
		shortenItems = append(shortenItems, &pb.ShortenItem{
			CorrelationId: item.CorrelationId,
			ShortUrl:      shortURL,
		})
	}
	return &pb.GenerateShortURLBatchResponse{Items: shortenItems}, nil
}

// GetFullURL handle get full url.
func (s *GRPCServer) GetFullURL(ctx context.Context, req *pb.GetFullURLRequest) (*pb.GetFullURLResponse, error) {
	fullURL, err := s.shortener.GetFullURL(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetFullURLResponse{FullUrl: fullURL}, nil
}

// Ping handle ping.
func (s *GRPCServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	err := s.storage.Ping()
	if err != nil {
		log.Info().Msg(fmt.Sprintf("postgres ping error: %s", err.Error()))
		return nil, err
	}

	return &pb.PingResponse{Message: "pong"}, nil
}

// GetUserURLs handle get user urls.
func (s *GRPCServer) GetUserURLs(ctx context.Context, req *pb.GetUserURLsRequest) (*pb.GetUserURLsResponse, error) {
	urls, err := s.shortener.GetUserURLs(ctx, req.UserId)
	var items []*pb.GetUserURLsHandlerItem
	if err != nil {
		return nil, err
	}
	for _, url := range urls {
		items = append(items, &pb.GetUserURLsHandlerItem{ShortUrl: url.GetShorter(), OriginalUrl: url.FullURL})
	}
	return &pb.GetUserURLsResponse{Urls: items}, nil
}

// DeleteURLs handle delete urls.
func (s *GRPCServer) DeleteURLs(ctx context.Context, req *pb.DeleteURLsRequest) (*pb.DeleteURLsResponse, error) {
	go s.shortener.DeleteURLs(req.Urls, req.UserId)
	return &pb.DeleteURLsResponse{}, nil
}

// GetInternalStats handle get internal stats.
func (s *GRPCServer) GetInternalStats(ctx context.Context, req *pb.GetInternalStatsRequest) (*pb.GetInternalStatsResponse, error) {
	urls, users, err := s.shortener.GetInternalStats(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetInternalStatsResponse{
		UrlCount:  urls,
		UserCount: users,
	}, nil
}

// StartGRPCServer starts the gRPC server.
func StartGRPCServer(addr string, cnt *container.Container) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	shortener, _ := cnt.Shortener()
	storage, _ := cnt.Storage()

	pb.RegisterShortenerServiceServer(grpcServer, NewGRPCServer(shortener, storage))
	return grpcServer.Serve(lis)
}
