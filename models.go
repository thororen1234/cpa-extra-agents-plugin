package main

import "github.com/router-for-me/CLIProxyAPI/v8/sdk/pluginapi"

const mythosDescription = "Anthropic's Mythos-class frontier model, for the most demanding reasoning and long-horizon agentic work"

const mythos51Tail = " extends Claude Mythos 5 at the same input and output prices, with cache reads at a quarter of the cost, and brings stronger long-running agentic coding, multistep research, and document, spreadsheet, and slide work."

func mythosModels() []pluginapi.ModelInfo {
	mk := func(id, name, desc string, created int64) pluginapi.ModelInfo {
		return pluginapi.ModelInfo{
			ID:                  id,
			Object:              "model",
			Created:             created,
			OwnedBy:             "anthropic",
			Type:                "claude",
			DisplayName:         name,
			Description:         desc,
			ContextLength:       1000000,
			MaxCompletionTokens: 128000,
			Thinking: &pluginapi.ThinkingSupport{
				ZeroAllowed:    false,
				DynamicAllowed: true,
				Levels:         []string{"low", "medium", "high", "xhigh", "max"},
			},
			SupportedInputModalities:  []string{"text", "image"},
			SupportedOutputModalities: []string{"text"},
		}
	}
	return []pluginapi.ModelInfo{
		mk("claude-mythos", "Claude Mythos", mythosDescription, 1781049600),
		mk("claude-mythos-5", "Claude Mythos 5", mythosDescription, 1781049600),
		mk("claude-mythos-5-1", "Claude Mythos 5.1", "Claude Mythos 5.1"+mythos51Tail, 1788220800),
		mk("claude-mythos-5-2", "Claude Mythos 5.2", "Claude Mythos 5.2"+mythos51Tail, 1794250800),
	}
}
