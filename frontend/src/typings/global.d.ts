export type Config = {
  appEnv: 'development' | 'production'
  baseUrl: string
}

declare global {
  interface Window {
    sharedBike: {
      config: Config
    }
  }
}
