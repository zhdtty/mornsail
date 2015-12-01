package main

import (
	//	"fmt"
	"servlet"
)

type ServerResource struct {
	psm *servlet.PlayerSessionManager
}

func NewServerResource() *ServerResource {
	sr := &ServerResource{
		psm: servlet.NewPlayerSessionManager(),
	}
	return sr
}

func (sr *ServerResource) GetPlayerSessionManager() *servlet.PlayerSessionManager {
	return sr.psm
}

var g_resource *ServerResource = NewServerResource()
