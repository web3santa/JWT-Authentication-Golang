import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Nav = () => {
  const [name, setName] = useState("")
  let menu;

  const logout = async() => {
    console.log("Logout");
    
    await fetch("http://localhost:3000/api/logout", {
      method: "GET",
      headers: {"Content-Type": "application/json"},
      credentials: "include"
    })
    
  }

  if (name == "") {
    menu = (
      <div className="flex justify-center font-bold p-5 m-5 gap-10">

        <Link to={"/"} state={{name: name}} >Home</Link>
        <Link to={"/login"}>Login</Link>
        <Link to={"/register"}>Register</Link>

      </div>
    )
  } else {
    menu = (
      <div className="flex justify-center font-bold p-5 m-5 gap-10">

        <Link to={"/"} state={{name: name}} >Home</Link>
        <Link to={"/login"}>Login</Link>
        <Link to={"/register"}>Register</Link>
        <button onClick={logout}>Logout</button>

      </div>
    )
  }


  useEffect(() => {
    (
      async() => {
        
      const response = await fetch("http://localhost:3000/api/user", {
        method: "GET",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
  
      const content = await response.json()
      setName(content.name)
      }
    )()
  },[name])
  return (
    <div>
      {menu}
    </div>
  )
};

export default Nav;
