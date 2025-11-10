export const getTasks = async () => {
  try {
    const response = await fetch("http://localhost:5000/tasks", {
      method: "GET",
    });
    if (!response.ok) {
      throw new Error("A requisição falhou");
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Falha ao buscar tasks:", error);
    return [];
  }
};
export const createTask = async (title, description, status) => {
  try {
    const response = await fetch("http://localhost:5000/tasks", {
      method: "POST",
      body: JSON.stringify({ title, description, status }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (!response.ok) {
      throw new Error("A requisição falhou");
    }
    const data = await response.json();
    console.log(data);
    return data;
  } catch (error) {
    console.error("Falha ao criar task", error);
  }
};

export const deleteTask = async (id) => {
  try {
    const response = await fetch(`http://localhost:5000/tasks/${id}`, {
      method: "DELETE",
    });
    if (!response.ok) {
      throw new Error("A requisicão falhou");
    }
  } catch (error) {
    console.error("Falha ao deletar task", error);
  }
};
