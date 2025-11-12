type EditorCursor struct[
    position: EditorPosition
    valid: bool
]

type EditorAnchor struct[
    position: EditorPosition
    valid: bool
]

type EditorKeyPressed struct[
    shift|ctrl|meta: bool
]

type EditorColumnItem struct[
    left|right|width: number
]

type EditorTokenItem struct[
    key|text|kind|color: string
    start|end: number
]

type EditorSelectionItem struct[
    lineIndex|start|end: number
    color: string
]

type EditorLineItem struct[
    chars: list[string]
    key|lineNo: string
    lineNoWidth|width: number
    colItems: list[EditorColumnItem]
    tokenItems: list[EditorTokenItem]
]

type TokenItem struct[
    line|start|end: number
    kind: string
]

type EditorPage struct[
    lines: list[list[string]],
    items: list[EditorLineItem]
    selectionItems: list[EditorSelectionItem]
    highlightItems: list[EditorSelectionItem]
    brackets: list[TokenItem]
    _freezeUndo: bool
    _undoStack: list[EditorTransaction]
    _redoIndex: number
]

type EditorPosition struct[
    lineIndex|colIndex: number
]

type EditorRange struct[
    start|end: EditorPosition
]

type EditorTransaction struct[
    command|text: string
    range: EditorRange
]

type ValueChangeListener function[EditorTransaction]

type EditorData struct[
    value: string
    fontSize: number
    page: EditorPage
    cursor: EditorCursor
    anchor: EditorAnchor
    keyPressed: EditorKeyPressed
    listeners: list[ValueChangeListener]
    compositionStartCursor: EditorPosition
    lineNoMaxWidth: number
]

type LineNoData struct[
    lineNo: string
    width: number
]

type EditorComponent component[
    Component

    data: EditorData
    private pageScrollTop: number
    private pageScrollLeft: number
    private pageWidth: number
    private pageLineCount: number
    private pageMargin: number

    static function computeLineNo: ContainerItemComputeFunc<LineNoData>
    static function computeLine: ContainerItemComputeFunc<EditorLineItem>
    static function computeSpan: ContainerItemComputeFunc<EditorTokenItem>
    static function computeSelectionItem: ContainerItemComputeFunc<EditorSelectionItem>

    function handlePageContextMenu: MouseEventHandler

    // public
    function setValue: function[string]
    function getValue: function[]
    function onValueChanged: function[]
    function publishChange: function[EditorTransaction]

    // util
    function newPosition: function[number, number]
    function copyPosition: function[EditorPosition]
    function samePosition: function[EditorPosition, EditorPosition]
    function newRange: function[EditorPosition, EditorPosition]

    // anchor
    function createAnchor: function[]
    function clearAnchor: function[]

    // cursor
    function checkWordItems: function[number]
    function moveCursorToPixel: function[number, number] // x,y
    function moveCursorLeftIgnoreRange: function[]
    function moveCursorLeft: function[]
    function moveCursorLeftByLine: function[]
    function moveCursorLeftByWord: function[]
    function moveCursorRightIgnoreRange: function[]
    function moveCursorRight: function[]
    function moveCursorRightByLine: function[]
    function moveCursorRightByWord: function[]
    function moveCursorUp: function[]
    function moveCursorDown: function[]
    function moveCursorToLine: function[number]
    function moveCursorToCol: function[number]
    function moveCursorToPosition: function[EditorPosition]
    function moveCursorToPageStart: function[]
    function moveCursorToPageEnd: function[]
    function hideCursor: function[]
    function refreshCursor: function[bool] // scrollPage

    // kb
    function handleKeyUp: KeyboardEventHandler
    function handleKeyDown: KeyboardEventHandler
    function handleCompositionUpdate: CompositionEventHandler
    function handleCompositionEnd: CompositionEventHandler
    function handleInput: InputEventHandler
    function handleFocus: EventHandler
    function handlePaste: ClipboardEventHandler
    function handleCopy: ClipboardEventHandler
    function handleCut: ClipboardEventHandler

    // mouse
    function handleMouseDown: MouseEventHandler
    function handleDoubleClick: MouseEventHandler
    function handleScrollTop: ScrollEventHandler
    function handleScrollLeft: ScrollEventHandler
    function _showCursorOnMouseEvent: function[]
    function getWordByCursor: function[]

    // cmd
    function doCreatePage: function[string]
    function doInsertText: function[string]
    function doDeleteText: function[]
    function doCopyText: function[]
    function doCutText: function[]
    function doPasteText: function[string]

    // page
    function createPageContent: function[string]
    function pageInsertText: function[string]
    function pageGetText: function[]
    function pageDeleteText: function[]
    function pageTokenize: function[]
    function _deletePageItemAfter: function[number]
    function _insertPageItemAfter: function[number]
    function _updatePageItem: function[number]
    function _updatePage: function[]
    function refreshPage: function[]

    // selection
    function updateSelection: function[]
    function getRange: function[]
    function hasVisibleRange: function[]
    function refreshSelection: function[]

    // highlight
    function refreshHighlight: function[]
    function highlightWordsByCursor: function[]

    // undo
    function _clearUndo: function[]
    function _pushUndo: function[string, string, EditorRange] // command, text, range
    function freezeUndo: function[]
    function executeUndo: function[]
    function executeRedo: function[]
]

component editor[EditorComponent] (
    pageScrollTop=contentEle.scrollTop
    pageScrollLeft=contentEle.scrollLeft
    pageWidth=contentEle.w
    pageMargin=1
) {
    activeLineEle:div(v=0, h=20, backgroundColor=webapp.theme.editorActiveLineColor)

    leftEle:container(list=true, virtual=true, scrollable=false, scrollTop=this.pageScrollTop,
        w=30 + (this.pageLineCount + '').length * 7.2, borderRight=1, borderColor=webapp.theme.grayBorderColor) {
        container_item<LineNoData>(compute=computeLineNo) {
            text(h=18, y=1, x=parent.w - parent.data.width - 20, text=parent.data.lineNo, color=webapp.theme.editorLineNoColor, fontFamily=mono)
        }
    }

    pageEle:div(x=prev.x2, w=parent.w - prev.w, onMouseDown=handleMouseDown, onDoubleClick=handleDoubleClick, cursor=text, onContextMenu=handlePageContextMenu) {
        highlightEle:container(x=this.pageMargin, w=this.pageWidth, scrollLeft=this.pageScrollLeft, scrollTop=this.pageScrollTop, scrollable=false, list=true, virtual=true, zIndex=background) {
            container_item<EditorSelectionItem>(compute=computeSelectionItem, backgroundColor=.data.color)
        }

        selectionEle:container(x=this.pageMargin, w=this.pageWidth, scrollLeft=this.pageScrollLeft, scrollTop=this.pageScrollTop, scrollable=false, list=true, virtual=true) {
            container_item<EditorSelectionItem>(compute=computeSelectionItem, backgroundColor=webapp.theme.editorSelectionColor)
        }

        cursorEle:div(v=1, x=100, y=100, w=2, h=20, animation=webapp.animation.blink, backgroundColor='black', cursor=text) {
            fakeInputEle:input(w=1000, opacity=0,
                onInput=handleInput,
                onCompositionUpdate=handleCompositionUpdate,
                onCompositionEnd=handleCompositionEnd,
                onKeyUp=handleKeyUp,
                onKeyDown=handleKeyDown,
                onFocus=handleFocus,
                onPaste=handlePaste,
                onCopy=handleCopy,
                onCut=handleCut,
            )
        }

        contentEle:container(x=this.pageMargin, list=true, virtual=true, onScrollTop=handleScrollTop, onScrollLeft=handleScrollLeft) {
            container_item<EditorLineItem>(compute=computeLine) {
                container(items=parent.data.tokenItems, list=true, y=1, h=18, scrollable=false, align=fill, tag=pre, reuseItem=false) {
                    container_item<EditorTokenItem>(
                        compute=computeSpan,
                        tag=span, position=static,
                        fontFamily=mono, fontSize=12, lineHeight=18,
                        innerText=.data.text,
                        color=.data.color,
                    )
                }
            }
        }
    }
}
