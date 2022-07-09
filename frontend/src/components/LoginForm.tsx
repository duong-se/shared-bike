import React from "react"
import { useNavigate } from 'react-router-dom'
import { useAuth } from "../hooks/AuthProvider"
import { Button } from "./Button";
import { Input } from "./Input";

type LoginFormProps = {
  onClickRegister: () => void;
}

export const LoginForm: React.FC<LoginFormProps> = ({ onClickRegister }) => {
  const { login } = useAuth()
  let navigate = useNavigate()
  const handleSubmit = React.useCallback((e: React.SyntheticEvent) => {
    e.preventDefault()
    const values = e.target
    console.log(values)
    login("test1", "password", () => navigate("/dashboard"))
  }, [login, navigate])
  return (
    <form onSubmit={handleSubmit} className="w-full">
      <Input
        type="text"
        name="username"
        id="username"
        placeholder="Your username"
        label="Username"
      />
      <Input
        type="password"
        id="password"
        placeholder="Your password"
        name="password"
        label="Password"
      />
      <div id="button" className="flex flex-col w-full my-5">
        <Button type="submit" variant="primary">Login</Button>
        <div className="flex justify-evenly mt-5">
          <Button onClick={onClickRegister} variant="ghost">Register</Button>
        </div>
      </div>
    </form>
  );
}
