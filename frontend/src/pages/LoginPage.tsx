import { memo, useCallback, useState } from "react";
import { LoginForm } from "../components/LoginForm";
import { RegisterForm } from "../components/RegisterForm";
import {ReactComponent as BikeLogo} from './logo.svg';

export const LoginPage: React.FC = () => {
  const [isLogin, setIsLogin] = useState(true);
  const switchToLogin = useCallback(() => {
    setIsLogin(true);
  }, [])
  const switchToRegister = useCallback(() => {
    setIsLogin(false);
  }, [])
  return (
    <div className="container px-6 mx-auto">
      <div
        className="flex flex-col text-center md:text-left md:flex-row h-screen justify-evenly md:items-center"
      >
        <div className="flex flex-col w-full">
          <div className="w-1/5">
            <BikeLogo />
          </div>
          <h1 className="text-5xl text-gray-800 font-bold">Shared Bike</h1>
          <p className="w-5/12 mx-auto md:mx-0 text-gray-500">
            Shared bike platform for everyone
          </p>
        </div>
        <div className="w-full md:w-full lg:w-9/12 mx-auto md:mx-0">
          <div className="bg-white p-10 flex flex-col w-full shadow-xl rounded-xl">
            <h2 className="text-2xl font-bold text-gray-800 text-left mb-5">
              Sigin
            </h2>
            {
              isLogin ? <LoginForm onClickRegister={switchToRegister} /> : <RegisterForm onClickLogin={switchToLogin} />
            }
          </div>
        </div>
      </div>
    </div>
  );
}

export default memo(LoginPage)
