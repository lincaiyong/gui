const g = {
    model: null,
    root: null,
    state: {},
    create() {
        g.log.info('create app.')
        document.documentElement.style.overflow = 'hidden';

        g.util.assert(g.model);
        Promise.all([]).then(() => {
            g.destroy();
            g._create();
        });
    },
    destroy() {
        if (g.root) {
            g.root._destroy();
            g.root = null;
        }
    },
    createElement(model, parent) {
        const ele = new model.Component(null, model);
        if (parent instanceof Component) {
            ele._create(parent);
        } else {
            g.log.error("invalid argument")
        }
        return ele;
    },
    createRootElement(parent) {
        g.root = new g.model.Component(null, g.model);
        if (parent instanceof Element) {
            g.root._create(parent);
        } else {
            g.log.error("invalid argument")
        }
    },
    _autoLayout() {
        const resize = () => [g.root.w, g.root.h] = [window.innerWidth, window.innerHeight];
        g.event.addListener(window, 'resize', g.util.debounce(resize, 20));
        resize();
    },
    _create() {
        g.createRootElement(document.body);
        g._autoLayout();
        g.root.v = 1;
        g.root._checkLoop();
    },
    removeElement(ele) {
        ele._destroy();
    },
    log: {
        level: 'debug',
        error() {
            console.error('[ERROR] ', ...arguments);
            document.body.innerText = [...arguments].join('\n');
        },
        info() {
            if (['info', 'debug', 'trace'].indexOf(g.log.level) !== -1) {
                console.log('[INFO ] ', ...arguments);
            }
        },
        debug() {
            if (['debug', 'trace'].indexOf(g.log.level) !== -1) {
                console.debug('[DEBUG] ', ...arguments);
            }
        },
        trace() {
            if (g.log.level === 'trace') {
                console.debug('[TRACE] ', ...arguments);
            }
        },
    },
    util: {
        assert(condition, failMsg = '') {
            if (!condition) {
                g.log.error(failMsg || 'assertion fail');
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
        fetch(url) {
            return new Promise((resolve, reject) => {
                fetch(url)
                    .then(resp => resp.text())
                    .then(v => resolve(v))
                    .catch(err => reject(err));
            });
        },
        _canvasCtx: undefined,
        textWidth(text, font, size = 12) {
            if (text === '') {
                return 0;
            }
            if (!g.util._canvasCtx) {
                const canvas = new OffscreenCanvas(1000, 40);
                g.util._canvasCtx = canvas.getContext("2d");
            }
            text = `=${text}=`
            const ctx = g.util._canvasCtx;
            ctx.font = `${size}px ${font}`;
            const metrics = ctx.measureText(text);
            const actual = Math.abs(metrics.actualBoundingBoxLeft) + Math.abs(metrics.actualBoundingBoxRight);
            const ret = Math.max(metrics.width, actual);
            return ret - 1.104 * size;
        },
    },
    event: {
        addListener(ref, name, handler) {
            if (handler instanceof Function) {
                g.log.trace('add event listener', name);
                ref.addEventListener(name, handler);
                return () => ref.removeEventListener(name, handler);
            }
        },
        onceListener(ref, name, handler) {
            if (handler instanceof Function) {
                g.log.trace('add once event listener', name);

                function handlerWrapper(ev) {
                    g.log.trace('once event listener removed', ref, name);
                    handler(ev);
                    ref.removeEventListener(name, handlerWrapper);
                }

                ref.addEventListener(name, handlerWrapper);
            }
        }
    },
    theme: {
        grayBorderColor: '#EBECF0',
        grayPaneColor: '#F7F8FA',
        dividerColor: '#C8CCD6',
        buttonColor: '#6C707E',
        buttonBgColor: '',
        buttonActiveBgColor: '#DFE1E4',
        buttonHoverBgColor: '#EBECF0',
        buttonSelectedBgColor: '#3475F0',
        buttonSelectedColor: '#FFFFFF',
        editorLineNoColor: '#AEB3C1',
        editorActiveLineColor: '#F6F8FE',
        editorSelectionColor: '#A6D2FF',
        editorHighlightColor: '#E6E6E6',
        editorBracketHighlightColor: '#93D9D9',
        scrollbarBgColor: '#7f7e80',
        treeFocusSelectedBgColor: '#D5E1FF',
        treeSelectedBgColor: '#DFE1E5',
    },
};