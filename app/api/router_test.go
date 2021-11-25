package api_test

import (
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/alexandre-slp/event-broker/mocks"
	"github.com/fasthttp/router"
	"github.com/golang/mock/gomock"
	"github.com/valyala/fasthttp"
	"reflect"
	"testing"
)

func TestNewRequestHandlerService(t *testing.T) {
	t.Parallel()
	rs := api.NewRouterService(router.New())
	type args struct {
		routerService *api.RouterService
	}
	tests := []struct {
		name string
		args args
		want *api.RequestHandlerService
	}{
		{
			name: "Test return value",
			args: args{routerService: rs},
			want: api.NewRequestHandlerService(rs),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.NewRequestHandlerService(tt.args.routerService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequestHandlerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestHandlerService_InitRequestHandler(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rs := api.NewRouterService(router.New())
	mpi := mocks.NewMockPathInitializer(ctrl)

	mpi.
		EXPECT().
		InitPaths(rs).
		Times(3).
		Return()

	rhs := api.NewRequestHandlerService(rs)
	rhs.InitRequestHandler(mpi, mpi, mpi)
}

func TestNewRouterService(t *testing.T) {
	t.Parallel()
	type args struct {
		router *router.Router
	}
	tests := []struct {
		name string
		args args
		want *api.RouterService
	}{
		{
			name: "Test new router service",
			args: args{},
			want: &api.RouterService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.NewRouterService(tt.args.router); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouterService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouterService_GET(t *testing.T) {
	t.Parallel()
	type fields struct {
		r *router.Router
	}
	type args struct {
		path    string
		handler func(*fasthttp.RequestCtx)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test GET",
			fields: fields{
				r: router.New(),
			},
			args: args{
				path:    "/test/get",
				handler: func(*fasthttp.RequestCtx) {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := api.NewRouterService(tt.fields.r)
			rs.GET(tt.args.path, tt.args.handler)
		})
	}
}

func TestRouterService_POST(t *testing.T) {
	t.Parallel()
	type fields struct {
		r *router.Router
	}
	type args struct {
		path    string
		handler func(*fasthttp.RequestCtx)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test POST",
			fields: fields{
				r: router.New(),
			},
			args: args{
				path:    "/test/post",
				handler: func(*fasthttp.RequestCtx) {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := api.NewRouterService(tt.fields.r)
			rs.POST(tt.args.path, tt.args.handler)
		})
	}
}

func TestRouterService_PUT(t *testing.T) {
	t.Parallel()
	type fields struct {
		r *router.Router
	}
	type args struct {
		path    string
		handler func(*fasthttp.RequestCtx)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test PUT",
			fields: fields{
				r: router.New(),
			},
			args: args{
				path:    "/test/put",
				handler: func(*fasthttp.RequestCtx) {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := api.NewRouterService(tt.fields.r)
			rs.PUT(tt.args.path, tt.args.handler)
		})
	}
}

func TestRouterService_PATCH(t *testing.T) {
	t.Parallel()
	type fields struct {
		r *router.Router
	}
	type args struct {
		path    string
		handler func(*fasthttp.RequestCtx)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test PATCH",
			fields: fields{
				r: router.New(),
			},
			args: args{
				path:    "/test/patch",
				handler: func(*fasthttp.RequestCtx) {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := api.NewRouterService(tt.fields.r)
			rs.PATCH(tt.args.path, tt.args.handler)
		})
	}
}

func TestRouterService_DELETE(t *testing.T) {
	t.Parallel()
	type fields struct {
		r *router.Router
	}
	type args struct {
		path    string
		handler func(*fasthttp.RequestCtx)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test DELETE",
			fields: fields{
				r: router.New(),
			},
			args: args{
				path:    "/test/delete",
				handler: func(*fasthttp.RequestCtx) {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := api.NewRouterService(tt.fields.r)
			rs.DELETE(tt.args.path, tt.args.handler)
		})
	}
}

func TestRouterService_Handler(t *testing.T) {
	t.Parallel()
	type fields struct {
		r *router.Router
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test Handler",
			fields: fields{
				r: router.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := api.NewRouterService(tt.fields.r)
			rs.Handler()
		})
	}
}

func TestRouterService_Group(t *testing.T) {
	t.Parallel()
	type fields struct {
		r *router.Router
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *router.Group
	}{
		{
			name:   "Test Group",
			fields: fields{
				r: router.New(),
			},
			args:   args{
				path: "/group",
			},
			want: router.New().Group("/group"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := api.NewRouterService(tt.fields.r)
			if got := rs.Group(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group() = %v, want %v", got, tt.want)
			}
		})
	}
}
