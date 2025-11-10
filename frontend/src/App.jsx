import { useEffect, useState } from "react"
import { getTasks, createTask } from "./services/taskService";

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
