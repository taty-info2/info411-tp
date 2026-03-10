import { Todos } from './todo.js';
import { TodosList } from './todo_vues.js';

export const API_BASE_URL = "http://localhost:3001";

const todos = new Todos();
await todos.getAll();
const todosList = new TodosList(todos);

document.body.append(todosList.el);
