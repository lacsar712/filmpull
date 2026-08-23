package fsm

import (
	"github.com/lacsar712/filmpull/internal/model"
)

var lineTransitions = map[model.LineState]map[string]model.LineState{
	model.LineIdle: {
		"thread": model.LineThreading,
		"shutdown": model.LineShutdown,
	},
	model.LineThreading: {
		"preheat": model.LinePreheat,
		"fault":   model.LineFault,
		"abort":   model.LineIdle,
	},
	model.LinePreheat: {
		"stretch": model.LineStretch,
		"fault":   model.LineFault,
		"coast":   model.LineCoast,
	},
	model.LineStretch: {
		"anneal": model.LineAnneal,
		"fault":  model.LineFault,
		"coast":  model.LineCoast,
	},
	model.LineAnneal: {
		"coast":    model.LineCoast,
		"shutdown": model.LineShutdown,
		"fault":    model.LineFault,
	},
	model.LineCoast: {
		"idle":     model.LineIdle,
		"shutdown": model.LineShutdown,
		"fault":    model.LineFault,
	},
	model.LineFault: {
		"reset": model.LineIdle,
	},
	model.LineShutdown: {},
}

func MustLine(from model.LineState, event string) (model.LineState, error) {
	table, ok := lineTransitions[from]
	if !ok {
		return "", model.Wrap("fsm", "line_state", model.ErrTransition)
	}
	next, ok := table[event]
	if !ok {
		return "", model.Wrap("fsm", "line_event", model.ErrTransition)
	}
	return next, nil
}

var nipTransitions = map[model.NipState]map[string]model.NipState{
	model.NipOpen: {
		"close": model.NipClosing,
		"fault": model.NipFault,
	},
	model.NipClosing: {
		"seated": model.NipClosed,
		"fault":  model.NipFault,
		"open":   model.NipOpen,
	},
	model.NipClosed: {
		"open":  model.NipOpen,
		"fault": model.NipFault,
	},
	model.NipFault: {
		"reset": model.NipOpen,
	},
}

func MustNip(from model.NipState, event string) (model.NipState, error) {
	table, ok := nipTransitions[from]
	if !ok {
		return "", model.Wrap("fsm", "nip_state", model.ErrTransition)
	}
	next, ok := table[event]
	if !ok {
		return "", model.Wrap("fsm", "nip_event", model.ErrTransition)
	}
	return next, nil
}