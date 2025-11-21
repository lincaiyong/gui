type TopComponent component[
    Component
    static function handleClickProjectBtn: MouseEventHandler
]

component top[TopComponent](tag=div) {
    titleBar:div(h=34, backgroundColor='black')
    div(y=prev.y2, h=parent.h - prev.h - next.h) {
        leftSideBar:div(w=32, borderColor=webapp.theme.grayBorderColor, borderRight=1, backgroundColor=webapp.theme.grayPaneColor) {
            projectBtn:button(x=3, y=4, selected=true, icon=folder, tooltip='Project', onClick=handleClickProjectBtn)
            commitBtn:button(x=prev.x, y=prev.y2 + 8, icon=commit, tooltip='Commit')
            divider(x=prev.x, y=prev.y2 + 8, w=prev.w)
            button(x=prev.x, y=prev.y2 + 8, icon=structure, tooltip='Structure')
        }

        div(x=prev.x2, w=parent.w - prev.w - next.w) {
            mainPane:div(h=next.y) {
                explorerPane:explorer_pane(w=next.x, v=1, backgroundColor=webapp.theme.grayPaneColor)
                bar(x=prev.v ? 200 : - .w, w=6, leftRight=[prev, next], backgroundColor='red', opacity=0.1)
                editorEle:editor(x=prev.x2, w=next.x - prev.x2)
                bar(x=parent.w - 400, w=6, leftRight=[prev, next], backgroundColor='red', opacity=0.1)
                noteEle:note(x=prev.x2, w=parent.w - .x)
            }

            bar(y=.v ? parent.h - 200 : parent.h, h=6, v=next.v, topBottom=[prev, next], cursor=rowResize, backgroundColor='red', opacity=0.1)

            bottomPane:div(y=prev.y2, h=parent.h - .y, v=1)
        }

        rightSideBar:div(x=prev.x2, w=32, borderColor=webapp.theme.grayBorderColor, borderLeft=1, backgroundColor=webapp.theme.grayPaneColor) {}
    }
    statusBar:div(y=prev.y2, h=24, borderColor=webapp.theme.grayBorderColor, borderTop=1, backgroundColor=webapp.theme.grayPaneColor)
}