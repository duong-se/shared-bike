import ReactDOM from 'react-dom/client'
import { Wrapper, Status } from '@googlemaps/react-wrapper'
import { BrowserRouter } from 'react-router-dom'
import App from './App'
import { Spinner } from './components/Spinner'
import './index.css'

// Init global config
window.sharedBike = {
  config: {
    googleMapApiKey: '',
    baseUrl: '',
    appEnv: 'development',
  },
}

const render = (status: Status) => {
  switch (status) {
  case Status.LOADING:
    return <Spinner />
  case Status.FAILURE:
    return <div>Failed...</div>
  case Status.SUCCESS:
    return <div>Success...</div>
  }
}

const renderApp = () => {
  const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
  )
  root.render(
    <Wrapper apiKey={window.sharedBike.config.googleMapApiKey} render={render}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </Wrapper>
  )
}

const run = async () => {
  await fetch('/config.json').then(response => response.json()).then(async (config) => {
    window.sharedBike.config = config
    renderApp()
  })
}

run()
