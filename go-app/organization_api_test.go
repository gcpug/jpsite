package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/favclip/ucon"
)

func TestOrganizationAPI_List(t *testing.T) {
	resp := ucon.MakeHandlerTestBed(t, "GET", "/api/1/organization", nil)
	if e, g := http.StatusOK, resp.StatusCode; e != g {
		t.Errorf("expected StatusCode %v, got %v", e, g)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	apiResp := &OrganizationAPIListResponse{}
	if err := json.Unmarshal(body, apiResp); err != nil {
		t.Errorf("unexpected ResponseBody.err=%+v", err)
	}
}
