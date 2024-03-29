import { useLocation } from "react-router-dom";
import Nav from "./components/Nav";

function App() {

  const { state } = useLocation();




  return (
    <div>
      <Nav />
      <div className="justify-center flex font-bold">
    <h1>{state.name ? state.name : "test"}</h1>
      </div>
    </div>
  );
}

export default App;
