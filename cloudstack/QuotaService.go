//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"net/url"
)

type QuotaIsEnabledParams struct {
	p map[string]interface{}
}

func (p *QuotaIsEnabledParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new QuotaIsEnabledParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaIsEnabledParams() *QuotaIsEnabledParams {
	p := &QuotaIsEnabledParams{}
	p.p = make(map[string]interface{})
	return p
}

// Return true if the plugin is enabled
func (s *QuotaService) QuotaIsEnabled(p *QuotaIsEnabledParams) (*QuotaIsEnabledResponse, error) {
	resp, err := s.cs.newRequest("quotaIsEnabled", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaIsEnabledResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaIsEnabledResponse struct {
	Isenabled bool `json:"isenabled"`
}

// Quota Balance
type QuotaBalanceParams struct {
	p map[string]interface{}
}

func (p *QuotaBalanceParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *QuotaBalanceParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

// Not required
func (p *QuotaBalanceParams) SetStartDate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
	return
}

// Not required
func (p *QuotaBalanceParams) SetEndDate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
	return
}

func (p *QuotaBalanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	for key, value := range p.p {
		u.Set(key, value.(string))
	}
	return u
}

// Return balance // Only simple request
func (s *QuotaService) QuotaBalance(p *QuotaBalanceParams) (*QuotaBalanceResponse, error) {
	resp, err := s.cs.newRequest("quotaBalance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaBalanceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaBalanceResponse struct {
	Balance QuotaBalance `json:"balance"`
}

type QuotaBalance struct {
	StartQuota float64 `json:"startquota"`
	EndQuota float64 `json:"endquota"`
	Credits []QuotaCredits `json:"credits"`
	Startdate string `json:"startdate"`
	Enddate string `json:"enddate"`
	Currency string `json:"currency"`
}

type QuotaCredits struct {
	Credits float64 `json:"credits"`
	UpdatedOon string `json:"updated_on"`
}


// Quota Summary - Undocumented API, return not Valid Id's ======================================
type QuotaSummaryParams struct {
	p map[string]interface{}
}

// Not required
func (p *QuotaSummaryParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

// Not required
func (p *QuotaSummaryParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *QuotaSummaryParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	for key, value := range p.p {
		u.Set(key, value.(string))
	}
	return u
}

// Return balance // Only simple request
func (s *QuotaService) QuotaSummary(p *QuotaSummaryParams) (*QuotaSummaryResponse, error) {
	resp, err := s.cs.newRequest("quotaSummary", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaSummaryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaSummaryResponse struct {
	Count     int64          `json:"count"`
	Summaries []QuotaSummary `json:"summary"`
}
// This is not valid id's!
type QuotaSummary struct {
	AccountId int64 `json:"accountid"`
	Account string `json:"account"`
	DomainId int64 `json:"domainid"`
	Domain string `json:"domain"`
	Balance float64 `json:"balance"`
	State string `json:"state"`
	Quota float64 `json:"quota"`
	Startdate string `json:"startdate"`
	Enddate string `json:"enddate"`
	Currency string `json:"currency"`
}
