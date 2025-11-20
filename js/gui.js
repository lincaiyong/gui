class Property {
    static id = 0;

    constructor(element, name, sources, sourceResolver, computeFunc) {
        this._element = element;
        this._name = name;
        this._value = undefined;
        this._subscribers = [];
        this._sources = Array.from(sources || []);
        this._sourceResolver = sourceResolver;
        this._resolvedSources = [];
        this._computeFunc = computeFunc;
        this._updatedListeners = [];
        this._id = Property.id++;
    }

    get id() {
        return `${this._id}(${this._element.id}.${this._name})`;
    }

    reset(sources = null, computeFunc = null) {
        this.unsubscribe();
        this._sources = Array.from(sources || []);
        this.subscribe();
        if (computeFunc) {
            this._computeFunc = computeFunc;
        }
    }

    update() {
        for (const source of this._resolvedSources) {
            if (source._value === undefined) {
                return;
            }
        }
        this._element[this._name] = this._computeFunc(this._element);
    }

    subscribe() {
        this._resolvedSources = this._sources.map(source => this._sourceResolver(source));
        this._resolvedSources.forEach(
            source => source?._subscribers.push(this)
        );
    }

    unsubscribe() {
        this._resolvedSources.forEach(
            source => source?._subscribers.splice(source?._subscribers.indexOf(this), 1)
        );
    }

    get value() {
        return this._value;
    }

    set value(val) {
        if (this._value !== val && val !== undefined) {
            this._value = val;
            this._updatedListeners.forEach(fun => fun(val));
            this._subscribers.forEach(sub => sub.update());
        }
    }

    onUpdated(fun) {
        this._updatedListeners.push(fun);
        return () => this._updatedListeners.splice(this._updatedListeners.indexOf(fun), 1);
    }
}

class BaseElement {
    constructor(parent, model) {
        this.properties = {};
        this.parent = parent;
        this.model = model;
        this.tag = model.tag;
        this.id = parent ? `${parent.id}.${model.id}` : model.id;
        this.ref = document.createElement(model.tag);
        this.ref.style.position = model.position || 'absolute';
        this.ref.style.overflow = model.overflow || 'hidden';
        this.ref.style.boxSizing = 'border-box';
        this.children = (model.children || []).map(child => g.createElement(this, child));
        this._sideEffects = {};

        const props = Object.assign({}, this._defaultProperties, model.properties || {});
        for (const k in props) {
            const [computeFunc, sources] = props[k];
            const sourceResolver = source => this.resolveSrc(source);
            this.properties[k] = new Property(this, k, sources, sourceResolver, computeFunc);
            if (k === 'hovered') {
                this.properties[k].onUpdated(v => {
                    this.onHover?.(this, v);
                    this.onUpdated?.(k, v);
                });
            } else {
                this.properties[k].onUpdated(v => {
                    this.onUpdated?.(k, v);
                });
            }
        }
    }

    get local() {
        const depth = this.model.depth || 0;
        if (depth === 0) {
            return this;
        }
        let current = this;
        for (let i = 0; i < depth && current?.parent; i++) {
            current = current.parent;
        }
        return current || this;
    }

    get child() {
        return this.children.length > 0 ? this.children[0] : null;
    }

    get prev() {
        if (this.parent && this.parent.children && this.parent.children.length > 0) {
            const selfIndex = this.parent.children.indexOf(this);
            if (selfIndex > 0) {
                return this.parent.children[selfIndex - 1];
            }
        }
        return null;
    }

    get next() {
        if (this.parent && this.parent.children && this.parent.children.length > 0) {
            const selfIndex = this.parent.children.indexOf(this);
            if (selfIndex >= 0 && selfIndex < this.parent.children.length - 1) {
                return this.parent.children[selfIndex + 1];
            }
        }
        return null;
    }

    _createAll(parent) {
        if (parent instanceof Element) {
            parent.appendChild(this.ref);
        } else if (parent instanceof BaseElement) {
            this.parent = parent;
            parent.children.push(this);
            parent.ref.appendChild(this.ref);
        } else {
            console.error("invalid argument")
        }
        Object.values(this.properties).forEach(p => p.subscribe());
        this.children.forEach(child => child._createAll(this.ref));
    }

    _initAll() {
        Object.values(this.properties).forEach(p => p.update());
        this.children.forEach(child => child._initAll());
    }

    _create(parent) {
        this._createAll(parent);
        this._initAll();
        this.onCreated?.();
    }

    _unInitAll() {
        this.children.forEach(child => child._unInitAll());
        Object.values(this._sideEffects).forEach(fun => fun?.());
        Object.values(this.properties).forEach(p => p.unsubscribe());
    }

    _destroyAll() {
        this.children.forEach(child => child._destroyAll());
        if (this.parent?.children) {
            const index = this.parent.children.indexOf(this);
            if (index !== -1) {
                this.parent.children.splice(index, 1);
            }
        }
        this.ref.parentElement?.removeChild(this.ref);
    }

    _destroy() {
        this.onDestroy?.();
        this._unInitAll();
        this._destroyAll();
    }

    _checkLoop() {
        const properties = this._collectProperties();
        this._topologicalSort(properties);
    }

    _collectProperties() {
        let ret = Object.values(this.properties);
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
                console.error("loop detected", '\n\t' + tmp.join('\n\t'));
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

    resolveSrc(source) {
        if (!source.includes('.')) {
            return this.resolveEle(source);
        }
        const [e, p] = source.split('.', 2);
        const target = this.resolveEle(e);
        return target?.properties[p];
    }

    resolveEle(name) {
        if (name === '') {
            return this;
        } else if (name === 'root') {
            return g.root;
        } else if (['local', 'parent', 'child', 'prev', 'next'].includes(name)) {
            return this[name];
        }
        return null;
    }

    get _defaultProperties() {
        return {
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
            fontFamily: [() => '-apple-system, BlinkMacSystemFont, "SF Pro", "SF Pro Text", "PingFang SC", "Helvetica Neue", Arial, sans-serif', []],
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
    }

    get background() {
        return this.properties.background.value;
    }

    get backgroundColor() {
        return this.properties.backgroundColor.value;
    }

    get borderBottom() {
        return this.properties.borderBottom.value;
    }

    get borderColor() {
        return this.properties.borderColor.value;
    }

    get borderLeft() {
        return this.properties.borderLeft.value;
    }

    get borderRadius() {
        return this.properties.borderRadius.value;
    }

    get borderRight() {
        return this.properties.borderRight.value;
    }

    get borderStyle() {
        return this.properties.borderStyle.value;
    }

    get borderTop() {
        return this.properties.borderTop.value;
    }

    get boxShadow() {
        return this.properties.boxShadow.value;
    }

    get caretColor() {
        return this.properties.caretColor.value;
    }

    get ch() {
        return this.properties.ch.value;
    }

    get color() {
        return this.properties.color.value;
    }

    get cursor() {
        return this.properties.cursor.value;
    }

    get cw() {
        return this.properties.cw.value;
    }

    get fontFamily() {
        return this.properties.fontFamily.value;
    }

    get fontSize() {
        return this.properties.fontSize.value;
    }

    get fontWeight() {
        return this.properties.fontWeight.value;
    }

    get fontVariantLigatures() {
        return this.properties.fontVariantLigatures.value;
    }

    get h() {
        return this.properties.h.value;
    }

    get hovered() {
        return this.properties.hovered.value;
    }

    get hoveredByMouse() {
        return this.properties.hoveredByMouse.value;
    }

    get innerText() {
        return this.properties.innerText.value;
    }

    get placeholder() {
        return this.properties.placeholder.value;
    }

    get srcdoc() {
        return this.properties.srcdoc.value;
    }

    get src() {
        return this.properties.src.value;
    }

    get lineHeight() {
        return this.properties.lineHeight.value;
    }

    get onActive() {
        return this.properties.onActive.value;
    }

    get onClick() {
        return this.properties.onClick.value;
    }

    get onClickOutside() {
        return this.properties.onClickOutside.value;
    }

    get onCompositionEnd() {
        return this.properties.onCompositionEnd.value;
    }

    get onCompositionStart() {
        return this.properties.onCompositionStart.value;
    }

    get onCompositionUpdate() {
        return this.properties.onCompositionUpdate.value;
    }

    get onCopy() {
        return this.properties.onCopy.value;
    }

    get onCut() {
        return this.properties.onCut.value;
    }

    get onDoubleClick() {
        return this.properties.onDoubleClick.value;
    }

    get onFocus() {
        return this.properties.onFocus.value;
    }

    get onHover() {
        return this.properties.onHover.value;
    }

    get onInput() {
        return this.properties.onInput.value;
    }

    get onKeyDown() {
        return this.properties.onKeyDown.value;
    }

    get onKeyUp() {
        return this.properties.onKeyUp.value;
    }

    get onMouseDown() {
        return this.properties.onMouseDown.value;
    }

    get onMouseMove() {
        return this.properties.onMouseMove.value;
    }

    get onMouseUp() {
        return this.properties.onMouseUp.value;
    }

    get onPaste() {
        return this.properties.onPaste.value;
    }

    get onScrollLeft() {
        return this.properties.onScrollLeft.value;
    }

    get onScrollTop() {
        return this.properties.onScrollTop.value;
    }

    get onWheel() {
        return this.properties.onWheel.value;
    }

    get opacity() {
        return this.properties.opacity.value;
    }

    get outline() {
        return this.properties.outline.value;
    }

    get position() {
        return this.properties.position.value;
    }

    get scrollLeft() {
        return this.properties.scrollLeft.value;
    }

    get scrollTop() {
        return this.properties.scrollTop.value;
    }

    get userSelect() {
        return this.properties.userSelect.value;
    }

    get v() {
        return this.properties.v.value;
    }

    get w() {
        return this.properties.w.value;
    }

    get x() {
        return this.properties.x.value;
    }

    get x2() {
        return this.properties.x2.value;
    }

    get y() {
        return this.properties.y.value;
    }

    get y2() {
        return this.properties.y2.value;
    }

    get zIndex() {
        return this.properties.zIndex.value;
    }

    set background(v) {
        if (this.background !== v) {
            this.properties.background.value = v;
            this.ref.style.background = v;
        }
    }

    set backgroundColor(v) {
        if (this.backgroundColor !== v) {
            this.properties.backgroundColor.value = v;
            this.ref.style.backgroundColor = v;
        }
    }

    set borderBottom(v) {
        if (this.borderBottom !== v) {
            this.properties.borderBottom.value = v;
            this.ref.style.borderBottomWidth = v + 'px';
        }
    }

    set borderColor(v) {
        if (this.borderColor !== v) {
            this.properties.borderColor.value = v;
            this.ref.style.borderColor = v;
        }
    }

    set borderLeft(v) {
        if (this.borderLeft !== v) {
            this.properties.borderLeft.value = v;
            this.ref.style.borderLeftWidth = v + 'px';
        }
    }

    set borderRadius(v) {
        if (this.borderRadius !== v) {
            this.properties.borderRadius.value = v;
            this.ref.style.borderRadius = v + 'px';
        }
    }

    set borderRight(v) {
        if (this.borderRight !== v) {
            this.properties.borderRight.value = v;
            this.ref.style.borderRightWidth = v + 'px';
        }
    }

    set borderStyle(v) {
        if (this.borderStyle !== v) {
            this.properties.borderStyle.value = v;
            this.ref.style.borderStyle = v;
        }
    }

    set borderTop(v) {
        if (this.borderTop !== v) {
            this.properties.borderTop.value = v;
            this.ref.style.borderTopWidth = v + 'px';
        }
    }

    set boxShadow(v) {
        if (this.boxShadow !== v) {
            this.properties.boxShadow.value = v;
            this.ref.style.boxShadow = v;
        }
    }

    set caretColor(v) {
        if (this.caretColor !== v) {
            this.properties.caretColor.value = v;
            this.ref.style.caretColor = v;
        }
    }

    set ch(v) {
        if (this.ch !== v) {
            this.properties.ch.value = v;
        }
    }

    set color(v) {
        if (this.color !== v) {
            this.properties.color.value = v;
            this.ref.style.color = v;
        }
    }

    set cursor(v) {
        if (this.cursor !== v) {
            this.properties.cursor.value = v;
            this.ref.style.cursor = v;
        }
    }

    set cw(v) {
        if (this.cw !== v) {
            this.properties.cw.value = v;
        }
    }

    set fontFamily(v) {
        if (this.fontFamily !== v) {
            this.properties.fontFamily.value = v;
            this.ref.style.fontFamily = v;
        }
    }

    set fontSize(v) {
        if (this.fontSize !== v) {
            this.properties.fontSize.value = v;
            this.ref.style.fontSize = v + 'px';
        }
    }

    set fontWeight(v) {
        if (this.fontWeight !== v) {
            this.properties.fontWeight.value = v;
            this.ref.style.fontWeight = v;
        }
    }

    set fontVariantLigatures(v) {
        if (this.fontVariantLigatures !== v) {
            this.properties.fontVariantLigatures.value = v;
            this.ref.style.fontVariantLigatures = v;
        }
    }

    set h(v) {
        if (this.h !== v) {
            this.properties.h.value = v;
            this.ref.style.height = v + 'px';
        }
    }

    set hovered(v) {
        if (this.hovered !== v) {
            this.properties.hovered.value = v;
        }
    }

    set hoveredByMouse(v) {
        if (this.hoveredByMouse !== v) {
            this.properties.hoveredByMouse.value = v;
        }
    }

    set innerText(v) {
        if (typeof (v) === 'string' && this.tag === 'span') {
            this.properties.innerText.value = v;
            this.ref.innerText = v;
        }
    }

    set placeholder(v) {
        if (typeof (v) === 'string' && this.tag === 'input') {
            this.properties.placeholder.value = v;
            this.ref.placeholder = v;
        }
    }

    set srcdoc(v) {
        if (typeof (v) === 'string' && this.tag === 'iframe') {
            this.properties.srcdoc.value = v;
            this.ref.srcdoc = v;
        }
    }

    set src(v) {
        if (typeof (v) === 'string' && (this.tag === 'svg' || this.tag === 'img')) {
            this.properties.src.value = v;
            if (this.tag === 'img') {
                this.ref.setAttribute(k, v);
            } else {
                g.fetchRes(v).then(data => this.ref.innerHTML = data).catch(err => g.log.error(err));
            }
        }
    }

    set lineHeight(v) {
        if (this.lineHeight !== v) {
            this.properties.lineHeight.value = v;
            this.ref.style.lineHeight = v + 'px';
        }
    }

    set onActive(v) {
        if (v instanceof Function) {
            this._addSideEffect('onActive', g.addListener(this.ref, 'mousedown', ev => {
                const fun = v(this, ev);
                g.onceListener(this.ref, 'mouseup', ev => fun?.(this, ev));
            }));
        }
    }

    set onClick(v) {
        if (v instanceof Function) {
            this._addSideEffect('onClick', g.addListener(this.ref, 'click', ev => {
                v(this, ev);
            }));
        }
    }

    set onClickOutside(v) {
        if (v instanceof Function) {
            this._addSideEffect('onClickOutside', g.addListener(document, 'click', ev => {
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
            this._addSideEffect('onCompositionEnd', g.addListener(this.ref, 'compositionend', ev => {
                v(this, ev);
            }));
        }
    }

    set onCompositionStart(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCompositionStart', g.addListener(this.ref, 'compositionstart', ev => {
                v(this, ev);
            }));
        }
    }

    set onCompositionUpdate(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCompositionUpdate', g.addListener(this.ref, 'compositionupdate', ev => {
                v(this, ev);
            }));
        }
    }

    set onCopy(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCopy', g.addListener(this.ref, 'copy', ev => {
                v(this, ev);
            }));
        }
    }

    set onCut(v) {
        if (v instanceof Function) {
            this._addSideEffect('onCut', g.addListener(this.ref, 'cut', ev => {
                v(this, ev);
            }));
        }
    }

    set onDoubleClick(v) {
        if (v instanceof Function) {
            this._addSideEffect('onDoubleClick', g.addListener(this.ref, 'dblclick', ev => {
                v(this, ev);
            }));
        }
    }

    set onFocus(v) {
        if (v instanceof Function) {
            this._addSideEffect('onFocus', g.addListener(this.ref, 'focus', ev => {
                const fun = v(this, ev);
                g.onceListener(this.ref, 'blur', ev => fun?.(this, ev));
            }));
        }
    }

    set onHover(v) {
        if (v instanceof Function) {
            this.properties.onHover.value = v;
            this._addSideEffect('mouseenter', g.addListener(this.ref, 'mouseenter', () => {
                g.onceListener(this.ref, 'mouseleave', () => this.hoveredByMouse = false);
                this.hoveredByMouse = true;
            }));
        }
    }

    set onInput(v) {
        if (v instanceof Function) {
            this._addSideEffect('onInput', g.addListener(this.ref, 'input', ev => {
                v(this, ev);
            }));
        }
    }

    set onKeyDown(v) {
        if (v instanceof Function) {
            this._addSideEffect('onKeyDown', g.addListener(this.ref, 'keydown', ev => {
                v(this, ev);
            }));
        }
    }

    set onKeyUp(v) {
        if (v instanceof Function) {
            this._addSideEffect('onKeyUp', g.addListener(this.ref, 'keyup', ev => {
                v(this, ev);
            }));
        }
    }

    set onMouseDown(v) {
        if (v instanceof Function) {
            this._addSideEffect('onMouseDown', g.addListener(this.ref, 'mousedown', ev => {
                v(this, ev);
            }));
        }
    }

    set onMouseMove(v) {
        if (v instanceof Function) {
            this._addSideEffect('onMouseMove', g.addListener(this.ref, 'mousemove', ev => {
                v(this, ev);
            }));
        }
    }

    set onMouseUp(v) {
        if (v instanceof Function) {
            this._addSideEffect('onMouseUp', g.addListener(this.ref, 'mouseup', ev => {
                v(this, ev);
            }));
        }
    }

    set onPaste(v) {
        if (v instanceof Function) {
            this._addSideEffect('onPaste', g.addListener(this.ref, 'paste', ev => {
                v(this, ev);
            }));
        }
    }

    set onScrollLeft(v) {
        if (this.onScrollLeft !== v) {
            this.properties.onScrollLeft.value = v;
        }
    }

    set onScrollTop(v) {
        if (this.onScrollTop !== v) {
            this.properties.onScrollTop.value = v;
        }
    }

    set onWheel(v) {
        if (v instanceof Function) {
            this._addSideEffect('onWheel', g.addListener(this.ref, 'wheel', ev => {
                v(this, ev);
            }));
        }
    }

    set opacity(v) {
        if (this.opacity !== v) {
            this.properties.opacity.value = v;
            this.ref.style.opacity = v;
        }
    }

    set outline(v) {
        if (this.outline !== v) {
            this.properties.outline.value = v;
            this.ref.style.outline = v;
        }
    }

    set position(v) {
        if (this.position !== v) {
            this.properties.position.value = v;
            this.ref.style.position = v;
        }
    }

    set scrollLeft(v) {
        if (this.scrollLeft !== v) {
            this.properties.scrollLeft.value = v;
            this.onScrollLeft?.(this, v);
        }
    }

    set scrollTop(v) {
        if (this.scrollTop !== v) {
            this.properties.scrollTop.value = v;
            this.onScrollTop?.(this, v);
        }
    }

    set userSelect(v) {
        if (this.userSelect !== v) {
            this.properties.userSelect.value = v;
            this.ref.style.userSelect = v;
            this.ref.style['-webkit-user-select'] = v;
            this.ref.style['-ms-user-select'] = v;
        }
    }

    set v(v) {
        if (this.v !== v) {
            this.properties.v.value = v;
            this.ref.style.visibility = v ? 'visible' : 'hidden';
        }
    }

    set w(v) {
        if (this.w !== v) {
            this.properties.w.value = v;
            this.ref.style.width = v + 'px';
        }
    }

    set x(v) {
        if (this.x !== v) {
            this.properties.x.value = v;
            this.ref.style.left = v + 'px';
        }
    }

    set x2(v) {
        if (this.x2 !== v) {
            this.properties.x2.value = v;
        }
    }

    set y(v) {
        if (this.y !== v) {
            this.properties.y.value = v;
            this.ref.style.top = v + 'px';
        }
    }

    set y2(v) {
        if (this.y2 !== v) {
            this.properties.y2.value = v;
        }
    }

    set zIndex(v) {
        if (this.zIndex !== v) {
            this.properties.zIndex.value = v;
            this.ref.style.zIndex = v;
        }
    }
}

const g = {
    root: null,
    createAll(domElement, model) {
        document.documentElement.style.overflow = 'hidden';
        if (g.root) {
            g.root.destroyElement();
        }
        g.root = g.createElement(null, model);
        g.root._create(domElement);
        const resize = () => [g.root.w, g.root.h] = [window.innerWidth, window.innerHeight];
        g.addListener(window, 'resize', g.debounce(resize, 20));
        resize();
        g.root.v = 1;
        g.root._checkLoop();
    },
    createElement(parent, model) {
        const instance = new BaseElement(parent, model);
        for (const k in model.methods || {}) {
            instance[k] = model.methods[k];
        }
        return instance;
    },
    destroyElement(e) {
        e._destroy();
    },
    addListener(ref, name, handler) {
        if (handler instanceof Function) {
            ref.addEventListener(name, handler);
            return () => ref.removeEventListener(name, handler);
        }
    },
    onceListener(ref, name, handler) {
        if (handler instanceof Function) {
            function handlerWrapper(ev) {
                handler(ev);
                ref.removeEventListener(name, handlerWrapper);
            }

            ref.addEventListener(name, handlerWrapper);
        }
    },
    debounce(fun, interval) {
        let timer;
        return function () {
            const args = arguments;
            clearTimeout(timer);
            timer = setTimeout(() => fun.apply(this, args), interval);
        };
    },
    _canvasCtx: null,
    textWidth(text, font, size = 12) {
        if (text === '') {
            return 0;
        }
        if (!g._canvasCtx) {
            const canvas = new OffscreenCanvas(1000, 40);
            g._canvasCtx = canvas.getContext("2d");
        }
        text = `=${text}=`
        const ctx = g._canvasCtx;
        ctx.font = `${size}px ${font}`;
        const metrics = ctx.measureText(text);
        const actual = Math.abs(metrics.actualBoundingBoxLeft) + Math.abs(metrics.actualBoundingBoxRight);
        const ret = Math.max(metrics.width, actual);
        return ret - 1.104 * size;
    },
    fetch(url) {
        return new Promise((resolve, reject) => {
            fetch(url)
                .then(resp => resp.text())
                .then(v => resolve(v))
                .catch(err => reject(err));
        });
    },
    resCache: {},
    fetchRes(uri) {
        if (uri in g.resCache) {
            return g.resCache[uri];
        }
        return new Promise((resolve, reject) => {
            fetch(uri)
                .then(resp => resp.text())
                .then(v => {
                    g.resCache[uri] = v;
                    resolve(v);
                })
                .catch(err => reject(err));
        });
    }
};

// const root = {
//     tag: 'div',
//     id: 'div',
//     depth: 0,
//     properties: {
//         ch: [e => ((e.h - e.borderTop) - e.borderBottom), ['.borderBottom', '.borderTop', '.h']],
//         cw: [e => ((e.w - e.borderLeft) - e.borderRight), ['.borderLeft', '.borderRight', '.w']],
//         hovered: [e => e.hoveredByMouse, ['.hoveredByMouse']],
//         x2: [e => (e.x + e.w), ['.w', '.x']],
//         y2: [e => (e.y + e.h), ['.h', '.y']],
//     },
//     methods: {},
//     children: [
//         {
//             tag: 'span',
//             id: 'span',
//             depth: 1,
//             properties: {
//                 ch: [e => ((e.h - e.borderTop) - e.borderBottom), ['.borderBottom', '.borderTop', '.h']],
//                 cw: [e => ((e.w - e.borderLeft) - e.borderRight), ['.borderLeft', '.borderRight', '.w']],
//                 fontSize: [e => Math.floor(((e.h * 2) / 3)), ['.h']],
//                 h: [e => 200, []],
//                 hovered: [e => e.hoveredByMouse, ['.hoveredByMouse']],
//                 innerText: [e => 'hello world', []],
//                 lineHeight: [e => e.h, ['.h']],
//                 v: [e => e.parent.v, ['parent.v']],
//                 w: [e => g.textWidth(e.innerText, e.fontFamily, e.fontSize), ['.fontFamily', '.fontSize', '.innerText']],
//                 x: [e => ((e.parent.w / 2) - (e.w / 2)), ['.w', 'parent.w']],
//                 x2: [e => (e.x + e.w), ['.w', '.x']],
//                 y: [e => 100, []],
//                 y2: [e => (e.y + e.h), ['.h', '.y']],
//                 zIndex: [e => e.parent.zIndex, ['parent.zIndex']],
//             },
//             methods: {},
//             children: [],
//             slot: null,
//         },
//     ],
//     slot: null,
// };
//
// g.createAll(document.body, root);