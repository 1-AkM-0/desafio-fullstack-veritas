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
