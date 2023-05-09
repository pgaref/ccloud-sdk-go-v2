// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_costs_billing_v1.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_billing_v1 "github.com/confluentinc/ccloud-sdk-go-v2/billing/v1"
)

// CostsBillingV1Api is a mock of CostsBillingV1Api interface
type CostsBillingV1Api struct {
	lockListBillingV1Costs sync.Mutex
	ListBillingV1CostsFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.ApiListBillingV1CostsRequest

	lockListBillingV1CostsExecute sync.Mutex
	ListBillingV1CostsExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.ApiListBillingV1CostsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.BillingV1CostList, *net_http.Response, error)

	calls struct {
		ListBillingV1Costs []struct {
			Ctx context.Context
		}
		ListBillingV1CostsExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.ApiListBillingV1CostsRequest
		}
	}
}

// ListBillingV1Costs mocks base method by wrapping the associated func.
func (m *CostsBillingV1Api) ListBillingV1Costs(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.ApiListBillingV1CostsRequest {
	m.lockListBillingV1Costs.Lock()
	defer m.lockListBillingV1Costs.Unlock()

	if m.ListBillingV1CostsFunc == nil {
		panic("mocker: CostsBillingV1Api.ListBillingV1CostsFunc is nil but CostsBillingV1Api.ListBillingV1Costs was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListBillingV1Costs = append(m.calls.ListBillingV1Costs, call)

	return m.ListBillingV1CostsFunc(ctx)
}

// ListBillingV1CostsCalled returns true if ListBillingV1Costs was called at least once.
func (m *CostsBillingV1Api) ListBillingV1CostsCalled() bool {
	m.lockListBillingV1Costs.Lock()
	defer m.lockListBillingV1Costs.Unlock()

	return len(m.calls.ListBillingV1Costs) > 0
}

// ListBillingV1CostsCalls returns the calls made to ListBillingV1Costs.
func (m *CostsBillingV1Api) ListBillingV1CostsCalls() []struct {
	Ctx context.Context
} {
	m.lockListBillingV1Costs.Lock()
	defer m.lockListBillingV1Costs.Unlock()

	return m.calls.ListBillingV1Costs
}

// ListBillingV1CostsExecute mocks base method by wrapping the associated func.
func (m *CostsBillingV1Api) ListBillingV1CostsExecute(r github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.ApiListBillingV1CostsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.BillingV1CostList, *net_http.Response, error) {
	m.lockListBillingV1CostsExecute.Lock()
	defer m.lockListBillingV1CostsExecute.Unlock()

	if m.ListBillingV1CostsExecuteFunc == nil {
		panic("mocker: CostsBillingV1Api.ListBillingV1CostsExecuteFunc is nil but CostsBillingV1Api.ListBillingV1CostsExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.ApiListBillingV1CostsRequest
	}{
		R: r,
	}

	m.calls.ListBillingV1CostsExecute = append(m.calls.ListBillingV1CostsExecute, call)

	return m.ListBillingV1CostsExecuteFunc(r)
}

// ListBillingV1CostsExecuteCalled returns true if ListBillingV1CostsExecute was called at least once.
func (m *CostsBillingV1Api) ListBillingV1CostsExecuteCalled() bool {
	m.lockListBillingV1CostsExecute.Lock()
	defer m.lockListBillingV1CostsExecute.Unlock()

	return len(m.calls.ListBillingV1CostsExecute) > 0
}

// ListBillingV1CostsExecuteCalls returns the calls made to ListBillingV1CostsExecute.
func (m *CostsBillingV1Api) ListBillingV1CostsExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_billing_v1.ApiListBillingV1CostsRequest
} {
	m.lockListBillingV1CostsExecute.Lock()
	defer m.lockListBillingV1CostsExecute.Unlock()

	return m.calls.ListBillingV1CostsExecute
}

// Reset resets the calls made to the mocked methods.
func (m *CostsBillingV1Api) Reset() {
	m.lockListBillingV1Costs.Lock()
	m.calls.ListBillingV1Costs = nil
	m.lockListBillingV1Costs.Unlock()
	m.lockListBillingV1CostsExecute.Lock()
	m.calls.ListBillingV1CostsExecute = nil
	m.lockListBillingV1CostsExecute.Unlock()
}
