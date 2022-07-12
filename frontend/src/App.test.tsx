import { render, screen, waitFor } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import App from './App';

describe('App', () => {
  it('renders app', async () => {
    render(<App />, { wrapper: BrowserRouter});
    await waitFor(() => {
      const linkElement = screen.getByText(/Shared bike platform for everyone/i);
      expect(linkElement).toBeInTheDocument();
    })
  });
})
