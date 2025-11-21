type ExplorerPaneComponent component[
    Component
    menuOpened: bool

    function handleClickOptions: MouseEventHandler
    function handleClickItem: MouseEventHandler
    function handleHover: HoverEventHandler
]

component explorer_pane[ExplorerPaneComponent](
    onHover=handleHover
    hovered=.hoveredByMouse || .menuOpened
) {
    div(h=32, borderColor=webapp.theme.grayBorderColor, borderBottom=1) {
        text(text='Project', x=7, y=.x, h=parent.h - .y * 2)
        menuContainer:div(x=prev.x2, w=parent.w-.x, v=this.hovered ? 1 : 0) {
            button(x=next.x - .w - 3, y=next.y, icon=select_opened_file, tooltip='Select Opened File', tooltipPos=bottom)
            button(x=next.x - .w - 3, y=next.y, icon=expand_all, tooltip='Expand All', tooltipPos=bottom)
            button(x=next.x - .w - 3, y=next.y, icon=collapse_all, tooltip='Collapse All', tooltipPos=bottom)
            button(x=next.x - .w - 3, y=next.y, icon=options, tooltip='Options', tooltipPos=bottom, onClick=handleClickOptions)
            button(x=parent.w - .w - 8, y=4, icon=hide, tooltip='Hide', tooltipPos=bottom)
        }
    }
    div(y=prev.y2, h=parent.h-.y) {
        treeEle:tree(onClickItem=handleClickItem)
    }
}