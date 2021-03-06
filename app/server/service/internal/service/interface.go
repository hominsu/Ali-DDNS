package service

import (
	v1 "Ali-DDNS/api/server/service/v1"
	"Ali-DDNS/app/server/service/internal/biz"
	"Ali-DDNS/app/server/service/jwt"
	"Ali-DDNS/pkg"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// Register a user
func (s *DDNSInterfaceService) Register(ctx context.Context, in *v1.RegisterReq) (*v1.RegisterReply, error) {
	// get the username and password from http request header
	username := in.GetUser().GetUsername()
	password := in.GetUser().GetPassword()

	// Check whether the username and password are empty
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username or password should not be empty")
	}

	// check whether the username already exist
	exists, err := s.domainUserUsecase.IsUserExists(context.TODO(), &biz.DomainUser{Username: username})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	if exists {
		return nil, status.Errorf(codes.AlreadyExists, "User is registered")
	} else {
		// add the user to data repo
		if _, err := s.domainUserUsecase.AddUser(context.TODO(), &biz.DomainUser{
			Username: username,
			Password: password,
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}

		// returns the message and redirects to the login page
		return &v1.RegisterReply{Status: true}, nil
	}
}

// Cancel a user
func (s *DDNSInterfaceService) Cancel(ctx context.Context, in *v1.CancelReq) (*v1.CancelReply, error) {
	// get the username from http request header
	username := in.GetUsername()
	if strings.Trim(username, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username should not be empty")
	}

	// check the username
	claimsUsername, err := jwt.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	if claimsUsername != username {
		return nil, status.Errorf(codes.Unauthenticated, "username not equal token's username")
	}

	// check whether the username already exist
	exists, err := s.domainUserUsecase.IsUserExists(context.TODO(), &biz.DomainUser{Username: username})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	if exists {
		if _, err := s.domainUserUsecase.DelUser(context.TODO(), &biz.DomainUser{
			Username: username,
		}); err != nil {
			return nil, err
		}
		// returns the message and redirects to the login page
		return &v1.CancelReply{Status: true}, nil
	} else {
		return nil, status.Errorf(codes.NotFound, "User is not registered")
	}
}

// Login .
func (s *DDNSInterfaceService) Login(ctx context.Context, in *v1.LoginReq) (*v1.LoginReply, error) {
	username := in.GetUser().GetUsername()
	password := in.GetUser().GetPassword()

	// check whether the username is null, return if null
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username or password should not be empty")
	}

	// check whether the user already exist, return if not
	userExists, err := s.domainUserUsecase.IsUserExists(context.TODO(), &biz.DomainUser{Username: username})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	if !userExists {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// check if the user and password matches, return if it does not exist
	userPassword, err := s.domainUserUsecase.GetUserPassword(context.TODO(), &biz.DomainUser{Username: username})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	if password != userPassword {
		return nil, status.Errorf(codes.PermissionDenied, "Authorized failed")
	}

	token, err := jwt.GenToken(username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &v1.LoginReply{
		Token:    token,
		Username: username,
	}, nil
}

// Logout .
func (s *DDNSInterfaceService) Logout(ctx context.Context, in *v1.LogoutReq) (*v1.LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

// ListDomainName list all domain-checks in redis
func (s *DDNSInterfaceService) ListDomainName(ctx context.Context, in *v1.ListDomainNameReq) (*v1.ListDomainNameReply, error) {
	username := in.GetUsername()

	if strings.Trim(username, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username should not be empty")
	}

	// check the username
	claimsUsername, err := jwt.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	if claimsUsername != username {
		return nil, status.Errorf(codes.Unauthenticated, "username not equal token's username")
	}

	// obtain the domain name of the current user
	domainNames, err := s.domainUserUsecase.GetDomainName(context.TODO(), &biz.DomainUser{Username: username})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	reply := &v1.ListDomainNameReply{}
	if domainNames != nil {
		for _, domainName := range domainNames {
			reply.DomainNames = append(reply.DomainNames, domainName)
		}
	}
	return reply, nil
}

// CreateDomainName create a domain-check
func (s *DDNSInterfaceService) CreateDomainName(ctx context.Context, in *v1.CreateDomainNameReq) (*v1.CreateDomainNameReply, error) {
	username := in.GetUsername()
	domainName := in.GetDomainName()

	if strings.Trim(username, " ") == "" || strings.Trim(domainName, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username or domain_name should not be empty")
	}

	// check the username
	claimsUsername, err := jwt.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	if claimsUsername != username {
		return nil, status.Errorf(codes.Unauthenticated, "username not equal token's username")
	}

	// add the domain name into date repo
	if _, err := s.domainUserUsecase.AddDomainName(context.TODO(), &biz.DomainUser{
		Username:   username,
		DomainName: domainName,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &v1.CreateDomainNameReply{
		Status:     true,
		DomainName: domainName,
	}, nil
}

// DeleteDomainName delete a domain-check
func (s *DDNSInterfaceService) DeleteDomainName(ctx context.Context, in *v1.DeleteDomainNameReq) (*v1.DeleteDomainNameReply, error) {
	username := in.GetUsername()
	domainName := in.GetDomainName()

	if strings.Trim(username, " ") == "" || strings.Trim(domainName, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username or domain_name should not be empty")
	}

	// check the username
	claimsUsername, err := jwt.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	if claimsUsername != username {
		return nil, status.Errorf(codes.Unauthenticated, "username not equal token's username")
	}

	// delete the domain name from data repo
	if _, err := s.domainUserUsecase.DelDomainName(context.TODO(), &biz.DomainUser{
		Username:   username,
		DomainName: domainName,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &v1.DeleteDomainNameReply{
		Status:     true,
		DomainName: domainName,
	}, nil
}

// ListDevice list all devices in redis
func (s *DDNSInterfaceService) ListDevice(ctx context.Context, in *v1.ListDeviceReq) (*v1.ListDeviceReply, error) {
	username := in.GetUsername()

	if strings.Trim(username, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username should not be empty")
	}

	// check the username
	claimsUsername, err := jwt.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	if claimsUsername != username {
		return nil, status.Errorf(codes.Unauthenticated, "username not equal token's username")
	}

	// obtain all devices of the current user from data repo
	devices, err := s.domainUserUsecase.GetDevice(context.TODO(), &biz.DomainUser{
		Username: username,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	reply := &v1.ListDeviceReply{}
	if devices != nil {
		for _, device := range devices {
			reply.Uuid = append(reply.Uuid, device)
		}
	}

	return reply, nil
}

// CreateDevice create a device
func (s *DDNSInterfaceService) CreateDevice(ctx context.Context, in *v1.CreateDeviceReq) (*v1.CreateDeviceReply, error) {
	username := in.GetUsername()

	if strings.Trim(username, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username should not be empty")
	}

	// check the username
	claimsUsername, err := jwt.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	if claimsUsername != username {
		return nil, status.Errorf(codes.Unauthenticated, "username not equal token's username")
	}

	// generate a uuid
	uuid, err := pkg.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	// save the uuid into data repo
	if _, err := s.domainUserUsecase.AddDevice(context.TODO(), &biz.DomainUser{
		Username: username,
		UUID:     uuid,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &v1.CreateDeviceReply{
		Status: true,
		Uuid:   uuid,
	}, nil
}

// DeleteDevice delete a device
func (s *DDNSInterfaceService) DeleteDevice(ctx context.Context, in *v1.DeleteDeviceReq) (*v1.DeleteDeviceReply, error) {
	username := in.GetUsername()
	uuid := in.GetUuid()

	if strings.Trim(username, " ") == "" || strings.Trim(uuid, " ") == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username or uuid should not be empty")
	}

	// check the username
	claimsUsername, err := jwt.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	if claimsUsername != username {
		return nil, status.Errorf(codes.Unauthenticated, "username not equal token's username")
	}

	// delete this device from data repo
	if _, err := s.domainUserUsecase.DelDevice(context.TODO(), &biz.DomainUser{
		Username: username,
		UUID:     uuid,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &v1.DeleteDeviceReply{
		Status: true,
		Uuid:   uuid,
	}, nil
}
