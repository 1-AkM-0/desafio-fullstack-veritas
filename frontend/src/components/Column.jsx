import { useState } from "react";
import Task from "./Task";

export default function Column({ title, tasks, onUpdate, onDelete, onAdd }) {

  const [newTitle, setNewTitle] = useState("")
  const [newDescription, setNewDescription] = useState("")

  const handleSubmit = (e) => {
    e.preventDefault()
    if (!newTitle) return
    onAdd(newTitle, newDescription)
    setNewTitle("")
    setNewDescription("")

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
        {onAdd && (
          <form onSubmit={handleSubmit} className="flex flex-col gap-3">
            <input
              type="text"
              value={newTitle}
              onChange={(e) => setNewTitle(e.target.value)}
              placeholder="Título da tarefa"
              required
              className={formInputStyles} />
            <textarea
              value={newDescription}
              onChange={(e) => setNewDescription(e.target.value)}
              placeholder="Descrição"
              rows="2"
              className={`${formInputStyles} resize-none`}
            />
            <button
              type="submit"
              className="w-full bg-blue-600 text-white p-2 rounded-md font-semibold hover:bg-blue-700 transition-colors"
            >
              Adicionar Tarefa
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
            />
          ))}
        </div>
      </div>
    </div>
  )

}
