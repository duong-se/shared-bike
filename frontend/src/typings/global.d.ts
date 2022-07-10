export type Config = {
  appEnv: 'development' | 'production'
  baseUrl: string
  googleMapApiKey: string
}

declare global {
  interface Window {
    sharedBike: {
      config: Config
    }
    google: typeof google
  }
}
