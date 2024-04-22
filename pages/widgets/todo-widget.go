package widgets

import (
	"gioFramework/models"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type TodoWidget struct {
	todo models.Todo
}

func NewTodoWidget(todo models.Todo) *TodoWidget {
	return &TodoWidget{todo: todo}
}

func (tw *TodoWidget) Layout(gtx C, th *material.Theme) D {
	return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(1, func(gtx C) D {
				title := material.Body1(th, tw.todo.Title)
				title.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				return title.Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				statusText := "Pending"
				if tw.todo.Completed {
					statusText = "Completed"
				}
				status := material.Caption(th, statusText)
				status.Color = color.NRGBA{R: 0, G: 128, B: 0, A: 255}
				return status.Layout(gtx)
			}),
		)
	})
}
