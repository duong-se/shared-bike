import { render, screen } from '@testing-library/react'
import { AlertError } from './AlertError'

describe('AlertError', () => {
  it('renders error correctly', async () => {
    const mockError = 'mock error string'
    render(<AlertError error={mockError} />)
    const error = (await screen.findByText(mockError)).textContent
    expect(error).toEqual(mockError)
  })

  it('does not render error and render only 1 div', async () => {
    const mockError = ''
    const { container } = render(<AlertError error={mockError} />)
    expect(container).toMatchSnapshot()
  })
})
