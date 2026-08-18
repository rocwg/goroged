package dashboard

import (
	"context"

	"github.com/rocwg/ged/orchestration"
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"

	edgeclients "github.com/rocwg/ged/examples/edge/clients"
)

type Service struct {
	dictArea dictv1.DictAreaServiceClient
	hello    hellov1.HelloServiceClient
}

func NewService(
	clients *edgeclients.Clients,
) *Service {
	return &Service{
		dictArea: clients.DictArea,
		hello:    clients.Hello,
	}
}

func (s *Service) overview(
	ctx context.Context,
) (*OverviewResponse, error) {

	var (
		dictResp  *dictv1.GetAreaByParentResponse
		helloResp *hellov1.SayHelloResponse
	)

	err := orchestration.ParallelFailFast(
		ctx,

		// ---------------------------------------------------------
		// 1. DictArea
		// ---------------------------------------------------------

		func(ctx context.Context) error {
			var err error
			dictResp, err = s.dictArea.GetAreaByParent(
				ctx,
				&dictv1.GetAreaByParentRequest{
					ParentCode: "0",
				},
			)
			return err
		},

		// ---------------------------------------------------------
		// 2. Hello
		// ---------------------------------------------------------

		func(ctx context.Context) error {
			var err error
			helloResp, err = s.hello.SayHello(
				ctx,
				&hellov1.SayHelloRequest{
					Name: "Goro-Edge",
				},
			)
			return err
		},
	)

	if err != nil {
		return nil, err
	}

	return &OverviewResponse{
		Areas:   dictResp.GetList(),
		Message: helloResp.GetMessage(),
	}, nil
}
