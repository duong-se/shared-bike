export type InputProps = {
  name: string
  id?: string
  label?: string
  type?: 'text' | 'password'
  placeholder?: string
}

export const Input: React.FC<InputProps> = ({ label, type="text", id, placeholder="", name }) => {
  return (
    <div id={`${id}-input`} className="flex flex-col w-full my-5">
      <label htmlFor={id} className="text-gray-500 mb-2">{label}</label>
      <input
        type={type}
        id={id}
        name={name}
        placeholder={placeholder}
        className="appearance-none border-2 border-gray-100 rounded-lg px-4 py-3 placeholder-gray-300 focus:outline-none focus:ring-2 focus:ring-green-600 focus:shadow-lg"
      />
    </div>
  )
}
