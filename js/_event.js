class Event {
    static listeners = {};
    static trace = false;
    static onceListeners = {};

    static _rand() {
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
            const r = (Math.random() * 16) | 0,  v = c === 'x' ? r : (r & 0x3) | 0x8;
            return v.toString(16);
        });
    }

    static onOnce(handler) {
        const name = Event._rand();
        Event.onceListeners[name] = handler;
        return name;
    }

    static emitOnce(name, data) {
        const fun = Event.onceListeners[name];
        if (fun) {
            fun(data);
            delete Event.onceListeners[name];
        }
    }

    static on(name, handler) {
        let arr = this.listeners[name];
        if (!arr) {
            arr = [];
            this.listeners[name] = arr;
        }
        arr.push(handler);
        return () => arr.splice(arr, arr.indexOf(handler), 1);
    }

    static emit(name, data) {
        if (Event.trace) {
            console.log(name);
        }
        const arr = this.listeners[name] || [];
        arr.forEach(fun => fun(data));
    }
}