class Component {
    constructor(parent, model) {
        const {properties, children=[]} = model;
        this._properties = {};
        this._parent = parent;
        this._model = model;
        this._id = parent ? `${parent.id}.${model.name}` : model.name;
        this._ref = document.createElement(model.tag);
        this._ref.style.position = model.position;
        this._ref.style.overflow = model.overflow || 'hidden';
        this._ref.style.boxSizing = 'border-box';
        this._children = children.map(childData => new childData.Component(this, childData));

        this._sideEffects = {};

        // initialize properties
        const props = Object.assign(this._defaultProperties, properties);
        for (const k in props) {
            const v = props[k];
            g.util.assert(v instanceof Array && v.length === 2 && v[0] instanceof Function && v[1] instanceof Array, `invalid argument ${v}`);
            const [computeFunc, sources] = v;
            const sourceResolver = source => this._(source);
            this._properties[k] = new Property(this, k, sources, sourceResolver, computeFunc);
            this._properties[k].onUpdated(v => {
                this._defaultOnUpdated(k, v);
                this.onUpdated(k, v);
            });
        }
    }

    get parent() {
        return this._parent;
    }

    get children() {
        return this._children;
    }

    get model() {
        return this._model;
    }

    get id() {
        return this._id;
    }

    get tag() {
        return this.model.tag;
    }

    get ref() {
        return this._ref;
    }

    get root() {
        return [...Array(this.model.depth).keys()].reduce(prev => prev?.parent, this);
    }


    onCreated() {
    }

    onUpdated(k, v) {
    }

    _defaultOnCreated() {
        if (this.onCreatedFn instanceof Function) {
            this.onCreatedFn();
        }
    }

    _defaultOnUpdated(k, v) {
        if (k === 'hovered') {
            this.onHover?.(this, v);
        }
        if (this.onUpdatedFn instanceof Function) {
            this.onUpdatedFn(k, v);
        }
    }

    _createAll(parent) {
        if (parent instanceof Element) {
            parent.appendChild(this.ref);
        } else if (parent instanceof Component) {
            this._parent = parent;
            parent.children.push(this);
            parent.ref.appendChild(this.ref);
        } else {
            g.log.error("invalid argument")
        }
        Object.values(this._properties).forEach(p => p.subscribe());
        this.children.forEach(child => child._createAll(this.ref));
    }

    _initAll() {
        Object.values(this._properties).forEach(p => p.update());
        this.children.forEach(child => child._initAll());
        this._defaultOnCreated();
        this.onCreated();
    }

    _create(parent) {
        this._createAll(parent);
        this._initAll();
    }

    _unInitAll() {
        this.children.forEach(child => child._unInitAll());
        Object.values(this._sideEffects).forEach(fun => fun?.());
        Object.values(this._properties).forEach(p => p.unsubscribe());
    }

    _destroyAll() {
        this.children.forEach(child => child._destroyAll());
        this.parent?.children.splice(this.parent?.children.indexOf(this), 1);
        this.ref.parentElement.removeChild(this.ref);
    }

    _destroy() {
        this._unInitAll();
        this._destroyAll();
    }

    _checkLoop() {
        const properties = this._collectProperties();
        this._topologicalSort(properties);
    }

    _collectProperties() {
        let ret = Object.values(this._properties);
        this.children.forEach(child => ret = ret.concat(child._collectProperties()));
        return ret;
    }

    _topologicalSort(properties) {
        const visited = {};
        let total = properties.length;
        let count = 0;
        for (; ;) {
            for (const prop of properties) {
                if (prop.id in visited) {
                    continue;
                }
                let ok = true;
                for (const source of prop._resolvedSources) {
                    if (!(source.id in visited)) {
                        ok = false;
                        break;
                    }
                }
                if (ok) {
                    visited[prop.id] = true;
                }
            }

            const newCount = Object.keys(visited).length;
            if (total === newCount) {
                break;
            }
            if (count === newCount) {
                const tmp = properties.filter(prop => !(prop.id in visited)).map(prop => `${prop.id}: ${prop._sources.join(', ')}`);
                g.log.error("loop detected", '\n\t' + tmp.join('\n\t'));
                return;
            }
            count = newCount;
        }
    }

    _addSideEffect(on, fun) {
        this._sideEffects[on]?.();
        this._sideEffects[on] = null;
        if (fun instanceof Function) {
            this._sideEffects[on] = fun;
        }
    }

    _(source) {
        g.util.assert(typeof (source) === 'string' && this instanceof Component);
        const ret = this._resolve(source);
        g.util.assert(ret);
        return ret;
    }

    _resolve(source) {
        if (!source.includes('.')) {
            return this._resolveEle(source);
        }
        const [e, p] = source.split('.', 2);
        const target = this._resolveEle(e);
        return target?._properties[p];
    }

    _resolveEle(name) {
        if (name === '') {
            return this;
        } else if (name === 'root') {
            return g.root;
        } else if (name === 'this') {
            return this.root;
        } else if (name === 'parent') {
            return this.parent;
        } else if (name === 'child') {
            return this.children[0];
        } else if (name === 'prev' || name === 'next') {
            return this.parent?.children[this.parent?.children.indexOf(this) + (name === 'prev' ? -1 : 1)];
        } else {
            const m = name.match(/^child([0-9])$/);
            if (m) {
                return this.children[parseInt(m[1])];
            }
            return this[name];
        }
    }

    get _defaultProperties() {
        return {
            onUpdatedFn: [() => null, []],
            onCreatedFn: [() => null, []],
            background: [() => '', []],
            backgroundColor: [() => '', []],
            borderBottom: [() => 0, []],
            borderColor: [() => 'black', []],
            borderLeft: [() => 0, []],
            borderRadius: [() => 0, []],
            borderRight: [() => 0, []],
            borderStyle: [() => 'solid', []],
            borderTop: [() => 0, []],
            boxShadow: [() => '', []],
            caretColor: [() => '', []],
            ch: [() => 0, []],
            color: [() => '', []],
            cursor: [() => 'inherit', []],
            cw: [() => 0, []],
            fontFamily: [() => 'SF Pro Display', []],
            fontSize: [() => 0, []],
            fontVariantLigatures: [() => 'none', []],
            h: [() => 0, []],
            hovered: [() => false, []],
            hoveredByMouse: [() => false, []],
            innerText: [() => '', []],
            lineHeight: [() => 0, []],
            onActive: [() => undefined, []],
            onClick: [() => undefined, []],
            onClickOutside: [() => undefined, []],
            onCompositionEnd: [() => undefined, []],
            onCompositionStart: [() => undefined, []],
            onCompositionUpdate: [() => undefined, []],
            onCopy: [() => undefined, []],
            onCut: [() => undefined, []],
            onDoubleClick: [() => undefined, []],
            onFocus: [() => undefined, []],
            onHover: [() => undefined, []],
            onInput: [() => undefined, []],
            onKeyDown: [() => undefined, []],
            onKeyUp: [() => undefined, []],
            onMouseDown: [() => undefined, []],
            onMouseMove: [() => undefined, []],
            onMouseUp: [() => undefined, []],
            onPaste: [() => undefined, []],
            onScrollLeft: [() => undefined, []],
            onScrollTop: [() => undefined, []],
            onWheel: [() => undefined, []],
            opacity: [() => 1, []],
            outline: [() => 'none', []],
            position: [() => 'absolute', []],
            scrollLeft: [() => 0, []],
            scrollTop: [() => 0, []],
            userSelect: [() => 'none', []],
            v: [() => 0, []],
            w: [() => 0, []],
            x: [() => 0, []],
            x2: [() => 0, []],
            y: [() => 0, []],
            y2: [() => 0, []],
            zIndex: [() => 0, []],
        };
    };

    // builtin properties
    get background() {
        return this._properties.background.value;
    }

    get backgroundColor() {
        return this._properties.backgroundColor.value;
    }

    get borderBottom() {
        return this._properties.borderBottom.value;
    }

    get borderColor() {
        return this._properties.borderColor.value;
    }

    get borderLeft() {
        return this._properties.borderLeft.value;
    }

    get borderRadius() {
        return this._properties.borderRadius.value;
    }

    get borderRight() {
        return this._properties.borderRight.value;
    }

    get borderStyle() {
        return this._properties.borderStyle.value;
    }

    get borderTop() {
        return this._properties.borderTop.value;
    }

    get boxShadow() {
        return this._properties.boxShadow.value;
    }

    get caretColor() {
        return this._properties.caretColor.value;
    }

    get ch() {
        return this._properties.ch.value;
    }

    get color() {
        return this._properties.color.value;
    }

    get cursor() {
        return this._properties.cursor.value;
    }

    get cw() {
        return this._properties.cw.value;
    }

    get fontFamily() {
        return this._properties.fontFamily.value;
    }

    get fontSize() {
        return this._properties.fontSize.value;
    }

    get fontWeight() {
        return this._properties.fontWeight.value;
    }

    get fontVariantLigatures() {
        return this._properties.fontVariantLigatures.value;
    }

    get h() {
        return this._properties.h.value;
    }

    get hovered() {
        return this._properties.hovered.value;
    }

    get hoveredByMouse() {
        return this._properties.hoveredByMouse.value;
    }

    get innerText() {
        return this._properties.innerText.value;
    }

    get lineHeight() {
        return this._properties.lineHeight.value;
    }

    get onActive() {
        return this._properties.onActive.value;
    }

    get onClick() {
        return this._properties.onClick.value;
    }

    get onClickOutside() {
        return this._properties.onClickOutside.value;
    }

    get onCompositionEnd() {
        return this._properties.onCompositionEnd.value;
    }

    get onCompositionStart() {
        return this._properties.onCompositionStart.value;
    }

    get onCompositionUpdate() {
        return this._properties.onCompositionUpdate.value;
    }

    get onCopy() {
        return this._properties.onCopy.value;
    }

    get onCut() {
        return this._properties.onCut.value;
    }

    get onDoubleClick() {
        return this._properties.onDoubleClick.value;
    }

    get onFocus() {
        return this._properties.onFocus.value;
    }

    get onHover() {
        return this._properties.onHover.value;
    }

    get onInput() {
        return this._properties.onInput.value;
    }

    get onKeyDown() {
        return this._properties.onKeyDown.value;
    }

    get onKeyUp() {
        return this._properties.onKeyUp.value;
    }

    get onMouseDown() {
        return this._properties.onMouseDown.value;
    }

    get onMouseMove() {
        return this._properties.onMouseMove.value;
    }

    get onMouseUp() {
        return this._properties.onMouseUp.value;
    }

    get onPaste() {
        return this._properties.onPaste.value;
    }

    get onScrollLeft() {
        return this._properties.onScrollLeft.value;
    }

    get onScrollTop() {
        return this._properties.onScrollTop.value;
    }

    get onWheel() {
        return this._properties.onWheel.value;
    }

    get opacity() {
        return this._properties.opacity.value;
    }

    get outline() {
        return this._properties.outline.value;
    }

    get position() {
        return this._properties.position.value;
    }

    get scrollLeft() {
        return this._properties.scrollLeft.value;
    }

    get scrollTop() {
        return this._properties.scrollTop.value;
    }

    get userSelect() {
        return this._properties.userSelect.value;
    }

    get v() {
        return this._properties.v.value;
    }

    get w() {
        return this._properties.w.value;
    }

    get x() {
        return this._properties.x.value;
    }

    get x2() {
        return this._properties.x2.value;
    }

    get y() {
        return this._properties.y.value;
    }

    get y2() {
        return this._properties.y2.value;
    }

    get zIndex() {
        return this._properties.zIndex.value;
    }

    get onCreatedFn() {
        return this._properties.onCreatedFn.value;
    }

    get onUpdatedFn() {
        return this._properties.onUpdatedFn.value;
    }

    set background(v) {
        if (this.background !== v) {
            this._properties.background.value = v;
            this.ref.style.background = v;
        }
    }

    set backgroundColor(v) {
        if (this.backgroundColor !== v) {
            this._properties.backgroundColor.value = v;
            this.ref.style.backgroundColor = v;
        }
    }

    set borderBottom(v) {
        if (this.borderBottom !== v) {
            this._properties.borderBottom.value = v;
            this.ref.style.borderBottomWidth = v + 'px';
        }
    }

    set borderColor(v) {
        if (this.borderColor !== v) {
            this._properties.borderColor.value = v;
            this.ref.style.borderColor = v;
        }
    }

    set borderLeft(v) {
        if (this.borderLeft !== v) {
            this._properties.borderLeft.value = v;
            this.ref.style.borderLeftWidth = v + 'px';
        }
    }

    set borderRadius(v) {
        if (this.borderRadius !== v) {
            this._properties.borderRadius.value = v;
            this.ref.style.borderRadius = v + 'px';
        }
    }

    set borderRight(v) {
        if (this.borderRight !== v) {
            this._properties.borderRight.value = v;
            this.ref.style.borderRightWidth = v + 'px';
        }
    }

    set borderStyle(v) {
        if (this.borderStyle !== v) {
            this._properties.borderStyle.value = v;
            this.ref.style.borderStyle = v;
        }
    }

    set borderTop(v) {
        if (this.borderTop !== v) {
            this._properties.borderTop.value = v;
            this.ref.style.borderTopWidth = v + 'px';
        }
    }

    set boxShadow(v) {
        if (this.boxShadow !== v) {
            this._properties.boxShadow.value = v;
            this.ref.style.boxShadow = v;
        }
    }

    set caretColor(v) {
        if (this.caretColor !== v) {
            this._properties.caretColor.value = v;
            this.ref.style.caretColor = v;
        }
    }

    set ch(v) {
        if (this.ch !== v) {
            this._properties.ch.value = v;
        }
    }

    set color(v) {
        if (this.color !== v) {
            this._properties.color.value = v;
            this.ref.style.color = v;
        }
    }

    set cursor(v) {
        if (this.cursor !== v) {
            this._properties.cursor.value = v;
            this.ref.style.cursor = v;
        }
    }

    set cw(v) {
        if (this.cw !== v) {
            this._properties.cw.value = v;
        }
    }

    set fontFamily(v) {
        if (this.fontFamily !== v) {
            this._properties.fontFamily.value = v;
            this.ref.style.fontFamily = v;
        }
    }

    set fontSize(v) {
        if (this.fontSize !== v) {
            this._properties.fontSize.value = v;
            this.ref.style.fontSize = v + 'px';
        }
    }

    set fontWeight(v) {
        if (this.fontWeight !== v) {
            this._properties.fontWeight.value = v;
            this.ref.style.fontWeight = v;
        }
    }

    set fontVariantLigatures(v) {
        if (this.fontVariantLigatures !== v) {
            this._properties.fontVariantLigatures.value = v;
            this.ref.style.fontVariantLigatures = v;
        }
    }

    set h(v) {
        if (this.h !== v) {
            this._properties.h.value = v;
            this.ref.style.height = v + 'px';
        }
    }

    set hovered(v) {
        if (this.hovered !== v) {
            this._properties.hovered.value = v;
        }
    }

    set hoveredByMouse(v) {
        if (this.hoveredByMouse !== v) {
            this._properties.hoveredByMouse.value = v;
        }
    }

    set innerText(v) {
        if (typeof (v) === 'string' && this.tag === 'span') {
            this._properties.innerText.value = v;
            this.ref.innerText = v;
        }
    }

    set lineHeight(v) {
        if (this.lineHeight !== v) {
            this._properties.lineHeight.value = v;
            this.ref.style.lineHeight = v + 'px';
        }
    }

    set onActive(v) {
        if (v instanceof Function) {
            this._addSideEffect('onActive', g.event.addListener(this.ref, 'mousedown', ev => {
                const fun = v(this, ev);
                g.event.onceListener(this.ref, 'mouseup', ev => fun?.(this, ev));
            }));
        }
    }

    set onClick(v) {
        if (v instanceof Function) {
            this._addSideEffect('onClick', g.event.addListener(this.ref, 'click', ev => {
                v(this, ev);
            }));
        }
    }

    set onClickOutside(v) {
        if (v instanceof Function) {
            this._addSideEffect('onClickOutside', g.event.addListener(document, 'click', ev => {
                const rect = this.ref.getBoundingClientRect();
                if (rect.x > ev.clientX || rect.y > ev.clientY || (rect.x + rect.width) < ev.clientX || (rect.y + rect.height) < ev.clientY) {
                    const isOutsideEvent = v(this, ev); // ev !== clickEv
                    if (isOutsideEvent) {
                        this._sideEffects.onClickOutside?.();
                    }
                }
            }));
        }
    }

    set onCompositionEnd(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCompositionEnd', g.event.addListener(this.ref, 'compositionend', ev => {
                v(this, ev);
            }));
        }
    }

    set onCompositionStart(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCompositionStart', g.event.addListener(this.ref, 'compositionstart', ev => {
                v(this, ev);
            }));
        }
    }

    set onCompositionUpdate(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCompositionUpdate', g.event.addListener(this.ref, 'compositionupdate', ev => {
                v(this, ev);
            }));
        }
    }

    set onCopy(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCopy', g.event.addListener(this.ref, 'copy', ev => {
                v(this, ev);
            }));
        }
    }

    set onCut(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCut', g.event.addListener(this.ref, 'cut', ev => {
                v(this, ev);
            }));
        }
    }

    set onDoubleClick(v) {
        if (v instanceof Function) {
            this._addSideEffect('onDoubleClick', g.event.addListener(this.ref, 'dblclick', ev => {
                v(this, ev);
            }));
        }
    }

    set onFocus(v) {
        if (v instanceof Function) {
            this._addSideEffect('onFocus', g.event.addListener(this.ref, 'focus', ev => {
                const fun = v(this, ev);
                g.event.onceListener(this.ref, 'blur', ev => fun?.(this, ev));
            }));
        }
    }

    set onHover(v) {
        if (v instanceof Function) {
            this._properties.onHover.value = v;
            this._addSideEffect('mouseenter', g.event.addListener(this.ref, 'mouseenter', () => {
                g.event.onceListener(this.ref, 'mouseleave', () => this.hoveredByMouse = false);
                this.hoveredByMouse = true;
            }));
        }
    }

    set onInput(v) {
        if (v instanceof Function) {
            this._addSideEffect('onInput', g.event.addListener(this.ref, 'input', ev => {
                v(this, ev);
            }));
        }
    }

    set onKeyDown(v) {
        if (v instanceof Function) {
            this._addSideEffect('onKeyDown', g.event.addListener(this.ref, 'keydown', ev => {
                v(this, ev);
            }));
        }
    }

    set onKeyUp(v) {
        if (v instanceof Function) {
            this._addSideEffect('onKeyUp', g.event.addListener(this.ref, 'keyup', ev => {
                v(this, ev);
            }));
        }
    }

    set onMouseDown(v) {
        if (v instanceof Function) {
            this._addSideEffect('onMouseDown', g.event.addListener(this.ref, 'mousedown', ev => {
                v(this, ev);
            }));
        }
    }

    set onMouseMove(v) {
        if (v instanceof Function) {
            this._addSideEffect('onMouseMove', g.event.addListener(this.ref, 'mousemove', ev => {
                v(this, ev);
            }));
        }
    }

    set onMouseUp(v) {
        if (v instanceof Function) {
            this._addSideEffect('onMouseUp', g.event.addListener(this.ref, 'mouseup', ev => {
                v(this, ev);
            }));
        }
    }

    set onPaste(v) {
        if (v instanceof Function) {
            this._addSideEffect('onPaste', g.event.addListener(this.ref, 'paste', ev => {
                v(this, ev);
            }));
        }
    }

    set onScrollLeft(v) {
        if (this.onScrollLeft !== v) {
            this._properties.onScrollLeft.value = v;
        }
    }

    set onScrollTop(v) {
        if (this.onScrollTop !== v) {
            this._properties.onScrollTop.value = v;
        }
    }

    set onWheel(v) {
        if (v instanceof Function) {
            this._addSideEffect('onWheel', g.event.addListener(this.ref, 'wheel', ev => {
                v(this, ev);
            }));
        }
    }

    set opacity(v) {
        if (this.opacity !== v) {
            this._properties.opacity.value = v;
            this.ref.style.opacity = v;
        }
    }

    set outline(v) {
        if (this.outline !== v) {
            this._properties.outline.value = v;
            this.ref.style.outline = v;
        }
    }

    set position(v) {
        if (this.position !== v) {
            this._properties.position.value = v;
            this.ref.style.position = v;
        }
    }

    set scrollLeft(v) {
        if (this.scrollLeft !== v) {
            this._properties.scrollLeft.value = v;
            this.onScrollLeft?.(this, v);
        }
    }

    set scrollTop(v) {
        if (this.scrollTop !== v) {
            this._properties.scrollTop.value = v;
            this.onScrollTop?.(this, v);
        }
    }

    set userSelect(v) {
        if (this.userSelect !== v) {
            this._properties.userSelect.value = v;
            this.ref.style.userSelect = v;
            this.ref.style['-webkit-user-select'] = v;
            this.ref.style['-ms-user-select'] = v;
        }
    }

    set v(v) {
        if (this.v !== v) {
            this._properties.v.value = v;
            this.ref.style.visibility = v ? 'visible' : 'hidden';
        }
    }

    set w(v) {
        if (this.w !== v) {
            this._properties.w.value = v;
            this.ref.style.width = v + 'px';
        }
    }

    set x(v) {
        if (this.x !== v) {
            this._properties.x.value = v;
            this.ref.style.left = v + 'px';
        }
    }

    set x2(v) {
        if (this.x2 !== v) {
            this._properties.x2.value = v;
        }
    }

    set y(v) {
        if (this.y !== v) {
            this._properties.y.value = v;
            this.ref.style.top = v + 'px';
        }
    }

    set y2(v) {
        if (this.y2 !== v) {
            this._properties.y2.value = v;
        }
    }

    set zIndex(v) {
        if (this.zIndex !== v) {
            this._properties.zIndex.value = v;
            this.ref.style.zIndex = v;
        }
    }

    set onCreatedFn(v) {
        if (this.onCreatedFn !== v) {
            this._properties.onCreatedFn.value = v;
        }
    }

    set onUpdatedFn(v) {
        if (this.onUpdatedFn !== v) {
            this._properties.onUpdatedFn.value = v;
        }
    }
}