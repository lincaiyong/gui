package all

import (
	"github.com/lincaiyong/gui/com"
	"github.com/lincaiyong/gui/com/bar"
	"github.com/lincaiyong/gui/com/button"
	"github.com/lincaiyong/gui/com/compare"
	"github.com/lincaiyong/gui/com/container"
	"github.com/lincaiyong/gui/com/div"
	"github.com/lincaiyong/gui/com/divider"
	"github.com/lincaiyong/gui/com/editor"
	"github.com/lincaiyong/gui/com/iframe"
	"github.com/lincaiyong/gui/com/img"
	"github.com/lincaiyong/gui/com/input"
	"github.com/lincaiyong/gui/com/root"
	"github.com/lincaiyong/gui/com/scrollbar"
	"github.com/lincaiyong/gui/com/table"
	"github.com/lincaiyong/gui/com/text"
	"github.com/lincaiyong/gui/com/toyeditor"
	"github.com/lincaiyong/gui/com/tree"
)

func HBar() *bar.Component {
	return bar.HBar()
}

func VBar() *bar.Component {
	return bar.VBar()
}

func Button() *button.Component {
	return button.Button()
}

func ToolButton() *button.Component {
	return button.ToolButton()
}

func SourceRootButton() *button.Component {
	return button.SourceRootButton()
}

func SourceDirButton() *button.Component {
	return button.SourceDirButton()
}

func SourceFileButton() *button.Component {
	return button.SourceFileButton()
}

func Compare() *compare.Component {
	return compare.Compare()
}

func VListContainer(children ...com.Component) *container.Component {
	return container.VListContainer(children...)
}

func ListContainer(children ...com.Component) *container.Component {
	return container.ListContainer(children...)
}

func Container(child com.Component) *container.Component {
	return container.Container(child)
}

func Div() *div.Component {
	return div.Div()
}

func HDivider() *divider.Component {
	return divider.HDivider()
}

func VDivider() *divider.Component {
	return divider.VDivider()
}

func Editor() *editor.Component {
	return editor.Editor()
}

func Iframe() *iframe.Component {
	return iframe.Iframe()
}

func Img(src string) *img.Component {
	return img.Img(src)
}

func Svg(src string) *img.Component {
	return img.Svg(src)
}

func Input() *input.Component {
	return input.Input()
}

func Root(children ...com.Component) *root.Component {
	return root.Root(children...)
}

func HScrollbar() *scrollbar.Component {
	return scrollbar.HScrollbar()
}

func VScrollbar() *scrollbar.Component {
	return scrollbar.VScrollbar()
}

func Table() *table.Component {
	return table.Table()
}

func Text(t string) *text.Component {
	return text.Text(t)
}

func ToyEditor() *toyeditor.Component {
	return toyeditor.ToyEditor()
}

func Tree() *tree.Component {
	return tree.Tree()
}
