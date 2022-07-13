import { useFormik } from 'formik'
import { NavigateFunction, useNavigate } from 'react-router-dom'
import jwtDecode from 'jwt-decode'
import * as Yup from 'yup'
import { useAuth } from '../hooks/AuthProvider'
import cx from 'classnames'
import { Input } from './Input'
import { RegisterResponse, User } from '../typings/types'
import { tokenKey } from '../constants/constants'
import { AlertError } from './AlertError'
import { useRegister } from '../hooks/useUsers'

type RegisterFormProps = {
  onClickLogin: () => void
}

const RegisterSchema = Yup.object().shape({
  username: Yup.string().required('Username is required'),
  password: Yup.string().required('Password is required'),
  confirmPassword: Yup.string().oneOf([Yup.ref('password'), null], 'Confirm passwords must match').required('Confirm password is required'),
  name: Yup.string().required('Name is required'),
})

export const onSuccessHandler = (
  setUser: React.Dispatch<React.SetStateAction<User | undefined>>,
  navigate: NavigateFunction
) => (data: RegisterResponse) => {
  localStorage.setItem(tokenKey, data.accessToken)
  const decoded = jwtDecode<User>(data.accessToken)
  setUser(decoded)
  navigate('/dashboard', { replace: true })
}



export const RegisterForm: React.FC<RegisterFormProps> = ({ onClickLogin }) => {
  const { setUser } = useAuth()
  const navigate = useNavigate()
  const { isLoading, mutate, error, isError } = useRegister({
    onSuccess: onSuccessHandler(setUser, navigate),
  })
  const formik = useFormik({
    validationSchema: RegisterSchema,
    initialValues: {
      username: '',
      password: '',
      confirmPassword: '',
      name: '',
    },
    onSubmit: values => {
      mutate({ username: values.username, password: values.password, name: values.name })
    },
  })
  return (
    <form onSubmit={formik.handleSubmit} className="w-full">
      <Input
        type="text"
        id="username"
        label="Username"
        placeholder="Your username"
        name="username"
        error={formik.errors.username}
        value={formik.values.username}
        onChange={formik.handleChange}
      />
      <Input
        type="password"
        id="password"
        placeholder="Your password"
        name="password"
        label="Password"
        error={formik.errors.password}
        value={formik.values.password}
        onChange={formik.handleChange}
      />
      <Input
        type="password"
        id="confirmPassword"
        name="confirmPassword"
        label="Confirm Password"
        placeholder="Your password confirm"
        error={formik.errors.confirmPassword}
        value={formik.values.confirmPassword}
        onChange={formik.handleChange}
      />
      <Input
        type="text"
        id="name"
        name="name"
        label="Name"
        placeholder="Your name"
        error={formik.errors.name}
        value={formik.values.name}
        onChange={formik.handleChange}
      />
      <div id="button" className="flex flex-col w-full my-5">
        <button
          type="submit"
          className={cx({ btn: true, loading: isLoading })}
        >
          Register
        </button>
        <div className="flex justify-evenly mt-5">
          <button onClick={onClickLogin} className="btn btn-link">Login</button>
        </div>
      </div>
      {isError && <AlertError error={error as string} />}
    </form>
  )
}
