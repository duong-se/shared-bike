export type InputProps = {
  name: string
  id?: string
  label?: string
  type?: 'text' | 'password'
  placeholder?: string
  error?: string
}

export const Input: React.FC<InputProps> = ({ label, type = "text", id, placeholder = "", name, error }) => {
  return (
    <div className="form-control w-full">
      <label className="label w-full">
        <span className="label-text">{label}</span>
      </label>
      <input type={type} placeholder={placeholder} className="input input-bordered w-full" />
    </div>
  )
}
