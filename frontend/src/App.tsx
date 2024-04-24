import { useEffect, useState } from "react";

function App() {
  interface taskObj {
    id: number;
    title: string;
    done: boolean;
  }

  const [tasks, setTasks] = useState<taskObj[]>([]);
  const [task, setTask] = useState<string>("");

  async function getTasks() {
    try {
      const res = await fetch("http://localhost:8080/api/all");
      const data = await res.json();
      // console.log(data);
      if (data.length > 0) {
        setTasks(data);
      }
    } catch (error) {
      console.log(error);
    }
  }

  const addTask = async () => {
    const newTask: taskObj = {
      id: tasks.length + 1,
      title: task,
      done: false,
    };
    setTask("");
    try {
      const res = await fetch("http://localhost:8080/api/add", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(newTask),
      });
      const data = await res.json();
      if (data) {
        setTasks([...tasks, newTask]);
      }
    } catch (error) {
      console.log(error);
    }
  };

  const deleteTask = async (id: number) => {
    try {
      const res = await fetch(`http://localhost:8080/api/delete/${id}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
      });
      const data = await res.json();
      getTasks();
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    getTasks();
  }, [deleteTask]);

  return (
    <div className="bg-[#333] h-screen">
      <div className="container py-4 size-full">
        <h1 className="text-center text-[crimson] text-4xl font-bold mb-3">GO TO DO APP</h1>
        <div className="bg-slate-600 w-2/3 m-auto p-4 rounded">
          <input
            placeholder="ur task?"
            type="text"
            value={task}
            onChange={(e) => setTask(e.target.value)}
            className="p-2 rounded outline-none mb-2"
          />
          <input
            value="+"
            type="submit"
            className="bg-green-600 text-white p-2 rounded cursor-pointer"
            onClick={addTask}
          />
          {tasks?.map((task) => (
            <div
              key={task.id}
              onClick={() => deleteTask(task.id)}
              className="bg-black text-white mb-2 p-2 rounded"
            >
              {task.title}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;
