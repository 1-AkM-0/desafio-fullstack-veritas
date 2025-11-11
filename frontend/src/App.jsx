import { useEffect, useState } from "react"
import { getTasks, createTask, deleteTask, updateTask } from "./services/taskService";
import Task from "./components/Task";
import Column from "./components/Column";

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
  const tasksByStatus = (status) => tasks.filter((task) => task.status === status)

  return (
    <div className="flex justify-center gap-6 p-6 min-h-screen bg-gray-200">
      <Column
        title={"A Fazer"}
        tasks={tasksByStatus("A Fazer")}
        onUpdate={handleUpdateTask}
        onDelete={handleDeleteTask}
        onAdd={handleCreateTask} />
      <Column
        title={"Em Progresso"}
        tasks={tasksByStatus("Em Progresso")}
        onUpdate={handleUpdateTask}
        onDelete={handleDeleteTask}
      />
      <Column
        title={"Concluídas"}
        tasks={tasksByStatus("Concluido")}
        onUpdate={handleUpdateTask}
        onDelete={handleDeleteTask}
      />
    </div>
  )
}

export default App
