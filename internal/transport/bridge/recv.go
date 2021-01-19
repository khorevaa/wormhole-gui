package bridge

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/Jacalz/wormhole-gui/internal/transport"
)

// RecvItem is the item that is being received
type RecvItem struct {
	URI    fyne.URI
	Status string
	Name   string
}

// RecvList is a list of progress bars that track send progress.
type RecvList struct {
	widget.List

	client *transport.Client

	Items []*RecvItem
}

// Length returns the length of the data.
func (p *RecvList) Length() int {
	return len(p.Items)
}

// CreateItem creates a new item in the list.
func (p *RecvList) CreateItem() fyne.CanvasObject {
	return container.New(&listLayout{}, widget.NewFileIcon(nil), widget.NewLabel("Waiting for filename..."), newRecvProgress())
}

// UpdateItem updates the data in the list.
func (p *RecvList) UpdateItem(i int, item fyne.CanvasObject) {
	item.(*fyne.Container).Objects[0].(*widget.FileIcon).SetURI(p.Items[i].URI)
	item.(*fyne.Container).Objects[1].(*widget.Label).SetText(p.Items[i].Name)
	item.(*fyne.Container).Objects[2].(*fyne.Container).Objects[0].(*recvProgress).setStatus(p.Items[i].Status)
}

// RemoveItem removes the item at the specified index.
func (p *RecvList) RemoveItem(i int) {
	copy(p.Items[i:], p.Items[i+1:])
	p.Items[p.Length()-1] = nil // Make sure that GC run on removed element
	p.Items = p.Items[:p.Length()-1]
	p.Refresh()
}

// OnSelected handles removing items and stopping send (in the future)
func (p *RecvList) OnSelected(i int) {
	dialog.ShowConfirm("Remove from list", "Do you wish to remove the item from the list?", func(remove bool) {
		if remove {
			p.RemoveItem(i)
			p.Refresh()
		}
	}, fyne.CurrentApp().Driver().AllWindows()[0])

	p.Unselect(i)
}

// NewReceive adds data about a new send to the list and then returns the channel to update the code.
func (p *RecvList) NewReceive(code string) {
	p.Items = append(p.Items, &RecvItem{Name: "Waiting for filename..."})
	p.Refresh()

	path := make(chan string)
	index := p.Length() - 1

	go func() {
		p.Items[index].Name = <-path
		if p.Items[index].Name != "Text Snippet" {
			p.Items[index].URI = storage.NewFileURI(p.Items[index].Name)
			p.Items[index].Name = p.Items[index].URI.Name()
		}

		close(path)
		p.Refresh()
	}()

	go func(code string) {
		if err := p.client.NewReceive(code, path); err != nil {
			p.Items[index].Status = "Failed"
			p.client.ShowNotification("Receive failed", "An error occurred when receiving the data.")
			dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
		} else {
			p.Items[index].Status = "Completed"
			p.client.ShowNotification("Receive completed", "The data was received successfully.")
		}

		p.Refresh()
	}(code)
}

// NewRecvList greates a list of progress bars.
func NewRecvList(client *transport.Client) *RecvList {
	p := &RecvList{client: client}
	p.List.Length = p.Length
	p.List.CreateItem = p.CreateItem
	p.List.UpdateItem = p.UpdateItem
	p.List.OnSelected = p.OnSelected
	p.ExtendBaseWidget(p)

	return p
}
