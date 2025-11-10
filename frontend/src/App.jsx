import { useEffect, useState } from "react"
import { getTasks, createTask, deleteTask, updateTask } from "./services/taskService";

function App() {
  const [tasks, setTasks] = useState([])
  useEffect(() => {
    fetchTasks();
  }, [])

  const fetchTasks = async () => {
    const data = await getTasks()
    setTasks(data)
  }
  const handleCreateTask = async (title = "teste", description = "", status = "A Fazer") => {
    const newTask = await createTask(title, description, status)
    if (newTask) {
      setTasks([...tasks, newTask])
    }
  }
  const handleDeleteTask = async (id) => {
    await deleteTask(id)
    setTasks(tasks.filter((task) => task.id !== id))
  }
  const handleUpdateTask = async (id, updates) => {
    console.log(updates)
    await updateTask(id, updates)
    setTasks(prevTasks => prevTasks.map(task => {
      if (task.id !== id) {
        return task
      }
      return { ...task, ...updates }
    }))
  }

  return (
    <>
      {tasks.length > 0 ? <div>
        {tasks.map((task) => (
          <div key={task.id}>
            <div>
              {task.title}
            </div>
            <div>
              {task.description}
            </div>
            <div>
              {task.status}
            </div>
            -----------
            <button type="button" onClick={() => handleDeleteTask(task.id)}>Deletar</button>
            <button type="button" onClick={() => handleUpdateTask(task.id, { status: "Em Progresso" })}>Atualizar</button>
          </div>
        ))}
      </div> : <div>
        Nao há tasks
      </div>}

      <button type="button" onClick={() => handleCreateTask()}>Criar task</button>
    </>
  )
}

export default App
