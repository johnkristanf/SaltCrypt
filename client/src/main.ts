import './assets/index.css';
import 'sweetalert2/dist/sweetalert2.min.css';

import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index'

const app = createApp(App)
app.use(router);
app.mount('#app');