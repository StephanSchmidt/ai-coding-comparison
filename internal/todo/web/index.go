package web

import (
	"fmt"
	"net/http"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, `<!DOCTYPE html>
		<html>
		<head>
		   	<script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.14.3/dist/cdn.min.js"></script>
		   	<script src="/js/todo.js"></script>
			<title>Todos</title>
		</head>
		<body>
			<h1>Todos</h1>
			<div x-data="retrieveData()" x-init="getData()">
				<input type="text" x-model="newItem" placeholder="Enter a new todo" />
				<button @click="if(newItem.trim() !== '') { item = { title: newItem }; todos.push(item); addTodo(item); newItem = ''; }">
					Add
				</button>
				<ul>
					<template x-for="(todo, index) in todos">
                    <li>
                        <span x-text="todo.title"></span>
                    </li>
                </template>
			</ul>
		</body>
		</html>`)
	}
}
