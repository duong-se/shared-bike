import React from "react"
import { useNavigate } from 'react-router-dom'
import { useAuth } from "../hooks/AuthProvider"
import classNames from 'classnames'
import { Input } from "./Input";

type LoginFormProps = {
  onClickRegister: () => void;
}

export const LoginForm: React.FC<LoginFormProps> = ({ onClickRegister }) => {
  const { login, isLoading, error } = useAuth()
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
        <button
          className={classNames({ btn: true, loading: isLoading })}
        >
          Login
        </button>
        <div className="flex justify-evenly mt-5">
          <button onClick={onClickRegister} className="btn btn-link">Register</button>
        </div>
      </div>
      {error && (
        <div className="alert alert-error shadow-lg">
          <div>
            <svg xmlns="http://www.w3.org/2000/svg" className="stroke-current flex-shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span>{error}</span>
          </div>
        </div>
      )}
    </form>
  );
}
