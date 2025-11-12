import { useState } from "react";
import Task from "./Task";

export default function Column({ title, tasks, onUpdate, onDelete, onAdd, loadingIds }) {

  const [newTitle, setNewTitle] = useState("")
  const [newDescription, setNewDescription] = useState("")
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState(null)

  const handleSubmit = async (e) => {
    e.preventDefault()
    console.log(isLoading)
    setIsLoading(true)
    setError(null)
    if (!newTitle) {
      setError("Título é obrigatório")
      setIsLoading(false)
      return
    }
    try {
      console.log(isLoading)
      await onAdd(newTitle, newDescription)
      setNewTitle("")
      setNewDescription("")

    } catch (error) {
      setError("Erro ao adicionar tarefa")
    }
    finally {
      setIsLoading(false)
      console.log(isLoading)
    }


  }
  const headerColors = {
    "A Fazer": "bg-blue-500",
    "Em Progresso": "bg-yellow-500",
    "Concluídas": "bg-green-500"
  }
  const bgColor = headerColors[title] || "bg-gray-400"

  const formInputStyles = "w-full border border-gray-300 rounded-md p-2 focus:outline-none focus:rind-2 focus:ring-blue-500"

  return (
    <div className="flex flex-col w-full min-w-[280px] max-w-[320px] bg-gray-100 rounded-lg shadow-md">
      <div className={`p-3 rounded-t-lg ${bgColor}`}>
        <h2 className="text-white text-lg font-bold text-center uppercase tracking-wide">
          {title}
        </h2>
      </div>
      <div className="p-4 flex flex-col gap-4">
        {error && <div className="border border-red-400 bg-red-50 text-red-700 text-sm p-2 rounded-md text-center ">
          {error}
        </div>}
        {onAdd && (
          <form onSubmit={handleSubmit} className="flex flex-col gap-3">
            <input
              type="text"
              value={newTitle}
              onChange={(e) => setNewTitle(e.target.value)}
              disabled={isLoading}
              placeholder="Título da tarefa"
              required
              className={formInputStyles} />
            <textarea
              value={newDescription}
              onChange={(e) => setNewDescription(e.target.value)}
              disabled={isLoading}
              placeholder="Descrição"
              rows="2"
              className={`${formInputStyles} resize-none`}
            />
            <button
              type="submit"
              className="w-full bg-blue-600 text-white p-2 rounded-md font-semibold hover:bg-blue-700 transition-colors"
              disabled={isLoading}
            >
              {isLoading ? "Adicionando..." : "Adicionar Tarefa"}
            </button>
          </form>
        )}
        <div className="flex flex-col gap-3 overflow-y-auto max-h-[70vh]">
          {tasks.map((task) => (
            <Task
              key={task.id}
              task={task}
              onUpdate={onUpdate}
              onDelete={onDelete}
              loadingIds={loadingIds}
            />
          ))}
        </div>
      </div>
    </div>
  )

}
