import { useState } from "react"

export default function Task({ task, onUpdate, onDelete, loadingIds }) {

  const isLoading = loadingIds.includes(task.id);

  const [isEditing, setIsEditing] = useState(false)

  const [editedTitle, setEditedTitle] = useState(task.title)
  const [editedDescription, setEditedDescription] = useState(task.description)

  const STATUS_TODO = "todo"
  const STATUS_IN_PROGRESS = "doing"
  const STATUS_DONE = "done"

  const handleSave = async () => {
    if (!editedTitle) return

    await onUpdate(task.id, {
      title: editedTitle,
      description: editedDescription,
      status: task.status
    })
    setIsEditing(false)
  }

  const handleCancel = () => {
    setIsEditing(false)
    setEditedTitle(task.title)
    setEditedDescription(task.description)
  }

  const ghostButtonStyles = "px-2 py-1 border rounded-md text-xs font-medium transition-colors duration-150"
  const formInputStyles = "w-full border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"

  const LoadingOverlay = () => (
    <div className="absolute inset-0 flex items-center justify-center bg-white/75">
    </div>
  )

  if (isEditing) {
    return (
      <div className="relative bg-white border-2 border-blue-500 rounded-lg p-4 shadow-lg flex flex-col gap-3">
        {isLoading && <LoadingOverlay />}
        <input
          type="text"
          value={editedTitle}
          onChange={(e) => setEditedTitle(e.target.value)}
          disabled={isLoading}
          required
          className={`${formInputStyles} px-3 py-2 text-base font-semibold`}
        />
        <textarea
          value={editedDescription}
          onChange={(e) => setEditedDescription(e.target.value)}
          disabled={isLoading}
          placeholder="Descrição"
          className={`${formInputStyles} px-3 py-2 text-sm min-h-[60px] resize-y`}
        />
        <div className="flex-gap-2 justify-end">
          <button
            onClick={handleCancel}
            disabled={isLoading}
            className={`${ghostButtonStyles} border-gray-400 text-gray-600 hover:bg-gray-100`}>
            Cancelar
          </button>
          <button
            onClick={handleSave}
            disabled={isLoading}
            className="px-3 py-1 rounded-md text-xs font-medium transition-colors duration-150 bg-green-500 text-white hover:bg-green-600 border border-green-500"
          >
            Salvar
          </button>
        </div>
      </div>
    )
  }



  return (
    <>
      <div className="relative bg-white border border-gray-200 rounded-lg p-4 shadow-sm flex flex-col gap-2">
        {isLoading && <LoadingOverlay />}
        <h4 className="text-base font-semibold text-gray-800">{task.title}</h4>
        {task.description && <p className="text-sm text-gray-600 whitespace-pre-wrap">{task.description}</p>}

        <div className="flex flex-wrap gap-2 justify-end pt-2 border-t border-gray-100 mt-2">
          <button
            onClick={() => setIsEditing(true)}
            disabled={isLoading}
            className={`${ghostButtonStyles} border-blue-500 text-blue-500 hover:bg-blue-500 hover:text-white`}
          >
            Editar
          </button>
          <button
            onClick={() => onDelete(task.id)}
            disabled={isLoading}
            className={`${ghostButtonStyles} border-red-500 text-red-500 hover:bg-red-500 hover:text-white`}
          >
            Deletar
          </button>
          {task.status === STATUS_TODO && (
            <button
              onClick={() => onUpdate(task.id, { status: STATUS_IN_PROGRESS })}
              disabled={isLoading}
              className={`${ghostButtonStyles} border-gray-400 text-gray-600 hover:bg-gray-100 `}
            >
              Iniciar
            </button>
          )}
          {task.status === STATUS_IN_PROGRESS && (
            <>
              <button
                onClick={() => onUpdate(task.id, { status: STATUS_TODO })}
                disabled={isLoading}
                className={`${ghostButtonStyles} border-gray-400 text-gray-600 hover:bg-gray-100 `}
              >
                Voltar
              </button>
              <button
                onClick={() => onUpdate(task.id, { status: STATUS_DONE })}
                disabled={isLoading}
                className={`${ghostButtonStyles} border-green-500 text-green-500 hover:bg-green-500 hover:text-white`}
              >
                Concluir
              </button>
            </>
          )}
          {task.status === STATUS_DONE && (
            <button
              onClick={() => onUpdate(task.id, { status: STATUS_IN_PROGRESS })}
              disabled={isLoading}
              className={`${ghostButtonStyles} border-gray-400 text-gray-600 hover:bg-gray-100`}
            >
              Reabrir
            </button>
          )}
        </div>
      </div>
    </>
  )

}
