import Vue from 'vue'
import App from './App.vue'
import store from "./store";
import router from './routes'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'

// icons
import {
  faChevronLeft,
  faTimesCircle,
  faTrash,
  faLock,
  faEnvelope,
  faEye,
  faEyeSlash,
  faAngleLeft,
  faAngleRight,
  faArrowUp,
  faArrowDown,
  faExclamationCircle
} from "@fortawesome/free-solid-svg-icons";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
library.add(
  faChevronLeft,
  faTimesCircle,
  faTrash,
  faLock,
  faEnvelope,
  faEye,
  faEyeSlash,
  faAngleLeft,
  faAngleRight,
  faArrowUp,
  faArrowDown,
  faExclamationCircle
);

Vue.component("font-awesome-icon", FontAwesomeIcon);

Vue.use(Buefy, {
  defaultIconPack: "fa",
  defaultIconComponent: FontAwesomeIcon
});

Vue.config.productionTip = false

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount('#app')
