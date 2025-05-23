import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'

// Create the app instance
const app = createApp(App)

// Set up Pinia store
const pinia = createPinia()
app.use(pinia)

// Mount the app
app.mount('#app')
