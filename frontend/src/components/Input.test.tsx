
import { fireEvent, render, screen } from '@testing-library/react'
import { Input } from './Input'

describe('Input', () => {
  it('should render correctly', () => {
    const mockProps = {
      label: "mockLabel",
      type: "text" as "text",
      id: "mockId",
      placeholder: "mockPlaceholder",
      name: "mockName",
      error: "mockError",
      onChange: jest.fn(),
      className: "mockClassName",
      value: "mockValue"
    }
    render(<Input {...mockProps} />)
    const label = screen.getByText(mockProps.label)
    expect(label).toBeInTheDocument()
    const input = screen.getByRole('textbox')
    expect(input).toBeInTheDocument()
    expect(Object.assign({},
      ...Array.from(input.attributes, ({ name, value }) => ({ [name]: value }))
    )).toStrictEqual({"aria-label": "mockLabel", "class": "input input-bordered w-full input-error mockClassName", "id": "mockId", "name": "mockName", "placeholder": "mockPlaceholder", "type": "text", "value": "mockValue"})
    const error = screen.getByText(mockProps.error)
    expect(error).toBeInTheDocument()
  })

  it('should call handle change', () => {
    const mockProps = {
      label: "mockLabel",
      type: "text" as "text",
      id: "mockId",
      placeholder: "mockPlaceholder",
      name: "mockName",
      onChange: jest.fn(),
      className: "mockClassName",
      value: "mockValue"
    }
    render(<Input {...mockProps} />)
    fireEvent.change(screen.getByRole('textbox'),{ target: { value: 'mockValueOne' } })
    expect(mockProps.onChange).toHaveBeenCalledTimes(1)
  })
})
