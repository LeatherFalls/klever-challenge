package controller

import (
	"context"
	"reflect"
	"testing"
	"upvote/grpc/proto/pb"
)

/* var cryptoServiceClient pb.CryptoServiceClient */

/* func init() {
	go StartServer()
}

func connectServer() (*grpc.ClientConn, context.CancelFunc) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error to connect to server: %v", err)
	}

	client := pb.NewCryptoServiceClient(conn)
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	cryptoServiceClient = client

	return conn, cancel
}

func TestCreate(t *testing.T) {
	conn, cancel := connectServer()
	defer conn.Close()
	defer cancel()

	req := &pb.Crypto{
		Id:          "634f6df1f56fbac87d062fbe",
		CryptoName:  "Bitcoin",
		CryptoLikes: 1,
	}

	res, err := cryptoServiceClient.Create(context.Background(), req)

	if err != nil {
		t.Errorf("Error while creating crypto: %v", err)
	}

	if res.Id != req.Id {
		t.Errorf("Error while creating crypto: %v", err)
	}
} */

func TestCryptoServer_Create(t *testing.T) {
	type fields struct {
		CryptoServiceServer pb.CryptoServiceServer
	}
	type args struct {
		ctx context.Context
		in  *pb.Crypto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CryptoId
		wantErr bool
	}{
		{
			name: "Create sucess",
			fields: fields{
				CryptoServiceServer: &CryptoServer{},
			},
			args: args{
				ctx: context.Background(),
				in: &pb.Crypto{
					Id:          "634f6df1f56fbac87d062fbe",
					CryptoName:  "Bitcoin",
					CryptoLikes: 1,
				},
			},
			want: &pb.CryptoId{
				Id: "634f6df1f56fbac87d062fbe",
			},
			wantErr: false,
		},
/* 		{
			name: "Create error",
			fields: fields{
				CryptoServiceServer: &CryptoServer{},
			},
			args: args{
				ctx: context.Background(),
				in: &pb.Crypto{
					Id:          "634f6df1f56fbac87d062fbe",
					CryptoName:  "Bitcoin",
					CryptoLikes: 1,
				},
			},
			want: 	nil,
			wantErr: true,
		}, */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCryptoController()
			got, err := c.Create(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CryptoServer.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
