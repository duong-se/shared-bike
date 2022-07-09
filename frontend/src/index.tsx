import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from "react-router-dom";
import App from './App'
import './index.css'

// Init global config
window.sharedBike = {
  config: {
    baseUrl: '',
    appEnv: 'development',
  },
}

const renderApp = () => {
  const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
  );
  root.render(
    <React.StrictMode>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </React.StrictMode>
  );
}

const run = async () => {
  await fetch('/config.json').then(response => response.json()).then(async(config) => {
    window.sharedBike.config = config;
    renderApp()
  })
}

run()
