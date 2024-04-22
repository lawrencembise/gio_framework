package home

import (
	"encoding/json"
	"gioFramework/icons"
	"gioFramework/models"
	page "gioFramework/pages"
	"gioFramework/pages/widgets"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Page holds the state for a page demonstrating the features of
// the NavDrawer component.
type Page struct {
	nonModalDrawer widget.Bool
	widget.List
	*page.Router
	TodoWidgets []*widgets.TodoWidget
}

// New constructs a Page with the provided router.
func New(router *page.Router) *Page {
	p := &Page{
		Router: router,
	}
	if err := p.FetchTodos(); err != nil {
		log.Println("Failed to fetch todos:", err)
		// Handle error more gracefully depending on your application requirements
	}
	return p
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Home",
		Icon: icons.HomeIcon,
	}
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	return material.List(th, &p.List).Layout(gtx, len(p.TodoWidgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
			return p.TodoWidgets[i].Layout(gtx, th)
		})
	})
}

func (p *Page) FetchTodos() error {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var todos []models.Todo
	if err := json.Unmarshal(body, &todos); err != nil {
		return err
	}

	p.TodoWidgets = make([]*widgets.TodoWidget, len(todos))
	for i, todo := range todos {
		mdl := models.Todo{
			Title:     todo.Title,
			Completed: todo.Completed,
		}
		p.TodoWidgets[i] = widgets.NewTodoWidget(mdl)
	}
	return nil
}
