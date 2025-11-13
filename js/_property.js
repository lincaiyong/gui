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