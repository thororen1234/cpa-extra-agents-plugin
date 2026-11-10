package main

import (
	"encoding/json"
	"strings"

	"github.com/MoeForge/moego/ojson"
	"github.com/router-for-me/CLIProxyAPI/v8/sdk/pluginapi"
)

const (
	perTurnControlBeta  = "per-turn-control-2026-07-01"
	thinkingBindingBeta = "thinking-binding-controls-2026-08-01"
	midConvClearAtBeta  = "mid-conversation-system-clear-at-2026-08-21"
)

func canonicalModel(model string) string {
	model = strings.ToLower(strings.TrimSpace(model))
	if i := strings.LastIndexByte(model, '/'); i >= 0 {
		model = model[i+1:]
	}
	return model
}

// perTurnEffort matches Mythos 5.1 and 5.2 with any suffix.
func perTurnEffort(model string) bool {
	return strings.HasPrefix(model, "claude-mythos-5-1") ||
		strings.HasPrefix(model, "claude-mythos-5-2") ||
		strings.HasPrefix(model, "claude-mythos-5.2")
}

func progressDisplay(model string) bool {
	for _, f := range []string{"claude-mythos-5-2", "claude-mythos-5.2"} {
		if model == f || strings.HasPrefix(model, f+"-") || strings.HasPrefix(model, f+"[") {
			return true
		}
	}
	return false
}

func interceptAfter(raw []byte) ([]byte, error) {
	var req pluginapi.RequestInterceptRequest
	if err := json.Unmarshal(raw, &req); err != nil {
		return nil, err
	}
	body := req.Body
	if !strings.EqualFold(req.ToFormat, "claude") || len(body) == 0 {
		return okEnvelope(pluginapi.RequestInterceptResponse{})
	}
	model := canonicalModel(req.Model)
	display := progressDisplay(model)
	thinkingType := strings.ToLower(strings.TrimSpace(ojson.GetBytes(body, "thinking.type").String()))

	var add []string
	if perTurnEffort(model) {
		add = append(add, perTurnControlBeta)
	}
	if display {
		add = append(add, midConvClearAtBeta)
		if thinkingType == "adaptive" {
			add = append(add, thinkingBindingBeta)
		}
	}
	if len(add) == 0 {
		return okEnvelope(pluginapi.RequestInterceptResponse{})
	}

	body, err := addBetas(body, add)
	if err != nil {
		return nil, err
	}
	if display && (thinkingType == "adaptive" || thinkingType == "enabled") &&
		!ojson.GetBytes(body, "thinking.display").Exists() {
		if body, err = ojson.SetBytes(body, "thinking.display", "updates"); err != nil {
			return nil, err
		}
	}
	return okEnvelope(pluginapi.RequestInterceptResponse{Body: body})
}

// addBetas merges betas into the body's "betas" array, which the Claude
// executor lifts into the Anthropic-Beta header.
func addBetas(body []byte, add []string) ([]byte, error) {
	seen := map[string]bool{}
	var all []string
	ojson.GetBytes(body, "betas").ForEach(func(_, v ojson.Result) bool {
		if s := v.String(); s != "" && !seen[s] {
			seen[s] = true
			all = append(all, s)
		}
		return true
	})
	for _, b := range add {
		if !seen[b] {
			seen[b] = true
			all = append(all, b)
		}
	}
	return ojson.SetBytes(body, "betas", all)
}
