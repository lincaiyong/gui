class Property {
    static id = 0;

    static alloc() {
        return Property.id++;
    }

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
        this._id = Property.alloc();
    }

    get id() {
        return `${this._id}(${this._element.id}.${this._name})`;
    }

    reset(sources=null, computeFunc=null) {
        g.util.assert(sources instanceof Array || sources === null);
        this.unsubscribe();
        this._sources = Array.from(sources || []);
        this.subscribe();
        if (computeFunc) {
            g.util.assert(computeFunc instanceof Function);
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

            // trace log
            if (g.log.level === 'trace') {
                g.log.trace(`update: ${this.id} = ${val}`);
                this._subscribers.forEach(sub => g.log.trace(`     -> ${this.id}`));
            }
        }
    }

    onUpdated(fun) {
        this._updatedListeners.push(fun);
        return () => this._updatedListeners.splice(this._updatedListeners.indexOf(fun), 1);
    }
}

class E {
    static root;
    static createAll(domElement, model) {
        document.documentElement.style.overflow = 'hidden';
        if (E.root) {
            E.root._destroy();
        }
        E.root = E.create(null, model);
        E.root._create(domElement);
        const resize = () => [E.root.w, E.root.h] = [window.innerWidth, window.innerHeight];
        E.addListener(window, 'resize', E.debounce(resize, 20));
        resize();
        E.root.v = 1;
        E.root._checkLoop();
    }
    static createElement(parent, model) {
        const instance = Object.create(E.prototype);
        E.call(instance, parent, model);
        for (const k in model.methods || {}) {
            instance[k] = model.methods[k];
        }
        return instance;
    }
    static destoryElement(e) {
        e._destroy();
    }
    static addListener(ref, name, handler) {
        if (handler instanceof Function) {
            ref.addEventListener(name, handler);
            return () => ref.removeEventListener(name, handler);
        }
    }
    static debounce(fun, interval) {
        let timer;
        return function () {
            const args = arguments;
            clearTimeout(timer);
            timer = setTimeout(() => fun.apply(this, args), interval);
        };
    }
    static _canvasCtx;
    static textWidth(text, font, size = 12) {
        if (text === '') {
            return 0;
        }
        if (!E._canvasCtx) {
            const canvas = new OffscreenCanvas(1000, 40);
            E._canvasCtx = canvas.getContext("2d");
        }
        text = `=${text}=`
        const ctx = E._canvasCtx;
        ctx.font = `${size}px ${font}`;
        const metrics = ctx.measureText(text);
        const actual = Math.abs(metrics.actualBoundingBoxLeft) + Math.abs(metrics.actualBoundingBoxRight);
        const ret = Math.max(metrics.width, actual);
        return ret - 1.104 * size;
    }
    constructor(parent, model) {
        this.properties = {};
        this.parent = parent;
        this.model = model;
        this.id = parent ? `${parent.id}.${model.id}` : model.id;
        this.ref = document.createElement(model.tag);
        this.ref.style.position = model.position || 'absolute';
        this.ref.style.overflow = model.overflow || 'hidden';
        this.ref.style.boxSizing = 'border-box';
        this.children = (model.children || []).map(child => E.create(this, child));
        this._sideEffects = {};

        for (const k in model.properties || {}) {
            const [computeFunc, sources] = model.properties[k];
            const sourceResolver = source => this._(source);
            this.properties[k] = new Property(this, k, sources, sourceResolver, computeFunc);
            if (k === 'hovered') {
                this.properties[k].onUpdated(v => { this.onHover?.(this, v); this.onUpdated?.(k, v); });
            } else {
                this.properties[k].onUpdated(v => { this.onUpdated?.(k, v); });
            }
        }
    }

    get root() {
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

    _createAll(parent) {
        if (parent instanceof Element) {
            parent.appendChild(this.ref);
        } else if (parent instanceof E) {
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
        this.onCreated?.();
    }

    _create(parent) {
        this._createAll(parent);
        this._initAll();
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

    _(source) {
        return this._resolve(source);
    }

    _resolve(source) {
        if (!source.includes('.')) {
            return this._resolveEle(source);
        }
        const [e, p] = source.split('.', 2);
        const target = this._resolveEle(e);
        return target?.properties[p];
    }

    _resolveEle(name) {
        if (name === '') {
            return this;
        } else if (name === 'root') {
            return E.root;
        } else if (name === 'this') {
            return this.root;
        } else if (name === 'parent') {
            return this.parent;
        } else if (name === 'child') {
            return this.children.?[0];
        } else if (name === 'prev' || name === 'next') {
            const siblings = this.parent?.children;
            if (!siblings) {
                return null;
            }
            const index = siblings.indexOf(this);
            if (index === -1) {
                return null;
            }
            const targetIndex = index + (name === 'prev' ? -1 : 1);
            if (targetIndex < 0 || targetIndex >= siblings.length) {
                return null;
            }
            return siblings[targetIndex];
        } else {
            const m = name.match(/^child([0-9])$/);
            if (m) {
                return this.children[parseInt(m[1])];
            }
            return this[name];
        }
    }
}