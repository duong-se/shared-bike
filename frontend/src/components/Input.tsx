import cx from "classnames";
export type InputProps = {
  name: string
  id?: string
  label?: string
  type?: 'text' | 'password'
  placeholder?: string
  error?: string
  onChange?: (e: React.ChangeEvent<any>) => void;
  className?: string
  value?: string
}

export const Input: React.FC<InputProps> = ({
  label,
  type = "text",
  id,
  placeholder = "",
  name,
  error,
  onChange,
  className,
  value
}) => {
  return (
    <div className="form-control w-full">
      <label htmlFor={id} className="label w-full">
        <span className="label-text">{label}</span>
      </label>
      <input
        aria-label={label}
        onChange={onChange}
        value={value}
        type={type}
        name={name}
        id={id}
        placeholder={placeholder}
        className={cx("input input-bordered w-full", { "input-error": error }, className)}
      />
      {error && (
        <label className="label">
          <span className="label-text-alt text-error">{error}</span>
        </label>
      )}
    </div>
  )
}
