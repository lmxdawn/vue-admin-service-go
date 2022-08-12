// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.5
// source: v1/ui_admin.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUiAdminCreateAdmin = "/api.ui.admin.v1.UiAdmin/CreateAdmin"

type UiAdminHTTPServer interface {
	CreateAdmin(context.Context, *CreateAdminRequest) (*CreateAdminReply, error)
}

func RegisterUiAdminHTTPServer(s *http.Server, srv UiAdminHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/admin", _UiAdmin_CreateAdmin0_HTTP_Handler(srv))
}

func _UiAdmin_CreateAdmin0_HTTP_Handler(srv UiAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAdminRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUiAdminCreateAdmin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAdmin(ctx, req.(*CreateAdminRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAdminReply)
		return ctx.Result(200, reply)
	}
}

type UiAdminHTTPClient interface {
	CreateAdmin(ctx context.Context, req *CreateAdminRequest, opts ...http.CallOption) (rsp *CreateAdminReply, err error)
}

type UiAdminHTTPClientImpl struct {
	cc *http.Client
}

func NewUiAdminHTTPClient(client *http.Client) UiAdminHTTPClient {
	return &UiAdminHTTPClientImpl{client}
}

func (c *UiAdminHTTPClientImpl) CreateAdmin(ctx context.Context, in *CreateAdminRequest, opts ...http.CallOption) (*CreateAdminReply, error) {
	var out CreateAdminReply
	pattern := "/admin/v1/admin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUiAdminCreateAdmin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
