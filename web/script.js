const API = "http://localhost:8080/tasks"

async function loadTasks() {

    const res = await fetch(API)
    const tasks = await res.json()

    const container = document.getElementById("tasks")
    container.innerHTML = ""

    tasks.forEach(task => {

        const div = document.createElement("div")
        div.className = "task"

        div.innerHTML = `
            <b>${task.title}</b><br>
            ${task.description}<br>

            <span class="status ${task.status}">
            ${task.status}
            </span>

            <br>

            <button onclick="editTask(${task.id}, '${task.title}', '${task.description}')">
            Edit
            </button>

            <button onclick="changeStatus(${task.id}, '${task.status}', '${task.title}', '${task.description}')">
            Change Status
            </button>

            <button onclick="deleteTask(${task.id})">
            Delete
            </button>
        `

        container.appendChild(div)
    })
}

async function createTask() {

    const title = document.getElementById("title").value
    const description = document.getElementById("description").value

    await fetch(API, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            title: title,
            description: description,
            status: "todo"
        })
    })

    loadTasks()
}

async function deleteTask(id) {

    await fetch(`${API}/${id}`, {
        method: "DELETE"
    })

    loadTasks()
}

async function editTask(id, title, description){
    const newTitle = prompt("New title:", title)
    const newDescription = prompt("New description:", description)

    if (!newTitle) return
    await fetch(`${API}/${id}`,{
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            title: newTitle,
            description: newDescription,
            status: "todo"
        })
    })
    loadTasks()
}

async function changeStatus(id, status, title, description) {
    let newStatus = "todo"

    if (status === "todo"){
        newStatus = "in_progress"
    } else if (status==='in_progress'){
        newStatus = "done"
    } else{
        newStatus = "todo"
    }
    await fetch(`${API}/${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            title: title,
            description: description,
            status: newStatus
        })
    })
    loadTasks()
}

loadTasks()