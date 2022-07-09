export type ButtonProps = React.PropsWithChildren & {
  variant: "primary" | "ghost"
  type?: "button" | "submit"
  onClick?: React.MouseEventHandler<HTMLButtonElement>
  disabled?: boolean
}

export const Button: React.FC<ButtonProps> = ({ variant, children, onClick, type="button", disabled = false }) => {
  if (variant === 'primary') {
    return (
      <button
        disabled={disabled}
        type={type}
        onClick={onClick}
        className="w-full py-4 bg-green-600 rounded-lg text-green-100"

      >
        <div className="flex flex-row items-center justify-center">
          <div className="font-bold">{children}</div>
        </div>
      </button>
    )
  }
  return (
    <button
    disabled={disabled}
    type={type}
    onClick={onClick}
    className="w-full text-center font-medium text-gray-500"

  >
    <div className="flex flex-row items-center justify-center">
      <div className="font-bold">{children}</div>
    </div>
  </button>
  )
}
