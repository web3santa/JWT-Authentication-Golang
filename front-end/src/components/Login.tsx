import { SyntheticEvent, useState } from "react";
import Nav from "./Nav";

const Login = () => {
  const [email, setEamil] = useState("");
  const [password, setPassword] = useState("");

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    await fetch("http://localhost:3000/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        email,
        password,
      }),
    });


  };

  return (
    <div>
      <Nav />
      <div className="justify-center flex align-middle pt-10 pb-10 bg-red-200">
        <main>
          <form onSubmit={submit} className="w-100 max-w-96 p-4 m-auto">
            <h1>Please sign in</h1>
            <input
              className="relative border  border-black h-auto p-3 font-bold"
              type="email"
              id="inputEmail"
              placeholder="Email address"
              required
              onChange={(e) => setEamil(e.target.value)}
            />
            <input
              className="relative border  border-black h-auto p-3 font-bold"
              type="password"
              id="inputPassword"
              placeholder="Password"
              required
              onChange={(e) => setPassword(e.target.value)}
            />
            <br></br>
            <button className="border border-black w-full m-5" type="submit">
              Sign in
            </button>
          </form>
        </main>
      </div>
    </div>
  );
};

export default Login;
