import { useEffect, useState } from "react"
import { getTasks, createTask, deleteTask, updateTask } from "./services/taskService";
import Column from "./components/Column";

function App() {
  const [tasks, setTasks] = useState([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState(null)
  const [actionLoadingIds, setActionLoadingIds] = useState([])

  useEffect(() => {
    fetchTasks();
  }, [])

  const fetchTasks = async () => {
    setIsLoading(true)
    setError(null)
    try {
      const data = await getTasks()
      setTasks(data)
    } catch (error) {
      setError("Falha ao carregar as tasks")
    } finally {
      setIsLoading(false)
    }

  }

  const handleCreateTask = async (title = "teste", description = "", status = "todo") => {
    setError(null)
    try {
      const newTask = await createTask(title, description, status)
      if (newTask) {
        setTasks([...tasks, newTask])
      }
    } catch (error) {
      setError("Falha ao criar tarefa")
    }
  }

  const handleDeleteTask = async (id) => {
    setActionLoadingIds(prev => [...prev, id])
    setError(null)
    try {
      await deleteTask(id)
      setTasks(tasks.filter((task) => task.id !== id))
    } catch (error) {
      setError("Falha ao deletar tarefa")
    } finally {
      setActionLoadingIds(prev => prev.filter(taskId => taskId !== id))
    }

  }
  const handleUpdateTask = async (id, updates) => {
    console.log(updates)
    setActionLoadingIds(prev => [...prev, id])
    setError(null)
    try {
      await updateTask(id, updates)
      setTasks(prevTasks => prevTasks.map(task => {
        if (task.id !== id) {
          return task
        }
        return { ...task, ...updates }
      }))
    }
    catch (error) {
      setError("Falha ao atualizar tarefa")

    } finally {
      setActionLoadingIds(prev => prev.filter(taskId => taskId !== id))
    }

  }
  const tasksByStatus = (status) => tasks.filter((task) => task.status === status)

  if (isLoading) {
    return (
      <div className="p-2 text-center">
        <h2>Carregando tarefas...</h2>
      </div>
    )
  }

  return (
    <div className="flex justify-center gap-6 p-6 min-h-screen bg-gray-200">
      {error && <div className="fixed bottom-4 right-4 bg-red-600 text-white px-3 py-2 rounded-md shadow">
        {error}
      </div>}
      <Column
        title={"A Fazer"}
        tasks={tasksByStatus("todo")}
        onUpdate={handleUpdateTask}
        onDelete={handleDeleteTask}
        onAdd={handleCreateTask}
        loadingIds={actionLoadingIds}
      />
      <Column
        title={"Em Progresso"}
        tasks={tasksByStatus("doing")}
        onUpdate={handleUpdateTask}
        onDelete={handleDeleteTask}
        loadingIds={actionLoadingIds}
      />
      <Column
        title={"Concluídas"}
        tasks={tasksByStatus("done")}
        onUpdate={handleUpdateTask}
        onDelete={handleDeleteTask}
        loadingIds={actionLoadingIds}
      />
    </div>
  )
}

export default App
