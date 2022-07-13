import React from 'react'
import { NavigateFunction, useNavigate } from 'react-router-dom'
import jwtDecode from 'jwt-decode'
import cx from 'classnames'
import {
  useFormik,
} from 'formik'
import * as Yup from 'yup'
import { useAuth } from '../hooks/AuthProvider'
import { LoginResponse, User } from '../typings/types'
import { Input } from './Input'
import { tokenKey } from '../constants/constants'
import { AlertError } from './AlertError'
import { useLogin } from '../hooks/useUsers'

type LoginFormProps = {
  onClickRegister: () => void;
}

const LoginSchema = Yup.object().shape({
  username: Yup.string().required('Username is required'),
  password: Yup.string().required('Password is required'),
})

export const onSuccessHandler = (
  setUser: React.Dispatch<React.SetStateAction<User | undefined>>,
  navigate: NavigateFunction
) => (data: LoginResponse) => {
  localStorage.setItem(tokenKey, data.accessToken)
  const decoded = jwtDecode<User>(data.accessToken)
  setUser(decoded)
  navigate('/dashboard', { replace: true })
}

export const LoginForm: React.FC<LoginFormProps> = ({ onClickRegister }) => {
  const { setUser } = useAuth()
  const navigate = useNavigate()
  const { isLoading, mutate, error, isError } = useLogin({
    onSuccess: onSuccessHandler(setUser, navigate),
  })
  const formik = useFormik({
    validationSchema: LoginSchema,
    initialValues: {
      username: '',
      password: '',
    },
    onSubmit: values => {
      const username = values.username
      const password = values.password
      mutate({ username: username, password: password })
    },
  })
  return (
    <form onSubmit={formik.handleSubmit}>
      <Input
        type="text"
        name="username"
        id="username"
        placeholder="Your username"
        label="Username"
        error={formik.touched.username ? formik.errors.username : undefined}
        value={formik.values.username}
        onChange={formik.handleChange}
      />
      <Input
        type="password"
        id="password"
        placeholder="Your password"
        name="password"
        label="Password"
        error={formik.touched.password ? formik.errors.password : undefined}
        value={formik.values.password}
        onChange={formik.handleChange}
      />
      <div id="button" className="flex flex-col w-full my-5">
        <button
          type="submit"
          className={cx({ btn: true, loading: isLoading })}
        >
          Login
        </button>
        <div className="flex justify-evenly mt-5">
          <button onClick={onClickRegister} className="btn btn-link">Register</button>
        </div>
      </div>
      {isError && <AlertError error={(error as Error).message} />}
    </form>
  )
}
