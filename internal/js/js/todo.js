console.log("todo.js loaded");

function retrieveData() {
    return {
        isLoading: false,
        todos: [],
        newItem: '',
        getData() {
            this.isLoading = true;
            fetch('/list/')
                .then((response) => response.json())
                .then((data) => {
                    this.todos = data;
                    this.isLoading = false;
                    this.newItem = '';
                });
        }
    }
}

function addTodo(data) {
    console.log("adding todo", data);
    fetch("/add/", {
        method: 'POST', // Specify the HTTP method
        headers: {
            'Content-Type': 'application/json', // Inform the server that the body contains JSON
        },
        body: JSON.stringify(data), // Convert the JavaScript object to JSON string
    })
}
