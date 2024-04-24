import { useState } from "react";

function App() {
  const [tasks, setTasks] = useState<string[]>([]);
  const [task, setTask] = useState<string>("");

  const addTask = () => {
    setTask("");
    setTasks([...tasks, task]);
  };
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
          {tasks.map((TASK) => (
            <div className="bg-black text-white mb-2 p-2 rounded">{TASK}</div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;
