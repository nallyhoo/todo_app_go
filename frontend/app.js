const API_URL = 'http://localhost:8080/todos';
let todos = [];

document.addEventListener('DOMContentLoaded', () => {
    fetchTodos();
    document.getElementById('todo-form').addEventListener('submit', handleFormSubmit);
    document.getElementById('cancel-edit').addEventListener('click', resetForm);
    setInterval(updateProgress, 1000);
});

async function fetchTodos() {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) throw new Error(`HTTP error: ${response.status}`);
        todos = await response.json();
        renderTodos();
    } catch (error) {
        console.error('Error fetching todos:', error);
        alert('Failed to fetch todos: ' + error.message);
    }
}

function renderTodos() {
    const todoList = document.getElementById('todo-list');
    todoList.innerHTML = '';
    todos.forEach(todo => {
        const li = document.createElement('li');
        li.className = todo.completed ? 'completed' : '';
        li.innerHTML = `
            <div>
                <span class="title">${todo.title}</span>
                <div class="details">
                    <small>Start: ${todo.start_time ? new Date(todo.start_time).toLocaleString() : 'N/A'}</small>
                    <small>End: ${todo.end_time ? new Date(todo.end_time).toLocaleString() : 'N/A'}</small>
                    <small>Progress: ${todo.progress.toFixed(1)}%</small>
                    <div class="progress-bar">
                        <div class="progress" style="width: ${todo.progress}%"></div>
                    </div>
                </div>
            </div>
            <div>
                <button onclick="editTodo(${todo.id})">Edit</button>
                <button class="delete" onclick="deleteTodo(${todo.id})">Delete</button>
            </div>
        `;
        todoList.appendChild(li);
    });
}

function calculateProgress(todo) {
    if (todo.completed) return 100;
    if (todo.progress >= 0 && todo.progress <= 100) return todo.progress;
    if (!todo.start_time || !todo.end_time) return 0;
    
    const now = new Date();
    const start = new Date(todo.start_time);
    const end = new Date(todo.end_time);
    
    if (now < start) return 0;
    if (now > end) return 100;
    
    const totalDuration = end - start;
    const elapsedDuration = now - start;
    if (totalDuration <= 0) return 0;
    
    return (elapsedDuration / totalDuration) * 100;
}

function updateProgress() {
    todos = todos.map(todo => ({
        ...todo,
        progress: calculateProgress(todo)
    }));
    renderTodos();
}

async function handleFormSubmit(event) {
    event.preventDefault();
    const id = document.getElementById('todo-id').value;
    const title = document.getElementById('todo-title').value;
    const description = document.getElementById('todo-description').value;
    const startTime = document.getElementById('todo-start-time').value;
    const endTime = document.getElementById('todo-end-time').value;
    const progress = parseFloat(document.getElementById('todo-progress').value) || -1;
    const completed = document.getElementById('todo-completed').checked;

    if (progress > 100 || (progress < 0 && progress !== -1)) {
        alert('Progress must be between 0 and 100');
        return;
    }

    const todo = {
        title,
        description,
        completed,
        start_time: startTime || null,
        end_time: endTime || null,
        progress: progress === -1 ? 0 : progress
    };

    try {
        const method = id ? 'PUT' : 'POST';
        const url = id ? `${API_URL}/${id}` : API_URL;
        const response = await fetch(url, {
            method,
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(todo)
        });

        if (response.ok) {
            resetForm();
            fetchTodos();
        } else {
            const errorText = await response.text();
            console.error('Error saving todo:', errorText);
            alert(`Failed to save TODO: ${errorText}`);
        }
    } catch (error) {
        console.error('Error:', error);
        alert(`Network error: ${error.message}`);
    }
}

async function editTodo(id) {
    try {
        const response = await fetch(`${API_URL}/${id}`);
        if (!response.ok) throw new Error(`HTTP error: ${response.status}`);
        const todo = await response.json();
        document.getElementById('todo-id').value = todo.id;
        document.getElementById('todo-title').value = todo.title;
        document.getElementById('todo-description').value = todo.description;
        document.getElementById('todo-completed').checked = todo.completed;
        document.getElementById('todo-start-time').value = todo.start_time ? new Date(todo.start_time).toISOString().slice(0, 16) : '';
        document.getElementById('todo-end-time').value = todo.end_time ? new Date(todo.end_time).toISOString().slice(0, 16) : '';
        document.getElementById('todo-progress').value = todo.progress > 0 ? todo.progress : '';
        document.getElementById('cancel-edit').style.display = 'inline';
        document.getElementById('todo-title').focus();
    } catch (error) {
        console.error('Error fetching todo:', error);
        alert('Failed to fetch todo: ' + error.message);
    }
}

async function deleteTodo(id) {
    if (!confirm('Are you sure you want to delete this TODO?')) {
        return; // Cancel deletion if user clicks "Cancel"
    }

    try {
        const response = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
        if (response.ok) {
            fetchTodos();
        } else {
            const errorText = await response.text();
            console.error('Error deleting todo:', errorText);
            alert(`Failed to delete TODO: ${errorText}`);
        }
    } catch (error) {
        console.error('Error:', error);
        alert(`Network error: ${error.message}`);
    }
}

function resetForm() {
    document.getElementById('todo-form').reset();
    document.getElementById('todo-id').value = '';
    document.getElementById('cancel-edit').style.display = 'none';
}