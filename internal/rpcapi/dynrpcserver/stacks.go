// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by ./generator. DO NOT EDIT.
package dynrpcserver

import (
	"context"
	"sync"

	tf1 "github.com/hashicorp/terraform/internal/rpcapi/terraform1"
)

type Stacks struct {
	impl tf1.StacksServer
	mu   sync.RWMutex
}

var _ tf1.StacksServer = (*Stacks)(nil)

func NewStacksStub() *Stacks {
	return &Stacks{}
}

func (s *Stacks) ApplyStackChanges(a0 *tf1.ApplyStackChanges_Request, a1 tf1.Stacks_ApplyStackChangesServer) error {
	impl, err := s.realRPCServer()
	if err != nil {
		return err
	}
	return impl.ApplyStackChanges(a0, a1)
}

func (s *Stacks) ClosePlan(a0 context.Context, a1 *tf1.CloseStackPlan_Request) (*tf1.CloseStackPlan_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.ClosePlan(a0, a1)
}

func (s *Stacks) CloseStackConfiguration(a0 context.Context, a1 *tf1.CloseStackConfiguration_Request) (*tf1.CloseStackConfiguration_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.CloseStackConfiguration(a0, a1)
}

func (s *Stacks) CloseState(a0 context.Context, a1 *tf1.CloseStackState_Request) (*tf1.CloseStackState_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.CloseState(a0, a1)
}

func (s *Stacks) FindStackConfigurationComponents(a0 context.Context, a1 *tf1.FindStackConfigurationComponents_Request) (*tf1.FindStackConfigurationComponents_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.FindStackConfigurationComponents(a0, a1)
}

func (s *Stacks) InspectExpressionResult(a0 context.Context, a1 *tf1.InspectExpressionResult_Request) (*tf1.InspectExpressionResult_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.InspectExpressionResult(a0, a1)
}

func (s *Stacks) OpenPlan(a0 tf1.Stacks_OpenPlanServer) error {
	impl, err := s.realRPCServer()
	if err != nil {
		return err
	}
	return impl.OpenPlan(a0)
}

func (s *Stacks) OpenStackConfiguration(a0 context.Context, a1 *tf1.OpenStackConfiguration_Request) (*tf1.OpenStackConfiguration_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.OpenStackConfiguration(a0, a1)
}

func (s *Stacks) OpenStackInspector(a0 context.Context, a1 *tf1.OpenStackInspector_Request) (*tf1.OpenStackInspector_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.OpenStackInspector(a0, a1)
}

func (s *Stacks) OpenState(a0 tf1.Stacks_OpenStateServer) error {
	impl, err := s.realRPCServer()
	if err != nil {
		return err
	}
	return impl.OpenState(a0)
}

func (s *Stacks) PlanStackChanges(a0 *tf1.PlanStackChanges_Request, a1 tf1.Stacks_PlanStackChangesServer) error {
	impl, err := s.realRPCServer()
	if err != nil {
		return err
	}
	return impl.PlanStackChanges(a0, a1)
}

func (s *Stacks) ValidateStackConfiguration(a0 context.Context, a1 *tf1.ValidateStackConfiguration_Request) (*tf1.ValidateStackConfiguration_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.ValidateStackConfiguration(a0, a1)
}

func (s *Stacks) ActivateRPCServer(impl tf1.StacksServer) {
	s.mu.Lock()
	s.impl = impl
	s.mu.Unlock()
}

func (s *Stacks) realRPCServer() (tf1.StacksServer, error) {
	s.mu.RLock()
	impl := s.impl
	s.mu.RUnlock()
	if impl == nil {
		return nil, unavailableErr
	}
	return impl, nil
}
