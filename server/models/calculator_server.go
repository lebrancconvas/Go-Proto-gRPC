package models 

import (
	context "context"
	"github.com/lebrancconvas/Go-Proto-gRPC/proto/gen/services/proto"
)

type CalculatorServer struct {

}

func NewCalculatorServer() CalculatorServer {
	return CalculatorServer{}; 
}

func (CalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (CalculatorServer) Add(ctx context.Context, req *Request) (*Response, error) {
	result := req.Number1 + req.Number2; 
	res := Response{
		Result: result,
	}
	return &res, nil; 
}