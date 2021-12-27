package auth

import (
	"github.com/ValeryBMSTU/evoModeler/internal/bl"
	"golang.org/x/net/context"
)

type Server struct {
	Bl bl.Bl
}

func (s *Server) CreateUser(c context.Context, request *CURequest) (response *CUResponse, err error) {
	pass := request.Pass
	login := request.Login

	sessionID, err := s.Bl.CreateUser(login, pass)

	response = &CUResponse{
		UserId: int64(sessionID),
	}

	return response, nil
}

func (s *Server) CreateSession(c context.Context, request *CSRequest) (response *CSResponse, err error) {
	pass := request.Pass
	login := request.Login

	sessionID, err := s.Bl.CreateSession(login, pass)

	response = &CSResponse{
		SessionId: int64(sessionID),
	}

	return response, nil
}

func (s *Server) RemoveSession(c context.Context, request *RSRequest) (response *RSResponse, err error) {
	err = s.Bl.RemoveSession(int(request.SessionId))

	response = &RSResponse{
		IsErr: err == nil,
	}

	return response, nil
}

func (s *Server) mustEmbedUnimplementedAuthServer() {

}
