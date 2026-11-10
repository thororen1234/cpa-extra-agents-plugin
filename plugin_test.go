package main

import (
	"encoding/json"
	"testing"

	"github.com/MoeForge/moego/ojson"
	"github.com/router-for-me/CLIProxyAPI/v8/sdk/pluginapi"
)

func intercept(t *testing.T, model, toFormat, body string) pluginapi.RequestInterceptResponse {
	t.Helper()
	raw, _ := json.Marshal(pluginapi.RequestInterceptRequest{Model: model, ToFormat: toFormat, Body: []byte(body)})
	out, err := interceptAfter(raw)
	if err != nil {
		t.Fatal(err)
	}
	var env envelope
	if err := json.Unmarshal(out, &env); err != nil || !env.OK {
		t.Fatalf("bad envelope: %s %v", out, err)
	}
	var resp pluginapi.RequestInterceptResponse
	if err := json.Unmarshal(env.Result, &resp); err != nil {
		t.Fatal(err)
	}
	return resp
}

func TestModelRegistration(t *testing.T) {
	out, err := handleMethod("model.register", nil)
	if err != nil {
		t.Fatal(err)
	}
	var env envelope
	_ = json.Unmarshal(out, &env)
	var resp pluginapi.ModelRegistrationResponse
	if err := json.Unmarshal(env.Result, &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Provider != "claude" || len(resp.Models) != 4 {
		t.Fatalf("got provider=%q models=%d", resp.Provider, len(resp.Models))
	}
	if resp.Models[3].ID != "claude-mythos-5-2" || len(resp.Models[3].Thinking.Levels) != 5 {
		t.Fatalf("unexpected model: %+v", resp.Models[3])
	}
}

func TestMythos52AddsBetasAndDisplay(t *testing.T) {
	resp := intercept(t, "claude-mythos-5-2[1m]", "claude", `{"thinking":{"type":"adaptive"},"betas":["x"]}`)
	if len(resp.Body) == 0 {
		t.Fatal("expected body")
	}
	got := ojson.GetBytes(resp.Body, "betas").Raw
	want := `["x","per-turn-control-2026-07-01","mid-conversation-system-clear-at-2026-08-21","thinking-binding-controls-2026-08-01"]`
	if got != want {
		t.Fatalf("betas = %s", got)
	}
	if ojson.GetBytes(resp.Body, "thinking.display").String() != "updates" {
		t.Fatalf("display not set: %s", resp.Body)
	}
}

func TestMythos51OnlyPerTurnControl(t *testing.T) {
	resp := intercept(t, "claude-mythos-5-1", "claude", `{"thinking":{"type":"adaptive"}}`)
	if ojson.GetBytes(resp.Body, "betas").Raw != `["per-turn-control-2026-07-01"]` {
		t.Fatalf("body = %s", resp.Body)
	}
	if ojson.GetBytes(resp.Body, "thinking.display").Exists() {
		t.Fatal("display must not be set for 5.1")
	}
}

func TestCallerDisplayKept(t *testing.T) {
	resp := intercept(t, "claude-mythos-5-2", "claude", `{"thinking":{"type":"enabled","display":"summarized"}}`)
	if ojson.GetBytes(resp.Body, "thinking.display").String() != "summarized" {
		t.Fatalf("display overwritten: %s", resp.Body)
	}
}

func TestNoOp(t *testing.T) {
	for _, c := range []struct{ model, format string }{
		{"claude-mythos-5", "claude"},
		{"claude-opus-5-5", "claude"},
		{"claude-mythos-5-2", "openai"},
	} {
		if resp := intercept(t, c.model, c.format, `{"thinking":{"type":"adaptive"}}`); len(resp.Body) != 0 {
			t.Fatalf("%v: unexpected body %s", c, resp.Body)
		}
	}
}
