import { Link } from "react-router-dom";

const Nav = () => {
  return (
    <div className="flex justify-center font-bold p-5 m-5 gap-10">
      <Link to={"/"}>Home</Link>
      <Link to={"/login"}>Login</Link>
      <Link to={"/register"}>Register</Link>
    </div>
  );
};

export default Nav;
