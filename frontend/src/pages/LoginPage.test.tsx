import { fireEvent, render, screen, waitFor } from "@testing-library/react"
import { QueryClient, QueryClientProvider } from "react-query"
import { BrowserRouter } from "react-router-dom"
import { AuthProvider } from "../hooks/AuthProvider"
import LoginPage from "./LoginPage"

describe('LoginPage', () => {
  it('should render correctly', async() => {
    const queryClient = new QueryClient();
    const Provider: React.FC<React.PropsWithChildren> = ({ children }) => (
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    );
    render(<Provider><AuthProvider><LoginPage /></AuthProvider></Provider>, { wrapper: BrowserRouter })
    const loginPageButtons = screen.getAllByRole('button')
    fireEvent.click(loginPageButtons[1])
    await waitFor(() => {
      const registerText =  screen.getByText('Register')
      expect(registerText).toBeInTheDocument()
    })
    const registerPageButtons = screen.getAllByRole('button')
    fireEvent.click(registerPageButtons[1])
    await waitFor(() => {
      const registerText =  screen.getByText('Login')
      expect(registerText).toBeInTheDocument()
    })
  })
})
