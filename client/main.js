function getAPI() {
    return fetch(`/getStatus`)
            .then((res) => res.json())
            .then((data) => {
                console.log(data);
                return data;
            });
}

class ListItem {
    constructor(url, res) {
        this.url = url;
        this.status = res.status;
        this.code = res.status === 2 && !res.code ? 'timeout' : res.code;
    }

    getStatus() {
        if (this.status === 0) {
            return 'down';
        } else if (this.status === 1) {
            return 'up';
        } else {
            return 'slow';
        }
    }

    render() {
        let li = document.createElement('li');
        let label = document.createElement('div');
        let link = document.createElement('a');
        let statusDot = document.createElement('div');
        let statusCode = document.createElement('div');

        label.setAttribute('class', 'list-label');
        
        link.setAttribute('href', this.url);
        link.setAttribute('target', "_blank");
        link.textContent = this.url;
        
        statusDot.setAttribute('class', `status ${this.getStatus()}`);

        if (this.code) {
            const isTimeout = this.code === 'timeout' ? ' timeout': '';
            statusCode.setAttribute('class', `status-code${isTimeout}`);
            statusCode.textContent = this.code;
            statusDot.appendChild(statusCode)
        }

        label.appendChild(link);
        li.appendChild(label);
        li.appendChild(statusDot);

        return li;
    }
}

class List {
    constructor(data, className) {
        this.data = data;
        this.className = className || 'list-container';
    }

    render() {
        let listContainer = document.createElement('div');
        let list = document.createElement('ul');

        listContainer.setAttribute('class', this.className);
        list.setAttribute('class', 'list');

        Object.keys(this.data).map((key) => {
            list.appendChild(new ListItem(key, this.data[key]).render());
        })

        listContainer.appendChild(list);

        return listContainer;
    }
}

class Loading {
    constructor() {}

    render() {
        let l = document.createElement('div');
        l.textContent = "Loading...";

        return l;
    }
}

window.onload = async () => {
    const app = document.getElementById('app');
    const loading = new Loading().render();
    app.appendChild(loading);
    
    const data = await getAPI();
    app.removeChild(loading);
    
    app.appendChild(new List(data).render());
}