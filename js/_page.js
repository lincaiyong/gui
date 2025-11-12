const page = {
    model: null,
    root: null,
    state: {},
    create() {
        page.log.info('create app.')
        document.documentElement.style.overflow = 'hidden';

        page.util.assert(page.model);
        Promise.all([]).then(() => {
            page.destroy();
            page._create();
        });
    },
    destroy() {
        if (page.root) {
            page.root._destroy();
            page.root = null;
        }
    },
    createElement(model, parent) {
        const ele = new model.Component(null, model);
        if (parent instanceof Component) {
            ele._create(parent);
        } else {
            page.log.error("invalid argument")
        }
        return ele;
    },
    createRootElement(parent) {
        page.root = new page.model.Component(null, page.model);
        if (parent instanceof Element) {
            page.root._create(parent);
        } else {
            page.log.error("invalid argument")
        }
    },
    _autoLayout() {
        const resize = () => [page.root.w, page.root.h] = [window.innerWidth, window.innerHeight];
        page.event.addListener(window, 'resize', page.util.debounce(resize, 20));
        resize();
    },
    _create() {
        page.createRootElement(document.body);
        page._autoLayout();
        page.root.v = 1;
        page.root._checkLoop();
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
            if (['info', 'debug', 'trace'].indexOf(page.log.level) !== -1) {
                console.log('[INFO ] ', ...arguments);
            }
        },
        debug() {
            if (['debug', 'trace'].indexOf(page.log.level) !== -1) {
                console.debug('[DEBUG] ', ...arguments);
            }
        },
        trace() {
            if (page.log.level === 'trace') {
                console.debug('[TRACE] ', ...arguments);
            }
        },
    },
    util: {
        assert(condition, failMsg = '') {
            if (!condition) {
                page.log.error(failMsg || 'assertion fail');
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
            if (!page.util._canvasCtx) {
                const canvas = new OffscreenCanvas(1000, 40);
                page.util._canvasCtx = canvas.getContext("2d");
            }
            text = `=${text}=`
            const ctx = page.util._canvasCtx;
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
                page.log.trace('add event listener', name);
                ref.addEventListener(name, handler);
                return () => ref.removeEventListener(name, handler);
            }
        },
        onceListener(ref, name, handler) {
            if (handler instanceof Function) {
                page.log.trace('add once event listener', name);

                function handlerWrapper(ev) {
                    page.log.trace('once event listener removed', ref, name);
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