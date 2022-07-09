import { Button } from "./Button";
import { Input } from "./Input";

type RegisterFormProps = {
  onClickLogin: () => void;
}

export const RegisterForm: React.FC<RegisterFormProps> = ({ onClickLogin }) => {
  return (
    <form action="" className="w-full">
      <Input
        type="text"
        id="username"
        placeholder="Your username"
        name="username"
      />
      <Input
        type="password"
        id="password"
        placeholder="Your password"
        name="password"
      />
      <Input
        type="password"
        id="confirmPassword"
        name="confirmPassword"
        placeholder="Your password confirm"
      />
      <Input
        type="text"
        id="name"
        name="name"
        placeholder="Your name"
      />
      <div id="button" className="flex flex-col w-full my-5">
        <Button variant="primary">Register</Button>
        <div className="flex justify-evenly mt-5">
          <Button onClick={onClickLogin} variant="ghost">Login</Button>
        </div>
      </div>
    </form>
  );
}
