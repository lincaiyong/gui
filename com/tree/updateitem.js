function updateItem(itemEle, k, v) {
    if (k === 'data') {
        if (v.leaf) {
            const ext = v.key.substring(v.key.lastIndexOf('.')+1);
            let src = 'svg/text.svg';
            switch (ext) {
                case 'go':
                    src = 'svg/go.svg';
                    break;
                case 'js':
                    src = 'svg/js.svg';
                    break;
                case 'py':
                    src = 'svg/python.svg';
                    break;
            }
            if (v.key === 'go.mod' || v.key.endsWith('/go.mod')) {
                src = 'svg/goMod.svg'
            }
            if (v.key === '.gitignore' || v.key.endsWith('/.gitignore')) {
                src = 'svg/ignored.svg'
            }
            itemEle.iconEle.src = src;
        } else {
            itemEle.iconEle.src = 'svg/folder.svg';
        }
    }
}