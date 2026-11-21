package main

import "github.com/router-for-me/CLIProxyAPI/v8/sdk/pluginapi"

const mythosDescription = "Claude Mythos 5 shares Claude Fable 5's specifications and is available only to organizations verified through Anthropic's verification programs"

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
		mk("claude-mythos-5", "Claude Mythos 5", mythosDescription, 1781049600),
		mk("claude-mythos-5-1", "Claude Mythos 5.1", "Claude Mythos 5.1"+mythos51Tail, 1788220800),
		mk("claude-mythos-5-2", "Claude Mythos 5.2", "Claude Mythos 5.2"+mythos51Tail, 1794250800),
	}
}

func antigravityModels() []pluginapi.ModelInfo {
	return []pluginapi.ModelInfo{
		{
			ID:                  "gemini-4-carbon-high",
			Object:              "model",
			OwnedBy:             "antigravity",
			Type:                "antigravity",
			DisplayName:         "Gemini 4 Carbon",
			Name:                "gemini-4-carbon-high",
			Description:         "Gemini 4 Carbon (High)",
			ContextLength:       1048576,
			MaxCompletionTokens: 1048576,
			Thinking: &pluginapi.ThinkingSupport{
				Min:            1,
				Max:            65535,
				DynamicAllowed: true,
				Levels:         []string{"low", "medium", "high"},
			},
			SupportedInputModalities:  []string{"text", "image", "audio", "video"},
			SupportedOutputModalities: []string{"text"},
		},
		{
			ID:                  "gemini-4-argon-high",
			Object:              "model",
			OwnedBy:             "antigravity",
			Type:                "antigravity",
			DisplayName:         "Gemini 4 Argon",
			Name:                "gemini-4-argon-high",
			Description:         "Gemini 4 Argon (High)",
			ContextLength:       1048576,
			MaxCompletionTokens: 1048576,
			Thinking: &pluginapi.ThinkingSupport{
				Min:            1,
				Max:            65535,
				DynamicAllowed: true,
				Levels:         []string{"low", "medium", "high"},
			},
			SupportedInputModalities:  []string{"text", "image", "audio", "video"},
			SupportedOutputModalities: []string{"text"},
		},
		{
			ID:                  "gemini-4-neon-xhigh",
			Object:              "model",
			OwnedBy:             "antigravity",
			Type:                "antigravity",
			DisplayName:         "Gemini 4 Neon",
			Name:                "gemini-4-neon-xhigh",
			Description:         "Gemini 4 Neon (XHigh)",
			ContextLength:       1048576,
			MaxCompletionTokens: 1048576,
			Thinking: &pluginapi.ThinkingSupport{
				Min:            1,
				Max:            65535,
				DynamicAllowed: true,
				Levels:         []string{"low", "medium", "high", "xhigh"},
			},
			SupportedInputModalities:  []string{"text", "image", "audio", "video"},
			SupportedOutputModalities: []string{"text"},
		},
		{
			ID:                  "gemini-4-hydrogen-xhigh",
			Object:              "model",
			OwnedBy:             "antigravity",
			Type:                "antigravity",
			DisplayName:         "Gemini 4 Hydrogen",
			Name:                "gemini-4-hydrogen-xhigh",
			Description:         "Gemini 4 Hydrogen (XHigh)",
			ContextLength:       1048576,
			MaxCompletionTokens: 1048576,
			Thinking: &pluginapi.ThinkingSupport{
				Min:            1,
				Max:            65535,
				DynamicAllowed: true,
				Levels:         []string{"low", "medium", "high", "xhigh"},
			},
			SupportedInputModalities:  []string{"text", "image", "audio", "video"},
			SupportedOutputModalities: []string{"text"},
		},
		{
			ID:                  "gemini-3.1-pro-preview",
			Object:              "model",
			OwnedBy:             "antigravity",
			Type:                "antigravity",
			DisplayName:         "Gemini 3.1 Pro",
			Name:                "gemini-3.1-pro-preview",
			Description:         "Gemini 3.1 Pro",
			ContextLength:       1048576,
			MaxCompletionTokens: 65536,
			Thinking: &pluginapi.ThinkingSupport{
				Min:            128,
				Max:            32768,
				DynamicAllowed: true,
				Levels:         []string{"low", "medium", "high"},
			},
			SupportedInputModalities:  []string{"text", "image", "audio", "video"},
			SupportedOutputModalities: []string{"text"},
		},
		{
			ID:                  "gemini-3.8-live",
			Object:              "model",
			OwnedBy:             "antigravity",
			Type:                "antigravity",
			DisplayName:         "Gemini 3.8 Live",
			Name:                "gemini-3.8-live",
			Description:         "Gemini 3.8 Live",
			ContextLength:       1048576,
			MaxCompletionTokens: 65536,
			Thinking: &pluginapi.ThinkingSupport{
				Min:            128,
				Max:            65535,
				DynamicAllowed: true,
				Levels:         []string{"low", "medium", "high"},
			},
			SupportedInputModalities:  []string{"text", "image", "audio", "video"},
			SupportedOutputModalities: []string{"text"},
		},
		{
			ID:                  "gemini-3.6-flash",
			Object:              "model",
			OwnedBy:             "antigravity",
			Type:                "antigravity",
			DisplayName:         "Gemini 3.6 Flash",
			Name:                "gemini-3.6-flash",
			Description:         "Gemini 3.6 Flash",
			ContextLength:       1048576,
			MaxCompletionTokens: 65536,
			Thinking: &pluginapi.ThinkingSupport{
				Min:            128,
				Max:            32768,
				DynamicAllowed: true,
				Levels:         []string{"minimal", "low", "medium", "high"},
			},
			SupportedInputModalities:  []string{"text", "image", "audio", "video"},
			SupportedOutputModalities: []string{"text"},
		},
		{
			ID:                  "gemini-3.1-flash-lite",
			Object:              "model",
			OwnedBy:             "antigravity",
			Type:                "antigravity",
			DisplayName:         "Gemini 3.1 Flash-Lite",
			Name:                "gemini-3.1-flash-lite",
			Description:         "Gemini 3.1 Flash-Lite",
			ContextLength:       1048576,
			MaxCompletionTokens: 65536,
			Thinking: &pluginapi.ThinkingSupport{
				Min:            128,
				Max:            32768,
				DynamicAllowed: true,
				Levels:         []string{"minimal", "low", "medium", "high"},
			},
			SupportedInputModalities:  []string{"text", "image", "audio", "video"},
			SupportedOutputModalities: []string{"text"},
		},
	}
}
