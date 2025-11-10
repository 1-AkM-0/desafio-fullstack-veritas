import { useEffect, useState } from "react"
import { getTasks } from "./services/taskService";

function App() {
  const [tasks, setTasks] = useState([])
  useEffect(() => {
    fetchTasks();
  }, [])

  const fetchTasks = async () => {
    const data = await getTasks()
    setTasks(data)
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
    </>
  )
}

export default App
